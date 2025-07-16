package errors

import "github.com/brokingSapphire/SapphireICICI/internal/utils"

func NewUnprocessableEntityError(message string) *APIError {
	if message == "" {
		message = "Unprocessable Entity"
	}
	return NewAPIError(utils.UNPROCESSABLE_ENTITY, message)
}
