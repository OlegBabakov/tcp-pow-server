package zap

import (
	"errors"
	"os"
	"syscall"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	"github.com/OlegBabakov/pow-server/config"
)

type ZapLogger struct {
	cfg    config.LoggerConfig
	logger *zap.SugaredLogger
}

var loggerLevelMap = map[string]zapcore.Level{
	"debug":  zapcore.DebugLevel,
	"info":   zapcore.InfoLevel,
	"warn":   zapcore.WarnLevel,
	"error":  zapcore.ErrorLevel,
	"dpanic": zapcore.DPanicLevel,
	"panic":  zapcore.PanicLevel,
	"fatal":  zapcore.FatalLevel,
}

// NewZapLogger instance logger.
func NewZapLogger(cfg config.LoggerConfig) *ZapLogger {
	return &ZapLogger{cfg: cfg}
}

func (l *ZapLogger) getLoggerLevel(lv string) zapcore.Level {
	level, exist := loggerLevelMap[lv]
	if !exist {
		return zapcore.DebugLevel
	}

	return level
}

// InitLogger init logger.
func (l *ZapLogger) InitLogger(name string) {
	logLevel := l.getLoggerLevel(l.cfg.Level)

	encoderCfg := zap.NewProductionEncoderConfig()
	encoderCfg.EncodeTime = zapcore.ISO8601TimeEncoder

	stderrSyncer := zapcore.Lock(os.Stderr)

	l.logger = zap.New(
		zapcore.NewCore(zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig()),
			stderrSyncer,
			zap.NewAtomicLevelAt(logLevel)),
		zap.AddCaller(),
		zap.AddStacktrace(zapcore.LevelOf(zap.ErrorLevel)),
		zap.AddCallerSkip(2)).
		Sugar().
		Named(name)

	if err := l.logger.Sync(); err != nil && !errors.Is(err, syscall.ENOTTY) {
		l.logger.Error(err)
	}
}

func (l *ZapLogger) Debug(args ...interface{}) {
	l.logger.Debug(args...)
}

func (l *ZapLogger) Debugf(template string, args ...interface{}) {
	l.logger.Debugf(template, args...)
}

func (l *ZapLogger) Info(args ...interface{}) {
	l.logger.Info(args...)
}

func (l *ZapLogger) Infof(template string, args ...interface{}) {
	l.logger.Infof(template, args...)
}

func (l *ZapLogger) Warn(args ...interface{}) {
	l.logger.Warn(args...)
}

func (l *ZapLogger) Warnf(template string, args ...interface{}) {
	l.logger.Warnf(template, args...)
}

func (l *ZapLogger) Error(args ...interface{}) {
	l.logger.Error(args...)
}

func (l *ZapLogger) Errorf(template string, args ...interface{}) {
	l.logger.Errorf(template, args...)
}

func (l *ZapLogger) DPanic(args ...interface{}) {
	l.logger.DPanic(args...)
}

func (l *ZapLogger) DPanicf(template string, args ...interface{}) {
	l.logger.DPanicf(template, args...)
}

func (l *ZapLogger) Panic(args ...interface{}) {
	l.logger.Panic(args...)
}

func (l *ZapLogger) Panicf(template string, args ...interface{}) {
	l.logger.Panicf(template, args...)
}

func (l *ZapLogger) Fatal(args ...interface{}) {
	l.logger.Fatal(args...)
}

func (l *ZapLogger) Fatalf(template string, args ...interface{}) {
	l.logger.Fatalf(template, args...)
}
