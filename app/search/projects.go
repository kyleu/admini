package search

import (
	"context"
	"strings"

	"github.com/kyleu/admini/app"
	"github.com/kyleu/admini/app/project"
)

func searchProjects(ctx context.Context, st *app.State, p *Params) (Results, error) {
	var ret Results
	prjs, _ := st.Services.Projects.List()
	for _, prj := range prjs {
		if m := projectMatches(prj, p.Q); len(m) > 0 {
			res := &Result{ID: prj.Key, Type: "project", Title: prj.Name(), Icon: prj.IconWithFallback(), URL: "/project/" + prj.Key, Matches: MatchesFrom(m), Data: prj}
			ret = append(ret, res)
		}
	}

	return ret, nil
}

func projectMatches(prj *project.Project, q string) []string {
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
