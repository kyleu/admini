package controller

import (
	"net/http"

	"github.com/kyleu/admini/app/util"

	"github.com/kyleu/admini/app/controller/cutil"

	"github.com/kyleu/admini/app"
	"github.com/kyleu/admini/views/vhelp"
)

var helpContent = map[string]interface{}{
	"_": "help",
	"urls": map[string]string{
		"TODO": "http://TODO",
	},
}

func Feedback(w http.ResponseWriter, r *http.Request) {
	act("feedback", w, r, func(as *app.State, ps *cutil.PageState) (string, error) {
		ps.Title = "Feedback"
		ps.Data = "to submit feedback, email " + util.AppContact
		return render(r, w, as, &vhelp.Feedback{}, ps, "feedback")
	})
}

func Help(w http.ResponseWriter, r *http.Request) {
	act("help", w, r, func(as *app.State, ps *cutil.PageState) (string, error) {
		ps.Title = "Help"
		ps.Data = helpContent
		return render(r, w, as, &vhelp.Help{}, ps, "help")
	})
}
