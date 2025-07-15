package errors

import "github.com/brokingSapphire/SapphireICICI/internal/utils"

func NewNotFoundError(message string) *APIError {
    if message == "" {
        message = "Not Found"
    }
    return NewAPIError(utils.NOT_FOUND, message)
}