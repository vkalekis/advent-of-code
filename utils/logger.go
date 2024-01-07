package utils

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var loggerLevelMap map[string]zapcore.Level

var Logger *zap.SugaredLogger

func initLoggerLevelMap() {
	loggerLevelMap = make(map[string]zapcore.Level)
	loggerLevelMap["debug"] = zap.DebugLevel
	loggerLevelMap["info"] = zap.InfoLevel
	loggerLevelMap["warn"] = zap.WarnLevel
	loggerLevelMap["error"] = zap.ErrorLevel
}

func NewLogger(level string) (*zap.SugaredLogger, error) {
	initLoggerLevelMap()

	stdout := zapcore.AddSync(os.Stdout)

	level1 := zap.NewAtomicLevelAt(loggerLevelMap[level])

	developmentCfg := zap.NewDevelopmentEncoderConfig()
	developmentCfg.EncodeLevel = zapcore.CapitalColorLevelEncoder

	consoleEncoder := zapcore.NewConsoleEncoder(developmentCfg)

	core := zapcore.NewTee(
		zapcore.NewCore(consoleEncoder, stdout, level1),
	)

	logger_ := zap.New(core)
	Logger = logger_.Sugar()

	return Logger, nil
}
