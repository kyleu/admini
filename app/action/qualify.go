package action

import (
	"fmt"
	"strings"

	"github.com/kyleu/admini/app/schema"
	"github.com/kyleu/admini/app/util"
	"github.com/pkg/errors"
)

type Request struct {
	Type   string        `json:"type"`
	Action string        `json:"action,omitempty"`
	Params util.ValueMap `json:"params,omitempty"`
}

func NewRequest(t string, a string, params ...interface{}) *Request {
	return &Request{Type: t, Action: a, Params: util.ValueMapFor(params...)}
}

type Result struct {
	Action []string `json:"act"`
	Path   []string `json:"path"`
	Debug  string   `json:"debug,omitempty"`
}

func (r *Result) String() string {
	if r.Debug == "" {
		return strings.Join(r.Link(), "/")
	}
	return fmt.Sprintf("%s (%s)", strings.Join(r.Link(), "/"), r.Debug)
}

func (r *Result) Link() []string {
	return append(r.Action, r.Path...)
}

type Results []*Result

func Qualify(req *Request, acts Actions, schemata schema.Schemata) (Results, error) {
	var ret Results
	for _, act := range acts {
		childResult, err := qualifyAct(req, act, schemata)
		if err != nil {
			return nil, err
		}
		ret = append(ret, childResult...)
	}
	return ret, nil
}

func qualifyAct(req *Request, act *Action, schemata schema.Schemata) (Results, error) {
	var ret Results

	switch req.Type {
	case TypeModel.Key:
		srcKey, err := req.Params.GetString(TypeSource.Key, false)
		if err != nil {
			return nil, err
		}
		modelPath, err := req.Params.GetString(TypeModel.Key, false)
		if err != nil {
			return nil, err
		}
		ret, err = qualifyModel(req, act, srcKey, util.SplitAndTrim(modelPath, "/"), schemata)
		if err != nil {
			return nil, err
		}
	default:
		return nil, errors.Errorf("unhandled qualify type [%s]", req.Type)
	}

	kids, err := Qualify(req, act.Children, schemata)
	if err != nil {
		return nil, err
	}
	ret = append(ret, kids...)
	return ret, nil
}

func qualifyModel(req *Request, act *Action, srcKey string, modelPath util.Pkg, schemata schema.Schemata) (Results, error) {
	if act.Type == TypeAll {
		sch, err := schemata.GetWithError(srcKey)
		if err != nil {
			return nil, errors.Wrapf(err, "can't find source in project with key [%s]", modelPath[0])
		}
		src := sch.ModelsByPackage()
		m, _ := src.Get(modelPath)
		if m != nil {
			return modelResults(req, act.Path(), append([]string{srcKey}, modelPath...), "all")
		}
		return nil, nil
	}

	if !checkSource(act, srcKey) {
		return nil, nil
	}
	switch act.Type {
	case TypeSource:
		src, err := getSource(act, schemata)
		if err != nil {
			return nil, err
		}
		m, _ := src.Get(modelPath)
		if m != nil {
			return modelResults(req, act.Path(), modelPath, "src")
		}
	case TypePackage:
		pkg, remaining, err := getPackage(act, schemata)
		if err != nil {
			return nil, err
		}
		if modelPath.StartsWith(pkg.Path()) {
			m, _ := pkg.Get(remaining)
			if m != nil {
				return modelResults(req, act.Path(), modelPath[len(pkg.Path()):], "pkg")
			}
		}
	case TypeModel:
		model, remaining, err := getModel(act, schemata)
		if err != nil {
			return nil, err
		}

		if modelPath.StartsWith(model.Path()) {
			return modelResults(req, act.Path(), remaining, "model")
		}
	}
	return nil, nil
}

func modelResults(req *Request, actPath []string, remaining []string, dbg string) (Results, error) {
	path := append([]string{}, remaining...)
	if req.Action == "view" {
		keys, err := req.Params.GetStringArray("keys", false)
		if err != nil {
			return nil, err
		}
		path = append(append(path, "v"), keys...)
	}
	return Results{{Action: actPath, Path: path, Debug: dbg}}, nil
}

func checkSource(act *Action, key string) bool {
	k, err := act.Config.GetString("source", false)
	if err != nil {
		return false
	}
	return k == key
}
