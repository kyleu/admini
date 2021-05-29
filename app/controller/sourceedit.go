package controller

import (
	"fmt"
	"net/http"
	"strconv"

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
		ret := currentApp.Sources.NewSource(key, title, description, schema.OriginFromString(typ))
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

func SourceSave(w http.ResponseWriter, r *http.Request) {
	act("source.save", w, r, func(as *app.State, ps *cutil.PageState) (string, error) {
		_ = r.ParseForm()
		key := mux.Vars(r)["key"]

		src, err := as.Sources.Load(key, false)
		if err != nil {
			return "", errors.Wrap(err, "unable to load source ["+key+"]")
		}

		src.Title = r.Form.Get("title")
		src.Description = r.Form.Get("description")

		switch src.Type {
		case schema.OriginPostgres:
			p := 0
			ps := r.Form.Get("port")
			if ps != "" {
				p, _ = strconv.Atoi(ps)
			}
			src.Config = util.ToJSONBytes(&database.DBParams{
				Host:     r.Form.Get("host"),
				Port:     p,
				Username: r.Form.Get("username"),
				Password: r.Form.Get("password"),
				Database: r.Form.Get("database"),
				Schema:   r.Form.Get("schema"),
			}, true)
		default:
			return "", errors.New("unable to parse config for source type [" + src.Type.String() + "]")
		}

		err = currentApp.Sources.Save(src, true)
		if err != nil {
			return "", errors.Wrap(err, "unable to save source ["+key+"]")
		}

		msg := fmt.Sprintf(`saved source "%v"`, key)
		return flashAndRedir(true, msg, as.Route("source.detail", "key", key), w, r, ps)
	})
}

func SourceDelete(w http.ResponseWriter, r *http.Request) {
	act("source.delete", w, r, func(as *app.State, ps *cutil.PageState) (string, error) {
		_ = r.ParseForm()
		key := mux.Vars(r)["key"]

		err := as.Sources.Delete(key)
		if err != nil {
			return "", errors.Wrap(err, "unable to delete source ["+key+"]")
		}

		msg := fmt.Sprintf(`deleted source "%v"`, key)
		return flashAndRedir(true, msg, as.Route("source.list"), w, r, ps)
	})
}
