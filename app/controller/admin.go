package controller

import (
	"github.com/kyleu/admini/app"
	"github.com/kyleu/admini/app/controller/cutil"
	"github.com/kyleu/admini/app/util"
	"github.com/kyleu/admini/views/vadmin"
	"github.com/pkg/errors"
	"net/http"
	"os"
	"runtime/pprof"
	"strings"
)

func Admin(w http.ResponseWriter, r *http.Request) {
	act("admin", w, r, func(as *app.State, ps *cutil.PageState) (string, error) {
		path := util.SplitAndTrim(strings.TrimPrefix(r.URL.Path, "/admin"), "/")
		if len(path) == 0 {
			return render(r, w, as, &vadmin.List{}, ps)
		}
		switch path[0] {
		case "cpu":
			switch path[1] {
			case "start":
				err := startCPUProfile()
				if err != nil {
					return "", err
				}
				return flashAndRedir(true, "started CPU profile", as.Route("admin"), w, r, ps)
			case "stop":
				pprof.StopCPUProfile()
				return flashAndRedir(true, "stopped CPU profile", as.Route("admin"), w, r, ps)
			default:
				return "", errors.New("unhandled CPU profile action [" + path[1] + "]")
			}
		case "heap":
			err := takeHeapProfile()
			if err != nil {
				return "", err
			}
			return flashAndRedir(true, "wrote heap profile", as.Route("admin"), w, r, ps)
		default:
			return "", errors.New("unhandled admin action [" + path[0] + "]")
		}
	})
}

func startCPUProfile() error {
	f, err := os.Create("./tmp/cpu.pprof")
	if err != nil {
		return err
	}
	return pprof.StartCPUProfile(f)
}

func takeHeapProfile() error {
	f, err := os.Create("./tmp/mem.pprof")
	if err != nil {
		return err
	}
	return pprof.WriteHeapProfile(f)
}
