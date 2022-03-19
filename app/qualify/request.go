package qualify

import (
	"admini.dev/admini/app/util"
)

type Request struct {
	Type   string        `json:"type"`
	Action string        `json:"action,omitempty"`
	Params util.ValueMap `json:"params,omitempty"`
}

func NewRequest(t string, a string, params ...any) *Request {
	return &Request{Type: t, Action: a, Params: util.ValueMapFor(params...)}
}
