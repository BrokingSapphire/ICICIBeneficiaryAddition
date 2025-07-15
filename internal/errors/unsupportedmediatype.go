package errors

import "github.com/brokingSapphire/SapphireICICI/internal/utils"

func NewUnsupportedMediaTypeError(message string) *APIError {
    if message == "" {
        message = "Unsupported Media Type"
    }
    return NewAPIError(utils.UNSUPPORTED_MEDIA_TYPE, message)
}