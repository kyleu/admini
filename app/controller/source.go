package controller

import (
	"fmt"
	"net/http"

	"github.com/kyleu/admini/app/controller/cutil"

	"github.com/kyleu/admini/app"
	"github.com/kyleu/admini/views/vsource"

	"github.com/gorilla/mux"
)

func SourceList(w http.ResponseWriter, r *http.Request) {
	act("source.list", w, r, func(as *app.State, ps *cutil.PageState) (string, error) {
		s, err := as.Sources.List()
		if err != nil {
			return "", fmt.Errorf("unable to load source list: %w", err)
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
			return "", fmt.Errorf("unable to load source ["+key+"]: %w", err)
		}
		sch, err := as.Sources.SchemaFor(key)
		if err != nil {
			return "", fmt.Errorf("unable to load schema for source ["+key+"]: %w", err)
		}
		ps.Data = map[string]interface{}{"source": src, "schema": sch}
		return render(r, w, as, &vsource.SourceDetail{Source: src, Schema: sch}, ps, "sources", src.Key)
	})
}

func SourceRefresh(w http.ResponseWriter, r *http.Request) {
	act("source.refresh", w, r, func(as *app.State, ps *cutil.PageState) (string, error) {
		key := mux.Vars(r)["key"]
		_, elapsedMillis, err := as.Sources.SchemaRefresh(key)
		if err != nil {
			return "", fmt.Errorf("unable to  refresh schema for source ["+key+"]: %w", err)
		}

		msg := fmt.Sprintf("refreshed in [%.3fms]", elapsedMillis)
		return flashAndRedir(true, msg, as.Route("source.detail", "key", key), w, r, ps)
	})
}
