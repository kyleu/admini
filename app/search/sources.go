package search

import (
	"context"
	"strings"

	"github.com/kyleu/admini/app"
	"github.com/kyleu/admini/app/source"
)

func searchSources(ctx context.Context, st *app.State, p *Params) (Results, error) {
	var ret Results
	prjs, _ := st.Services.Sources.List()
	for _, prj := range prjs {
		if m := sourceMatches(prj, p.Q); len(m) > 0 {
			res := &Result{ID: prj.Key, Type: "source", Title: prj.Name(), Icon: prj.IconWithFallback(), URL: "/source/" + prj.Key, Matches: MatchesFrom(m), Data: prj}
			ret = append(ret, res)
		}
	}

	return ret, nil
}

func sourceMatches(prj *source.Source, q string) []string {
	var ret []string
	ql := strings.ToLower(q)
	f := func(k string, v string) {
		if strings.Contains(strings.ToLower(v), ql) {
			ret = append(ret, k+": "+v)
		}
	}
	f("key", prj.Key)
	f("title", prj.Title)
	f("description", prj.Description)
	return ret
}
