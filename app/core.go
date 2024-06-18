package app

import (
	"context"

	"admini.dev/admini/app/util"
)

type CoreServices struct{}

//nolint:revive
func initCoreServices(ctx context.Context, st *State, logger util.Logger) CoreServices {
	return CoreServices{}
}
