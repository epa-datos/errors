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
	}
}

func getSeverityFromCode(code Code) Severity {
	switch code {
	case InternalError:
		return ErrorSeverity
	case BadRequest:
		return WarningSeverity
	case NotFound:
		return ErrorSeverity
	}
	return ErrorSeverity
}
