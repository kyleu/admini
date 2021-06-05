package controller

import (
	"fmt"

	"github.com/kyleu/admini/app/util"
	"github.com/valyala/fasthttp"

	"github.com/kyleu/admini/app/controller/cutil"

	"github.com/kyleu/admini/app"
	"github.com/kyleu/admini/views/vhelp"
)

var helpContent = util.ValueMap{
	"_": "help",
	"urls": map[string]string{
		"home": "/",
	},
}

func Feedback(ctx *fasthttp.RequestCtx) {
	act("feedback", ctx, func(as *app.State, ps *cutil.PageState) (string, error) {
		ps.Title = "Feedback"
		ps.Data = fmt.Sprintf("to submit feedback, email %s", util.AppContact)
		return render(ctx, as, &vhelp.Feedback{}, ps, "feedback")
	})
}

func Help(ctx *fasthttp.RequestCtx) {
	act("help", ctx, func(as *app.State, ps *cutil.PageState) (string, error) {
		ps.Title = "Help"
		ps.Data = helpContent
		return render(ctx, as, &vhelp.Help{}, ps, "help")
	})
}
