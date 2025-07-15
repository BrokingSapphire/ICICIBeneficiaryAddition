package middleware

import (
    "github.com/brokingSapphire/SapphireICICI/internal/errors"
    "github.com/gin-gonic/gin"
)

// NotFoundHandler handles 404 errors
func NotFoundHandler() gin.HandlerFunc {
    return func(c *gin.Context) {
        err := errors.NewNotFoundError("Route not found")
        
        c.JSON(err.StatusCode, gin.H{
            "error": gin.H{
                "code":    err.StatusCode,
                "message": err.Message,
            },
        })
    }
}