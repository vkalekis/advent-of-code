package utils

import (
	"go.uber.org/zap"
)

func NewLogger() (*zap.SugaredLogger, error) {
	logger, err := zap.NewDevelopment()
	if err != nil {
		return nil, err
	}

	sugar := logger.Sugar()
	return sugar, nil
}
