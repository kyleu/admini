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

func Qualify(req *Request, acts Actions) ([]interface{}, error) {
	var ret []interface{}
	for _, act := range acts {
		childResult, err := qualifyAct(req, act)
		if err != nil {
			return nil, err
		}
		ret = append(ret, childResult...)
	}
	return ret, nil
}

func qualifyAct(req *Request, act *Action) ([]interface{}, error) {
	var ret []interface{}

	switch act.Type {
	case TypeFolder:
	case TypeStatic:
	case TypeSeparator:

	case TypeAll:
		if req.Type == TypeModel.Key {
			ret = append(ret, "All!!!")
		}
	case TypeSource:
		if req.Type == TypeModel.Key {
			ret = append(ret, "Source!!!")
		}
	case TypePackage:
		if req.Type == TypeModel.Key {
			ret = append(ret, "Package!!!")
		}
	case TypeModel:
		if req.Type == TypeModel.Key {
			ret = append(ret, "Model!!!")
		}
	case TypeActivity:
		ret = append(ret, "Activity!!!")

	default:
		return nil, errors.New("unhandled action type [" + act.Type.Key + "] in qualify attempt")
	}
	kids, err := Qualify(req, act.Children)
	if err != nil {
		return nil, err
	}
	ret = append(ret, kids...)
	return ret, nil
}
