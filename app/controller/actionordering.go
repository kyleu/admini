package controller

import (
	"fmt"
	"time"

	"github.com/pkg/errors"
	"github.com/valyala/fasthttp"

	"admini.dev/admini/app"
	"admini.dev/admini/app/action"
	"admini.dev/admini/app/controller/cutil"
	"admini.dev/admini/app/util"
)

func ActionOrdering(rc *fasthttp.RequestCtx) {
	act("action.ordering", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		key, err := cutil.RCRequiredString(rc, "key", false)
		if err != nil {
			return "", err
		}
		prj, err := as.Services.Projects.LoadRequired(key, false)
		if err != nil {
			return "", errors.Wrapf(err, "unable to load project [%s]", key)
		}

		frm, err := cutil.ParseForm(rc)
		if err != nil {
			return "", errors.Wrap(err, "unable to parse form")
		}
		ordering := frm.GetStringOpt("ordering")

		actOrders := action.Orderings{}
		err = util.FromJSON([]byte(ordering), &actOrders)
		if err != nil {
			return "", err
		}
		startNanos := time.Now().UnixNano()
		prj.Actions, err = action.ReorderActions(prj.Actions, actOrders)
		if err != nil {
			return "", err
		}

		err = as.Services.Projects.Save(prj, true)
		if err != nil {
			return "", err
		}
		elapsedMillis := float64((time.Now().UnixNano()-startNanos)/int64(time.Microsecond)) / float64(1000)

		_, err = as.Services.Projects.LoadRequired(prj.Key, true)
		if err != nil {
			return "", err
		}
		count := prj.Actions.Size()
		msg := fmt.Sprintf("saved [%d] %s in [%.3fms]", count, util.StringPlural(count, "action"), elapsedMillis)
		return flashAndRedir(true, msg, fmt.Sprintf("/project/%s", key), rc, ps)
	})
}
