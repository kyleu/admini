package cmd

import (
	"github.com/kyleu/admini/app"
	"go.uber.org/zap"
)

func Run(bi *app.BuildInfo) (*zap.SugaredLogger, error) {
	_buildInfo = bi

	err := rootCmd().Execute()
	if err != nil {
		return _logger, err
	}
	return _logger, nil
}
