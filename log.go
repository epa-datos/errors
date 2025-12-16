package errors

import log "github.com/sirupsen/logrus"

const (
	WarningSeverity = "Warning"
	DebugSeverity   = "Debug"
	InfoSeverity    = "Info"
	ErrorSeverity   = "Error"
)

func Log(err error) {
	var er Error
	var isErrType bool

	if er, isErrType = err.(Error); isErrType {
		logEntry := log.WithFields(log.Fields{
			"operation":     er.Operation,
			"code":          er.Kind.Code,
			"reason":        er.Message,
			"trace":         er.StackTrace,
			"severity_type": er.Severity,
			"http_status":   er.Kind.HttpStatus,
		})
		displayError(logEntry, er)
	} else {
		log.Error("Undefined Error:", err.Error())
	}
}

func displayError(logEntry *log.Entry, err Error) {
	switch err.Severity {
	case WarningSeverity:
		logEntry.Warning(err.Message)
	case ErrorSeverity:
		logEntry.Error(err.Message)
	case DebugSeverity:
		logEntry.Debug(err.Message)
	case InfoSeverity:
		logEntry.Info(err.Message)
	default:
		logEntry.Error(err.Message)
	}
}

func getSeverityFromCode(code Code) Severity {
	switch code {
	case InternalError, GatewayTimeout, ServiceUnavailable, ExternalServiceUnavailable:
		return ErrorSeverity
	case BadRequest, UnprocessableEntity, UnsupportedMediaType:
		return WarningSeverity
	case NotFound, MaintenanceMode, RequestTimeout:
		return InfoSeverity
	}
	return ErrorSeverity
}
