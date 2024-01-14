package logger

// Logger is an interface of log application.
type Logger interface {
	// InitLogger created logger instance
	InitLogger(name string)
	// Debug represents a debug log
	Debug(args ...interface{})
	// Debugf represents a debug formatted log
	Debugf(template string, args ...interface{})
	// Info represents a debug log
	Info(args ...interface{})
	// Infof represents a debug formatted log
	Infof(template string, args ...interface{})
	// Warn represents a debug log
	Warn(args ...interface{})
	// Warnf represents a debug formatted log
	Warnf(template string, args ...interface{})
	// Error represents a debug log
	Error(args ...interface{})
	// Errorf represents a debug formatted log
	Errorf(template string, args ...interface{})
	// DPanic represents a debug log
	DPanic(args ...interface{})
	// DPanicf represents a debug formatted log
	DPanicf(template string, args ...interface{})
	// Fatal represents a debug log
	Fatal(args ...interface{})
	// Fatalf represents a debug formatted log
	Fatalf(template string, args ...interface{})
}
