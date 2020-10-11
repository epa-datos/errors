package errors

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Severity string

var messageList map[Code]string

type Kind struct {
	Desc       string
	HttpStatus int
	Code       Code
}

type Error struct {
	Kind       Kind
	Message    Message
	StackTrace string
	Operation  Operation
	Severity   Severity
}

func (e Error) Error() string {
	return string(e.Message)
}

func (e *Error) Wrap(op Operation, msg Message) {
	e.StackTrace += fmt.Sprintf("\n operation: %s, error: %s ;", op, msg)
}

func Build(args ...interface{}) error {
	var newError Error

	for _, arg := range args {
		switch val := arg.(type) {
		case Code:
			newError.Kind.Code = val
		case Operation:
			newError.Operation = val
		case Message:
			newError.Message = val
		case Severity:
			newError.Severity = val
		}
	}

	if newError.Kind.Code == 0 {
		newError.Kind.Code = Undefined
	}

	if newError.Severity == "" {
		newError.Severity = getSeverityFromCode(newError.Kind.Code)
	}

	newError.Wrap(newError.Operation, newError.Message)

	return newError
}

func IsErrType(err error) bool {
	_, isErrType := err.(Error)
	return isErrType
}

func JSON(c *gin.Context, err error) {
	e, IsErrType := err.(Error)
	Log(e)
	if !IsErrType {
		c.JSON(httpStatusCodes[InternalError], nil)
		return
	}

	if _, exists := httpStatusCodes[e.Kind.Code]; !exists {
		c.JSON(httpStatusCodes[e.Kind.Code], nil)
		return
	}

	switch e.Kind.Code {
	case InternalError, MaintenanceMode, Undefined:
		c.JSON(httpStatusCodes[e.Kind.Code], nil)
	default:
		c.JSON(httpStatusCodes[e.Kind.Code], gin.H{
			"message": e.Message,
		})
	}
}

func GetHttpStatusCode(err error) int {
	if errpyError, isErrpyError := err.(Error); isErrpyError {
		return errpyError.getHttpStatusCode()
	}
	return http.StatusInternalServerError
}

func (e Error) getHttpStatusCode() int {
	if httpCode, codeExists := httpStatusCodes[e.Kind.Code]; codeExists {
		return httpCode
	}
	return http.StatusInternalServerError
}

func GetErrpyCode(httpStatusCode int) Code {
	if errpyCode, isMapped := errpyCodes[httpStatusCode]; isMapped {
		return errpyCode
	}
	return Undefined
}
