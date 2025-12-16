package errors

import "net/http"

type Code uint

const (
	_ Code = iota
	InternalError
	ExternalServiceUnavailable
	BadRequest
	NotFound
	Forbidden
	Busy
	MaintenanceMode
	DuplicatedEntry
	Undefined
	Conflict
	Unauthorized
	TooManyRequests
	RequestTimeout
	UnprocessableEntity
	ServiceUnavailable
	GatewayTimeout
	UnsupportedMediaType
)

var httpStatusCodes = map[Code]int{
	InternalError:              http.StatusInternalServerError,
	NotFound:                   http.StatusNotFound,
	BadRequest:                 http.StatusBadRequest,
	Forbidden:                  http.StatusForbidden,
	MaintenanceMode:            http.StatusServiceUnavailable,
	Undefined:                  http.StatusInternalServerError,
	Conflict:                   http.StatusConflict,
	Unauthorized:               http.StatusUnauthorized,
	TooManyRequests:            http.StatusTooManyRequests,
	DuplicatedEntry:            http.StatusConflict,
	ExternalServiceUnavailable: http.StatusBadGateway,
	RequestTimeout:             http.StatusRequestTimeout,
	UnprocessableEntity:        http.StatusUnprocessableEntity,
	ServiceUnavailable:         http.StatusServiceUnavailable,
	GatewayTimeout:             http.StatusGatewayTimeout,
	UnsupportedMediaType:       http.StatusUnsupportedMediaType,
}

var errpyCodes = map[int]Code{
	http.StatusInternalServerError:  InternalError,
	http.StatusNotFound:             NotFound,
	http.StatusBadRequest:           BadRequest,
	http.StatusForbidden:            Forbidden,
	http.StatusServiceUnavailable:   MaintenanceMode,
	http.StatusConflict:             Conflict,
	http.StatusUnauthorized:         Unauthorized,
	http.StatusTooManyRequests:      TooManyRequests,
	http.StatusBadGateway:           ExternalServiceUnavailable,
	http.StatusRequestTimeout:       RequestTimeout,
	http.StatusUnprocessableEntity:  UnprocessableEntity,
	http.StatusGatewayTimeout:       GatewayTimeout,
	http.StatusUnsupportedMediaType: UnsupportedMediaType,
}
