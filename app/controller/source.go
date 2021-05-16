package controller

import (
	"fmt"
	"github.com/kyleu/admini/app/database"
	"github.com/kyleu/admini/app/schema"
	"github.com/kyleu/admini/app/source"
	"github.com/kyleu/admini/app/util"
	"github.com/kyleu/admini/views"
	"net/http"

	"github.com/pkg/errors"

	"github.com/kyleu/admini/app/controller/cutil"

	"github.com/kyleu/admini/app"
	"github.com/kyleu/admini/views/vsource"

	"github.com/gorilla/mux"
)

func SourceList(w http.ResponseWriter, r *http.Request) {
	act("source.list", w, r, func(as *app.State, ps *cutil.PageState) (string, error) {
		s, err := as.Sources.List()
		if err != nil {
			return "", errors.Wrap(err, "unable to load source list")
		}
		ps.Data = s
		return render(r, w, as, &vsource.SourceList{Sources: s}, ps, "sources")
	})
}

func SourceDetail(w http.ResponseWriter, r *http.Request) {
	act("source.detail", w, r, func(as *app.State, ps *cutil.PageState) (string, error) {
		key := mux.Vars(r)["key"]
		src, err := as.Sources.Load(key)
		if err != nil {
			return "", errors.Wrap(err, "unable to load source ["+key+"]")
		}
		sch, err := as.Sources.SchemaFor(key)
		if err != nil {
			return "", errors.Wrap(err, "unable to load schema for source ["+key+"]")
		}
		ps.Data = map[string]interface{}{"source": src, "schema": sch}
		return render(r, w, as, &vsource.SourceDetail{Source: src, Schema: sch}, ps, "sources", src.Key)
	})
}

func SourceNew(w http.ResponseWriter, r *http.Request) {
	act("source.new", w, r, func(as *app.State, ps *cutil.PageState) (string, error) {
		src := &source.Source{}
		ps.Data = src
		return render(r, w, as, &views.TODO{Message: "TODO: New source"}, ps, "sources", "New")
	})
}

func SourceEdit(w http.ResponseWriter, r *http.Request) {
	act("source.edit", w, r, func(as *app.State, ps *cutil.PageState) (string, error) {
		key := mux.Vars(r)["key"]
		src, err := as.Sources.Load(key)
		if err != nil {
			return "", errors.Wrap(err, "unable to load source ["+key+"]")
		}
		ps.Data = src
		pcfg := &database.DBParams{}
		if src.Type == schema.OriginPostgres {
			err = util.FromJSON(src.Config, pcfg)
			if err != nil {
				return "", errors.Wrap(err, "can't parse config")
			}
		}
		gcfg := map[string]string{"key": "graphql"}
		page := &vsource.SourceForm{OrigKey: key, Source: src, PostgresConfig: pcfg, GraphQLConfig: gcfg}
		return render(r, w, as, page, ps, "sources", src.Title, "Edit")
	})
}

func SourceInsert(w http.ResponseWriter, r *http.Request) {
	act("source.insert", w, r, func(as *app.State, ps *cutil.PageState) (string, error) {
		key := "TODO"
		return flashAndRedir(true, "saved new source", as.Route("source.detail", "key", key), w, r, ps)
	})
}

func SourceUpdate(w http.ResponseWriter, r *http.Request) {
	act("source.update", w, r, func(as *app.State, ps *cutil.PageState) (string, error) {
		key := mux.Vars(r)["key"]
		msg := fmt.Sprintf("saved source \"%v\"", key)
		return flashAndRedir(true, msg, as.Route("source.detail", "key", key), w, r, ps)
	})
}

func SourceRefresh(w http.ResponseWriter, r *http.Request) {
	act("source.refresh", w, r, func(as *app.State, ps *cutil.PageState) (string, error) {
		key := mux.Vars(r)["key"]
		_, elapsedMillis, err := as.Sources.SchemaRefresh(key)
		if err != nil {
			return "", errors.Wrap(err, "unable to  refresh schema for source ["+key+"]")
		}

		msg := fmt.Sprintf("refreshed in [%.3fms]", elapsedMillis)
		return flashAndRedir(true, msg, as.Route("source.detail", "key", key), w, r, ps)
	})
}
