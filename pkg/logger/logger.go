package logger

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var loggerLevelMap map[string]zapcore.Level

var logger *zap.SugaredLogger

func Debugf(format string, args ...interface{}) {
	logger.Debugf(format, args...)
}

func Infof(format string, args ...interface{}) {
	logger.Infof(format, args...)
}

func Warnf(format string, args ...interface{}) {
	logger.Warnf(format, args...)
}

func Errorf(format string, args ...interface{}) {
	logger.Errorf(format, args...)
}

func Fatalf(format string, args ...interface{}) {
	logger.Fatalf(format, args...)
}

func Debugln(args ...interface{}) {
	logger.Debug(args...)
}

func Infoln(args ...interface{}) {
	logger.Info(args...)
}

func Warnln(args ...interface{}) {
	logger.Warn(args...)
}

func Errorln(args ...interface{}) {
	logger.Error(args...)
}

func Fatalln(args ...interface{}) {
	logger.Fatal(args...)
}

func initLoggerLevelMap() {
	loggerLevelMap = make(map[string]zapcore.Level)
	loggerLevelMap["debug"] = zap.DebugLevel
	loggerLevelMap["info"] = zap.InfoLevel
	loggerLevelMap["warn"] = zap.WarnLevel
	loggerLevelMap["error"] = zap.ErrorLevel
}

func NewLogger(level string) error {
	initLoggerLevelMap()

	stdout := zapcore.AddSync(os.Stdout)

	level1 := zap.NewAtomicLevelAt(loggerLevelMap[level])

	developmentCfg := zap.NewDevelopmentEncoderConfig()
	developmentCfg.EncodeLevel = zapcore.CapitalColorLevelEncoder

	consoleEncoder := zapcore.NewConsoleEncoder(developmentCfg)

	core := zapcore.NewTee(
		zapcore.NewCore(consoleEncoder, stdout, level1),
	)

	logger_ := zap.New(core, zap.AddCaller())
	logger = logger_.Sugar()

	return nil
}
