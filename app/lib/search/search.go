// Content managed by Project Forge, see [projectforge.md] for details.
package search

import (
	"context"
	"strings"

	"github.com/pkg/errors"

	"admini.dev/admini/app"
	"admini.dev/admini/app/lib/search/result"
	"admini.dev/admini/app/lib/telemetry"
	"admini.dev/admini/app/util"
)

type Provider func(context.Context, *app.State, *Params, util.Logger) (result.Results, error)

func Search(ctx context.Context, as *app.State, params *Params, logger util.Logger) (result.Results, []error) {
	ctx, span, logger := telemetry.StartSpan(ctx, "search", logger)
	defer span.Complete()

	if params.Q == "" {
		return nil, nil
	}
	var allProviders []Provider
	// $PF_SECTION_START(search_functions)$
	projectFunc := func(ctx context.Context, as *app.State, p *Params, logger util.Logger) (result.Results, error) {
		return as.Services.Projects.Search(ctx, p.Q, logger)
	}
	sourceFunc := func(ctx context.Context, as *app.State, p *Params, logger util.Logger) (result.Results, error) {
		return as.Services.Sources.Search(p.Q, logger)
	}
	allProviders = append(allProviders, projectFunc, sourceFunc)
	// $PF_SECTION_END(search_functions)$
	// $PF_INJECT_START(codegen)$
	// $PF_INJECT_END(codegen)$
	if len(allProviders) == 0 {
		return nil, []error{errors.New("no search providers configured")}
	}

	params.Q = strings.TrimSpace(params.Q)

	results, errs := util.AsyncCollect(allProviders, func(item Provider) (result.Results, error) {
		return item(ctx, as, params, logger)
	})

	ret := make(result.Results, 0, len(results)*len(results))
	for _, x := range results {
		ret = append(ret, x...)
	}

	ret.Sort()
	return ret, errs
}
