// Package search - Content managed by Project Forge, see [projectforge.md] for details.
package search

import (
	"context"
	"strings"

	"github.com/pkg/errors"
	"github.com/samber/lo"

	"admini.dev/admini/app"
	"admini.dev/admini/app/controller/cutil"
	"admini.dev/admini/app/lib/search/result"
	"admini.dev/admini/app/lib/telemetry"
	"admini.dev/admini/app/util"
)

type Provider func(context.Context, *Params, *app.State, *cutil.PageState, util.Logger) (result.Results, error)

func Search(ctx context.Context, params *Params, as *app.State, page *cutil.PageState) (result.Results, []error) {
	ctx, span, logger := telemetry.StartSpan(ctx, "search", page.Logger)
	defer span.Complete()

	if params.Q == "" {
		return nil, nil
	}
	var allProviders []Provider
	// $PF_SECTION_START(search_functions)$
	projectFunc := func(ctx context.Context, p *Params, as *app.State, page *cutil.PageState, logger util.Logger) (result.Results, error) {
		return as.Services.Projects.Search(ctx, p.Q, logger)
	}
	sourceFunc := func(ctx context.Context, p *Params, as *app.State, page *cutil.PageState, logger util.Logger) (result.Results, error) {
		return as.Services.Sources.Search(p.Q, logger)
	}
	allProviders = append(allProviders, projectFunc, sourceFunc)
	// $PF_SECTION_END(search_functions)$
	if len(allProviders) == 0 {
		return nil, []error{errors.New("no search providers configured")}
	}

	params.Q = strings.TrimSpace(params.Q)
	results, errs := util.AsyncCollect(allProviders, func(p Provider) ([]*result.Result, error) {
		return p(ctx, params, as, page, logger)
	})
	var ret result.Results = lo.Flatten(results)
	return ret.Sort(), errs
}
