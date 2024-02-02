package logger

import (
	"os"

	gelf "github.com/snovichkov/zap-gelf"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var globalLogger *zap.SugaredLogger

func InitLocal(debug bool) {
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

func Init(serviceName string, graylogEndpoint string, debug bool) {
	level := zapcore.InfoLevel
	if debug {
		level = zapcore.DebugLevel
	}

	core, err := gelf.NewCore(
		gelf.Host(serviceName),
		gelf.Addr(graylogEndpoint),
		gelf.Level(level),
	)
	if err != nil {
		globalLogger.Panic("gelf.NewCore", "error", err)
	}

	globalLogger = zap.New(
		zapcore.NewTee(core, createConsoleCore(level)),
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
