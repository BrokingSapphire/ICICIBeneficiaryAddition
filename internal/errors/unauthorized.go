package errors

import "github.com/brokingSapphire/SapphireICICI/internal/utils"

func NewUnauthorizedError(message string) *APIError {
	if message == "" {
		message = "Unauthorized"
	}
	return NewAPIError(utils.UNAUTHORIZED, message)
}
