// Package cmd - Content managed by Project Forge, see [projectforge.md] for details.
package cmd

import (
	"admini.dev/admini/app"
	"admini.dev/admini/app/util"
)

func Run(bi *app.BuildInfo) (util.Logger, error) {
	_buildInfo = bi

	if err := rootCmd().Execute(); err != nil {
		return _logger, err
	}
	return _logger, nil
}
