package logger

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var globalLogger *zap.SugaredLogger

func Init(debug bool) {
	level := zapcore.InfoLevel
	if debug {
		level = zapcore.DebugLevel
	}

	globalLogger = zap.New(
		createConsoleCore(level),
		zap.AddStacktrace(zap.LevelEnablerFunc(func(l zapcore.Level) bool {
			return zapcore.PanicLevel.Enabled(l)
		})),
	).Sugar()
}

func Sync() {
	_ = globalLogger.Sync()
}

//////////////////////////////////////////////////////////////////////

func init() {
	globalLogger = zap.New(createConsoleCore(zapcore.InfoLevel)).Sugar()
}

func createConsoleCore(minLevel zapcore.Level) zapcore.Core {
	return zapcore.NewCore(
		zapcore.NewConsoleEncoder(zap.NewDevelopmentEncoderConfig()),
		zapcore.AddSync(os.Stderr),
		zap.LevelEnablerFunc(func(level zapcore.Level) bool {
			return level >= minLevel
		}),
	)
}
