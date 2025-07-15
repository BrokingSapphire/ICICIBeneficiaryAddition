package errors

import "github.com/brokingSapphire/SapphireICICI/internal/utils"

func NewForbiddenError(message string) *APIError {
    if message == "" {
        message = "Forbidden"
    }
    return NewAPIError(utils.FORBIDDEN, message)
}