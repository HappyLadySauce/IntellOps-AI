package glog

import (
	"context"

	"gofly/utils/tools/command"
	"gofly/utils/tools/gconv"
	"gofly/utils/tools/grpool"
)

// ILogger is the API interface for logger.
type ILogger interface {
	Print(ctx context.Context, v ...interface{})
	Printf(ctx context.Context, format string, v ...interface{})
	Debug(ctx context.Context, v ...interface{})
	Debugf(ctx context.Context, format string, v ...interface{})
	Info(ctx context.Context, v ...interface{})
	Infof(ctx context.Context, format string, v ...interface{})
	Notice(ctx context.Context, v ...interface{})
	Noticef(ctx context.Context, format string, v ...interface{})
	Warning(ctx context.Context, v ...interface{})
	Warningf(ctx context.Context, format string, v ...interface{})
	Error(ctx context.Context, v ...interface{})
	Errorf(ctx context.Context, format string, v ...interface{})
	Critical(ctx context.Context, v ...interface{})
	Criticalf(ctx context.Context, format string, v ...interface{})
	Panic(ctx context.Context, v ...interface{})
	Panicf(ctx context.Context, format string, v ...interface{})
	Fatal(ctx context.Context, v ...interface{})
	Fatalf(ctx context.Context, format string, v ...interface{})
}

const (
	commandEnvKeyForDebug = "gf.glog.debug"
)

var (
	// Ensure Logger implements ILogger.
	_ ILogger = &Logger{}

	// Default logger object, for package method usage.
	defaultLogger = New()

	// Goroutine pool for async logging output.
	// It uses only one asynchronous worker to ensure log sequence.
	asyncPool = grpool.New(1)

	// defaultDebug enables debug level or not in default,
	// which can be configured using command option or system environment.
	defaultDebug = true
)

func init() {
	defaultDebug = gconv.Bool(command.GetOptWithEnv(commandEnvKeyForDebug, "true"))
	SetDebug(defaultDebug)
}

// DefaultLogger returns the default logger.
func DefaultLogger() *Logger {
	return defaultLogger
}

// SetDefaultLogger sets the default logger for package glog.
// Note that there might be concurrent safety issue if calls this function
// in different goroutines.
func SetDefaultLogger(l *Logger) {
	defaultLogger = l
}
