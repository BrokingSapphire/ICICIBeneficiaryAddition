package errors

import "github.com/brokingSapphire/SapphireICICI/internal/utils"

func NewConflictError(message string) *APIError {
    if message == "" {
        message = "Conflict"
    }
    return NewAPIError(utils.CONFLICT, message)
}