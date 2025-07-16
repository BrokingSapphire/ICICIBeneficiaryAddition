package errors

import "github.com/brokingSapphire/SapphireICICI/internal/utils"

func NewMethodNotAllowedError(message string) *APIError {
	if message == "" {
		message = "Method Not Allowed"
	}
	return NewAPIError(utils.METHOD_NOT_ALLOWED, message)
}
