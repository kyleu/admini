// Content managed by Project Forge, see [projectforge.md] for details.
package cmd

import (
	"context"

	"github.com/pkg/errors"

	"admini.dev/admini/app"
	"admini.dev/admini/app/lib/filesystem"
	"admini.dev/admini/app/lib/telemetry"
	"admini.dev/admini/app/util"
)

func buildDefaultAppState(flags *Flags, logger util.Logger) (*app.State, error) {
	f := filesystem.NewFileSystem(flags.ConfigDir, logger)
	telemetryDisabled := util.GetEnvBool("disable_telemetry", false)
	st, err := app.NewState(flags.Debug, _buildInfo, f, !telemetryDisabled, logger)
	if err != nil {
		return nil, err
	}

	ctx, span, logger := telemetry.StartSpan(context.Background(), "app:init", logger)
	defer span.Complete()

	svcs, err := app.NewServices(ctx, st)
	if err != nil {
		return nil, errors.Wrap(err, "error creating services")
	}
	st.Services = svcs

	return st, nil
}
