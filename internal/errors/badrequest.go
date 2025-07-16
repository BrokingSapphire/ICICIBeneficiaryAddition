package errors

import "github.com/brokingSapphire/SapphireICICI/internal/utils"

func NewBadRequestError(message string) *APIError {
	if message == "" {
		message = "Bad Request"
	}
	return NewAPIError(utils.BAD_REQUEST, message)
}
