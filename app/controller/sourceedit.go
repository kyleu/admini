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
		frm, err := cutil.ParseForm(r)
		if err != nil {
			return "", errors.Wrap(err, "unable to parse form")
		}

		key, err := frm.GetString("key", false)
		if err != nil {
			return flashError(err, as.Route("source.new", "key", key), w, r, ps)
		}
		title, err := frm.GetString("title", true)
		if err != nil {
			return "", err
		}
		description, err := frm.GetString("description", true)
		if err != nil {
			return "", err
		}
		typ, err := frm.GetString("type", true)
		if err != nil {
			return "", err
		}
		ret := currentApp.Sources.NewSource(key, title, description, schema.OriginFromString(typ))
		err = currentApp.Sources.Save(ret, false)
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
		frm, err := cutil.ParseForm(r)
		if err != nil {
			return "", errors.Wrap(err, "unable to parse form")
		}

		key := mux.Vars(r)["key"]

		src, err := as.Sources.Load(key, false)
		if err != nil {
			return "", errors.Wrap(err, "unable to load source ["+key+"]")
		}

		src.Title, err = frm.GetString("title", true)
		if err != nil {
			return "", err
		}
		src.Description, err = frm.GetString("description", true)
		if err != nil {
			return "", err
		}

		switch src.Type {
		case schema.OriginPostgres:
			ps, _ := frm.GetString("port", true)
			params := &database.DBParams{}
			params.Host, err = frm.GetString("host", false)
			if err != nil {
				return "", err
			}
			if ps != "" {
				params.Port, _ = strconv.Atoi(ps)
			}
			params.Username, _ = frm.GetString("username", true)
			params.Password, _ = frm.GetString("password", true)
			params.Database, _ = frm.GetString("database", true)
			params.Schema, _ = frm.GetString("schema", true)

			src.Config = util.ToJSONBytes(params, true)
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
		key := mux.Vars(r)["key"]

		err := as.Sources.Delete(key)
		if err != nil {
			return "", errors.Wrap(err, "unable to delete source ["+key+"]")
		}

		msg := fmt.Sprintf(`deleted source "%v"`, key)
		return flashAndRedir(true, msg, as.Route("source.list"), w, r, ps)
	})
}
