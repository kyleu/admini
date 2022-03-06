package qualify

import (
	"admini.dev/app/util"
)

type Request struct {
	Type   string        `json:"type"`
	Action string        `json:"action,omitempty"`
	Params util.ValueMap `json:"params,omitempty"`
}

func NewRequest(t string, a string, params ...interface{}) *Request {
	return &Request{Type: t, Action: a, Params: util.ValueMapFor(params...)}
}
