package app

import (
	"context"

	"admini.dev/admini/app/util"
)

type CoreServices struct{}

func initCoreServices(ctx context.Context, st *State, logger util.Logger) CoreServices {
	return CoreServices{}
}
