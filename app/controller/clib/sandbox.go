package clib

import (
	"fmt"
	"net/http"

	"admini.dev/admini/app"
	"admini.dev/admini/app/controller"
	"admini.dev/admini/app/controller/cutil"
	"admini.dev/admini/app/lib/sandbox"
	"admini.dev/admini/app/lib/telemetry"
	"admini.dev/admini/app/util"
	"admini.dev/admini/views"
	"admini.dev/admini/views/vpage"
	"admini.dev/admini/views/vsandbox"
)

func SandboxList(w http.ResponseWriter, r *http.Request) {
	controller.Act("sandbox.list", w, r, func(as *app.State, ps *cutil.PageState) (string, error) {
		if title := r.URL.Query().Get("title"); title != "" {
			ps.SetTitleAndData(title, title)
			return controller.Render(r, as, &views.Debug{}, ps, title)
		}
		ps.SetTitleAndData("Sandboxes", sandbox.AllSandboxes)
		return controller.Render(r, as, &vsandbox.List{}, ps, "sandbox")
	})
}

func SandboxRun(w http.ResponseWriter, r *http.Request) {
	controller.Act("sandbox.run", w, r, func(as *app.State, ps *cutil.PageState) (string, error) {
		key, err := cutil.PathString(r, "key", false)
		if err != nil {
			return "", err
		}

		sb := sandbox.AllSandboxes.Get(key)
		if sb == nil {
			return controller.ERsp("no sandbox with key [%s]", key)
		}

		argRes := util.FieldDescsCollect(r, sb.Args)
		if argRes.HasMissing() {
			ps.Data = argRes
			url := fmt.Sprintf("/admin/sandbox/%s", sb.Key)
			return controller.Render(r, as, &vpage.Args{URL: url, Directions: "Choose your options", Results: argRes}, ps, "sandbox", sb.Key)
		}

		ctx, span, logger := telemetry.StartSpan(ps.Context, "sandbox."+key, ps.Logger)
		defer span.Complete()

		ret, err := sb.Run(ctx, as, argRes.Values, logger.With("sandbox", key))
		if err != nil {
			return "", err
		}
		ps.SetTitleAndData(sb.Title, ret)
		if sb.Key == "testbed" {
			return controller.Render(r, as, &vsandbox.Testbed{}, ps, "sandbox", sb.Key)
		}
		return controller.Render(r, as, &vsandbox.Run{Key: key, Title: sb.Title, Icon: sb.Icon, Result: ret}, ps, "sandbox", sb.Key)
	})
}
