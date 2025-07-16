package errors

var (
	// BadRequestError constructor
	BadRequest = NewBadRequestError

	// UnauthorizedError constructor
	Unauthorized = NewUnauthorizedError

	// ForbiddenError constructor
	Forbidden = NewForbiddenError

	// NotFoundError constructor
	NotFound = NewNotFoundError

	// MethodNotAllowedError constructor
	MethodNotAllowed = NewMethodNotAllowedError

	// ConflictError constructor
	Conflict = NewConflictError

	// UnprocessableEntityError constructor
	UnprocessableEntity = NewUnprocessableEntityError

	// UnsupportedMediaTypeError constructor
	UnsupportedMediaType = NewUnsupportedMediaTypeError

	// InternalServerError constructor
	InternalServer = NewInternalServerError
)
