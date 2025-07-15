package utils

import "net/http"

const (
    OK                     = http.StatusOK
    CREATED                = http.StatusCreated
    ACCEPTED               = http.StatusAccepted
    NO_CONTENT             = http.StatusNoContent
    BAD_REQUEST            = http.StatusBadRequest
    UNAUTHORIZED           = http.StatusUnauthorized
    FORBIDDEN              = http.StatusForbidden
    NOT_FOUND              = http.StatusNotFound
    METHOD_NOT_ALLOWED     = http.StatusMethodNotAllowed
    CONFLICT               = http.StatusConflict
    UNPROCESSABLE_ENTITY   = http.StatusUnprocessableEntity
    UNSUPPORTED_MEDIA_TYPE = http.StatusUnsupportedMediaType
    INTERNAL_SERVER_ERROR  = http.StatusInternalServerError
)