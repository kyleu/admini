package qualify

import (
	"fmt"
	"github.com/kyleu/admini/app/action"
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
	Icon   string   `json:"icon,omitempty"`
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

func Qualify(req *Request, acts action.Actions, schemata schema.Schemata) (Results, error) {
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

func qualifyAct(req *Request, act *action.Action, schemata schema.Schemata) (Results, error) {
	var ret Results

	switch req.Type {
	case action.TypeModel.Key:
		srcKey, err := req.Params.GetString(action.TypeSource.Key, false)
		if err != nil {
			return nil, err
		}
		modelPath, err := req.Params.GetString(action.TypeModel.Key, false)
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
