package middleware

import (
    "github.com/brokingSapphire/SapphireICICI/internal/env"
    "github.com/brokingSapphire/SapphireICICI/internal/errors"
    "github.com/brokingSapphire/SapphireICICI/internal/logger"
    "github.com/brokingSapphire/SapphireICICI/internal/utils"
    "github.com/gin-gonic/gin"
    "github.com/lib/pq"
    "github.com/sirupsen/logrus" 
)

type ErrorResponse struct {
    Error ErrorDetail `json:"error"`
}

type ErrorDetail struct {
    Code    int      `json:"code"`
    Message string   `json:"message"`
    Details []string `json:"details,omitempty"`
}

// ErrorHandler is the global error handling middleware
func ErrorHandler() gin.HandlerFunc {
    return func(c *gin.Context) {
        c.Next()
        
        // Check if there are any errors
        if len(c.Errors) > 0 {
            err := c.Errors.Last().Err
            
            // Log the error
            logger.ErrorWithFields(logrus.Fields{
                "method": c.Request.Method,
                "path":   c.Request.URL.Path,
                "ip":     c.ClientIP(),
            }, err)
            
            // Handle different error types
            handleError(c, err)
        }
    }
}

func handleError(c *gin.Context, err error) {
    // Don't send response if headers already sent
    if c.Writer.Written() {
        return
    }
    
    // Handle APIError
    if apiErr, ok := err.(*errors.APIError); ok {
        c.JSON(apiErr.StatusCode, ErrorResponse{
            Error: ErrorDetail{
                Code:    apiErr.StatusCode,
                Message: apiErr.Message,
            },
        })
        return
    }
    
    // Handle PostgreSQL errors
    if pqErr, ok := err.(*pq.Error); ok {
        handleDatabaseError(c, pqErr)
        return
    }
    
    // Handle other known error types
    switch err.Error() {
    case "record not found":
        c.JSON(utils.NOT_FOUND, ErrorResponse{
            Error: ErrorDetail{
                Code:    utils.NOT_FOUND,
                Message: "No result found",
            },
        })
        return
    }
    
    // Default internal server error
    errorResponse := ErrorResponse{
        Error: ErrorDetail{
            Code:    utils.INTERNAL_SERVER_ERROR,
            Message: "Something went wrong!",
        },
    }
    
    // Add error details in development
    if env.Env.Env == "development" {
        errorResponse.Error.Details = getErrorDetails(err)
    }
    
    c.JSON(utils.INTERNAL_SERVER_ERROR, errorResponse)
}

func handleDatabaseError(c *gin.Context, pqErr *pq.Error) {
    switch pqErr.Code {
    case "23505": // unique_violation
        c.JSON(utils.BAD_REQUEST, ErrorResponse{
            Error: ErrorDetail{
                Code:    utils.BAD_REQUEST,
                Message: "Duplicate entry",
            },
        })
    case "23503": // foreign_key_violation
        c.JSON(utils.BAD_REQUEST, ErrorResponse{
            Error: ErrorDetail{
                Code:    utils.BAD_REQUEST,
                Message: "Invalid reference",
            },
        })
    case "23514": // check_violation
        c.JSON(utils.BAD_REQUEST, ErrorResponse{
            Error: ErrorDetail{
                Code:    utils.BAD_REQUEST,
                Message: "Validation failed",
            },
        })
    case "23502": // not_null_violation
        c.JSON(utils.BAD_REQUEST, ErrorResponse{
            Error: ErrorDetail{
                Code:    utils.BAD_REQUEST,
                Message: "Required field missing",
            },
        })
    default:
        c.JSON(utils.INTERNAL_SERVER_ERROR, ErrorResponse{
            Error: ErrorDetail{
                Code:    utils.INTERNAL_SERVER_ERROR,
                Message: "Database error",
            },
        })
    }
}

func getErrorDetails(err error) []string {
    details := []string{err.Error()}
    return details
}