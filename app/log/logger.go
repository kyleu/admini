package log

import (
	"github.com/pkg/errors"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func InitLogging(debug bool) (*zap.SugaredLogger, error) {
	var logger *zap.Logger
	var err error
	if debug {
		logger, err = initDevLogging()
	} else {
		logger, err = initProdLogging()
	}
	if err != nil {
		return nil, errors.Wrap(err, "error initializing logging")
	}
	return logger.Sugar(), nil
}

func initDevLogging() (*zap.Logger, error) {
	err := zap.RegisterEncoder("custom", func(cfg zapcore.EncoderConfig) (zapcore.Encoder, error) {
		return NewEncoder(cfg), nil
	})
	if err != nil {
		return nil, errors.Wrap(err, "error registering encoder")
	}
	config := zap.NewDevelopmentConfig()
	config.EncoderConfig = zapcore.EncoderConfig{}

	config.Encoding = "custom"
	config.Development = true
	return config.Build(zap.AddStacktrace(zap.PanicLevel), zap.AddCaller())
}

func initProdLogging() (*zap.Logger, error) {
	config := zap.NewProductionConfig()
	config.EncoderConfig.FunctionKey = "func"
	return config.Build(zap.AddStacktrace(zap.PanicLevel), zap.AddCaller())
}
