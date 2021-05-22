package controller

import (
	"fmt"
	"net/http"

	"github.com/kyleu/admini/views"

	"github.com/kyleu/admini/app/database"
	"github.com/kyleu/admini/app/schema"
	"github.com/kyleu/admini/app/source"
	"github.com/kyleu/admini/app/util"
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
		ps.Title = "Sources"
		ps.Data = s
		return render(r, w, as, &vsource.List{Sources: s}, ps, "sources")
	})
}

func SourceNew(w http.ResponseWriter, r *http.Request) {
	act("source.new", w, r, func(as *app.State, ps *cutil.PageState) (string, error) {
		ps.Title = "New Source"
		ps.Data = &source.Source{}
		return render(r, w, as, &vsource.New{}, ps, "sources", "New")
	})
}

func SourceInsert(w http.ResponseWriter, r *http.Request) {
	act("source.insert", w, r, func(as *app.State, ps *cutil.PageState) (string, error) {
		key := "TODO"
		return flashAndRedir(true, "saved new source", as.Route("source.detail", "key", key), w, r, ps)
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
		ps.Title = src.Name()
		ps.Data = map[string]interface{}{"source": src, "schema": sch}
		return render(r, w, as, &vsource.Detail{Source: src, Schema: sch}, ps, "sources", src.Key)
	})
}

func SourceEdit(w http.ResponseWriter, r *http.Request) {
	act("source.edit", w, r, func(as *app.State, ps *cutil.PageState) (string, error) {
		key := mux.Vars(r)["key"]
		src, err := as.Sources.Load(key)
		if err != nil {
			return "", errors.Wrap(err, "unable to load source ["+key+"]")
		}
		ps.Title = "Edit [" + src.Name() + "]"
		ps.Data = src

		switch src.Type {
		case schema.OriginPostgres:
			pcfg := &database.DBParams{}
			err = util.FromJSON(src.Config, pcfg)
			if err != nil {
				return "", errors.Wrap(err, "can't parse postgres config")
			}
			return render(r, w, as, &vsource.EditPostgres{Source: src, Cfg: pcfg}, ps, "sources", src.Key, "Edit")
		default:
			return render(r, w, as, &views.TODO{Message: "unhandled source type [" + src.Type.String() + "]"}, ps, "sources", src.Key, "Edit")
		}
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
