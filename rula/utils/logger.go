package utils

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/atrariksa/fastrogos/rula/configs"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/sirupsen/logrus"
)

const (
	// FatalLevel level. Logs and then calls `logger.Exit(1)`. It will exit even if the
	// logging level is set to Panic.
	FatalLevel = "FATAL"
	// ErrorLevel level. Logs. Used for errors that should definitely be noted.
	// Commonly used for hooks to send errors to an error tracking service.
	ErrorLevel = "ERROR"
	// WarnLevel level. Non-critical entries that deserve eyes.
	WarnLevel = "WARN"
	// InfoLevel level. General operational entries about what's going on inside the
	// application.
	InfoLevel = "INFO"
	// DebugLevel level. Usually only enabled when debugging. Very verbose logging.
	DebugLevel = "DEBUG"
	// TraceLevel level. Designates finer-grained informational events than the Debug.
	TraceLevel = "TRACE"
)

func GetLogger(cfg *configs.Config) *logrus.Logger {
	logger := logrus.New()
	logger.Formatter = &logrus.JSONFormatter{
		DisableTimestamp: true,
	}
	logLevel := strings.ToUpper(strings.TrimSpace(cfg.App.LogLevel))
	switch logLevel {
	case FatalLevel:
		logger.SetLevel(logrus.DebugLevel)
	case ErrorLevel:
		logger.SetLevel(logrus.ErrorLevel)
	case WarnLevel:
		logger.SetLevel(logrus.WarnLevel)
	case DebugLevel:
		logger.SetLevel(logrus.DebugLevel)
	case TraceLevel:
		logger.SetLevel(logrus.TraceLevel)
	default:
		logger.SetLevel(logrus.InfoLevel)
	}
	return logger
}

func NewStructuredLogger(logger *logrus.Logger) func(next http.Handler) http.Handler {
	return middleware.RequestLogger(&StructuredLogger{logger})
}

type StructuredLogger struct {
	Logger *logrus.Logger
}

func (l *StructuredLogger) NewLogEntry(r *http.Request) middleware.LogEntry {
	entry := &StructuredLoggerEntry{Logger: logrus.NewEntry(l.Logger)}
	logFields := logrus.Fields{}

	logFields["ts"] = time.Now().UTC().Format(time.RFC1123)

	if reqID := middleware.GetReqID(r.Context()); reqID != "" {
		logFields["req_id"] = reqID
	}

	scheme := "http"
	if r.TLS != nil {
		scheme = "https"
	}
	logFields["http_scheme"] = scheme
	logFields["http_proto"] = r.Proto
	logFields["http_method"] = r.Method

	logFields["remote_addr"] = r.RemoteAddr
	logFields["user_agent"] = r.UserAgent()

	logFields["uri"] = fmt.Sprintf("%s://%s%s", scheme, r.Host, r.RequestURI)

	entry.Logger = entry.Logger.WithFields(logFields)

	entry.Logger.Infoln("request started")

	return entry
}

type StructuredLoggerEntry struct {
	Logger logrus.FieldLogger
}

func (l *StructuredLoggerEntry) Write(status, bytes int, header http.Header, elapsed time.Duration, extra interface{}) {
	l.Logger = l.Logger.WithFields(logrus.Fields{
		"resp_status": status, "resp_bytes_length": bytes,
		"resp_elapsed_ms": float64(elapsed.Nanoseconds()) / 1000000.0,
	})

	l.Logger.Infoln("request complete")
}

func (l *StructuredLoggerEntry) Panic(v interface{}, stack []byte) {
	l.Logger = l.Logger.WithFields(logrus.Fields{
		"stack": string(stack),
		"panic": fmt.Sprintf("%+v", v),
	})
}

// Helper methods used by the application to get the request-scoped
// logger entry and set additional fields between handlers.
//
// This is a useful pattern to use to set state on the entry as it
// passes through the handler chain, which at any point can be logged
// with a call to .Print(), .Info(), etc.

func GetLogEntry(r *http.Request) logrus.FieldLogger {
	entry := middleware.GetLogEntry(r).(*StructuredLoggerEntry)
	return entry.Logger
}

func LogEntrySetField(r *http.Request, key string, value interface{}) {
	if entry, ok := r.Context().Value(middleware.LogEntryCtxKey).(*StructuredLoggerEntry); ok {
		entry.Logger = entry.Logger.WithField(key, value)
	}
}

func LogEntrySetFields(r *http.Request, fields map[string]interface{}) {
	if entry, ok := r.Context().Value(middleware.LogEntryCtxKey).(*StructuredLoggerEntry); ok {
		entry.Logger = entry.Logger.WithFields(fields)
	}
}
