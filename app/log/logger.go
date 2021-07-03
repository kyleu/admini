package log

import (
	"github.com/pkg/errors"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func InitLogging(debug bool, json bool) (*zap.SugaredLogger, error) {
	var logger *zap.Logger
	var err error
	if json {
		logger, err = initJSONLogging()
	} else {
		if debug {
			logger, err = initDevLogging()
		} else {
			logger, err = initSimpleLogging()
		}
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
	return config.Build(zap.AddStacktrace(zap.PanicLevel), zap.AddCaller())
}

func initJSONLogging() (*zap.Logger, error) {
	config := zap.NewProductionConfig()
	config.EncoderConfig.FunctionKey = "func"
	return config.Build(zap.AddStacktrace(zap.PanicLevel), zap.AddCaller())
}

func initSimpleLogging() (*zap.Logger, error) {
	err := zap.RegisterEncoder("custom", func(cfg zapcore.EncoderConfig) (zapcore.Encoder, error) {
		return SimpleEncoder(cfg), nil
	})
	if err != nil {
		return nil, errors.Wrap(err, "error registering encoder")
	}
	config := zap.NewDevelopmentConfig()
	config.EncoderConfig = zapcore.EncoderConfig{}

	config.Encoding = "custom"
	config.EncoderConfig.FunctionKey = "func"
	return config.Build(zap.AddStacktrace(zap.PanicLevel), zap.AddCaller())
}
