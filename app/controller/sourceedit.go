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

func SourceNew(w http.ResponseWriter, r *http.Request) {
	act("source.new", w, r, func(as *app.State, ps *cutil.PageState) (string, error) {
		ps.Title = "New Source"
		t := schema.OriginPostgres
		s := &source.Source{Type: t}
		ps.Data = s
		return render(r, w, as, &vsource.New{Origin: t}, ps, "sources", "New")
	})
}

func SourceInsert(w http.ResponseWriter, r *http.Request) {
	act("source.insert", w, r, func(as *app.State, ps *cutil.PageState) (string, error) {
		_ = r.ParseForm()
		key := r.Form.Get("key")
		if key == "" {
			return flashAndRedir(false, "key must be provided", as.Route("source.new", "key", key), w, r, ps)
		}
		title := r.Form.Get("title")
		description := r.Form.Get("description")
		typ := r.Form.Get("type")
		ret := &source.Source{Key: key, Title: title, Description: description, Type: schema.OriginFromString(typ)}
		err := currentApp.Sources.Save(ret, false)
		if err != nil {
			return "", errors.Wrap(err, "unable to save source")
		}
		return flashAndRedir(true, "saved new source", as.Route("source.edit", "key", key), w, r, ps)
	})
}

func SourceEdit(w http.ResponseWriter, r *http.Request) {
	act("source.edit", w, r, func(as *app.State, ps *cutil.PageState) (string, error) {
		key := mux.Vars(r)["key"]
		src, err := as.Sources.Load(key, false)
		if err != nil {
			return "", errors.Wrap(err, "unable to load source ["+key+"]")
		}
		ps.Title = "Edit [" + src.Name() + "]"
		ps.Data = src

		switch src.Type {
		case schema.OriginPostgres:
			pcfg := &database.DBParams{}
			if len(src.Config) > 0 {
				err = util.FromJSON(src.Config, pcfg)
				if err != nil {
					return "", errors.Wrap(err, "can't parse postgres config")
				}
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
