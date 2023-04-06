// Content managed by Project Forge, see [projectforge.md] for details.
package app

import (
	"context"
	"fmt"
	"sync"
	"time"

	"admini.dev/admini/app/lib/auth"
	"admini.dev/admini/app/lib/filesystem"
	"admini.dev/admini/app/lib/telemetry"
	"admini.dev/admini/app/lib/theme"
	"admini.dev/admini/app/util"
)

var once sync.Once

type BuildInfo struct {
	Version string `json:"version"`
	Commit  string `json:"commit"`
	Date    string `json:"date"`
}

func (b *BuildInfo) String() string {
	if b.Date == "unknown" {
		return b.Version
	}
	d, _ := util.TimeFromJS(b.Date)
	return fmt.Sprintf("%s (%s)", b.Version, util.TimeToYMD(d))
}

type State struct {
	Debug     bool
	BuildInfo *BuildInfo
	Files     filesystem.FileLoader
	Auth      *auth.Service
	Themes    *theme.Service
	Services  *Services
	Started   time.Time
}

func NewState(debug bool, bi *BuildInfo, f filesystem.FileLoader, enableTelemetry bool, port uint16, logger util.Logger) (*State, error) {
	var loadLocationError error
	once.Do(func() {
		loc, err := time.LoadLocation("UTC")
		if err != nil {
			loadLocationError = err
			return
		}
		time.Local = loc
	})
	if loadLocationError != nil {
		return nil, loadLocationError
	}

	_ = telemetry.InitializeIfNeeded(enableTelemetry, bi.Version, logger)
	as := auth.NewService("", port, logger)
	ts := theme.NewService(f)

	return &State{
		Debug:     debug,
		BuildInfo: bi,
		Files:     f,
		Auth:      as,
		Themes:    ts,
		Started:   time.Now(),
	}, nil
}

func (s State) Close(ctx context.Context, logger util.Logger) error {
	defer func() { _ = telemetry.Close(ctx) }()
	return s.Services.Close(ctx, logger)
}
