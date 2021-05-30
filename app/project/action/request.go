package action

import (
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
}

type Results []*Result

func Qualify(req *Request, acts Actions) (Results, error) {
	var ret Results
	for _, act := range acts {
		childResult, err := qualifyAct(req, act)
		if err != nil {
			return nil, err
		}
		ret = append(ret, childResult...)
	}
	return ret, nil
}

func qualifyAct(req *Request, act *Action) (Results, error) {
	var ret Results

	switch req.Type {
	case TypeModel.Key:
		modelPath, err := req.Params.GetString(TypeModel.Key, false)
		if err != nil {
			return nil, err
		}
		srcKey, err := req.Params.GetString(TypeSource.Key, false)
		if err != nil {
			return nil, err
		}
		ret, err = qualifyModel(req, act, srcKey, modelPath)
		if err != nil {
			return nil, err
		}
	default:
		return nil, errors.New("unhandled qualify type [" + req.Type + "]")
	}

	kids, err := Qualify(req, act.Children)
	if err != nil {
		return nil, err
	}
	ret = append(ret, kids...)
	return ret, nil
}

func qualifyModel(req *Request, act *Action, srcKey string, modelPath string) (Results, error) {
	p := util.SplitAndTrim(modelPath, "/")
	switch act.Type {
	case TypeAll:
		return Results{{Action: act.Path(), Path: p}}, nil
	case TypeSource:
		return Results{{Action: act.Path(), Path: p}}, nil
	case TypePackage:
		return Results{{Action: act.Path(), Path: p}}, nil
	case TypeModel:
		return Results{{Action: act.Path(), Path: p}}, nil
	}
	return nil, nil
}
