package errors

import "github.com/brokingSapphire/SapphireICICI/internal/utils"

func NewInternalServerError(message string) *APIError {
    if message == "" {
        message = "Internal Server Error"
    }
    return NewAPIError(utils.INTERNAL_SERVER_ERROR, message)
}