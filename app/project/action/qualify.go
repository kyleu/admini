package action

import (
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
		return nil, errors.New("unhandled qualify type [" + req.Type + "]")
	}

	kids, err := Qualify(req, act.Children, schemata)
	if err != nil {
		return nil, err
	}
	ret = append(ret, kids...)
	return ret, nil
}

func qualifyModel(req *Request, act *Action, srcKey string, modelPath util.Pkg, schemata schema.Schemata) (Results, error) {
	if !checkSource(act, srcKey) {
		return nil, nil
	}
	switch act.Type {
	case TypeAll:
		//return Results{{Action: act.Path(), Path: p}}, nil
	case TypeSource:
		src, err := getSource(act, schemata)
		if err != nil {
			return nil, err
		}
		m, _ := src.Get(modelPath)
		if m != nil {
			return Results{{Action: act.Path(), Path: modelPath, Debug: "src"}}, nil
		}
		//return Results{{Action: act.Path(), Path: p}}, nil
	case TypePackage:
		pkg, remaining, err := getPackage(act, schemata)
		if err != nil {
			return nil, err
		}
		if modelPath.StartsWith(pkg.Path()) {
			m, _ := pkg.Get(remaining)
			if m != nil {
				return Results{{Action: act.Path(), Path: modelPath[len(pkg.Path()):], Debug: "pkg"}}, nil
			}
		}
	case TypeModel:
		model, remaining, err := getModel(act, schemata)
		if err != nil {
			return nil, err
		}

		if modelPath.StartsWith(model.Path()) {
			return Results{{Action: act.Path(), Path: remaining, Debug: "model"}}, nil
		}
	}
	return nil, nil
}

func checkSource(act *Action, key string) bool {
	k, err := act.Config.GetString("source", false)
	if err != nil {
		return false
	}
	return k == key
}
