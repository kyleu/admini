package cutil

import (
	"github.com/pkg/errors"
	"net/http"
	"sort"
	"strings"
)

type FormValue struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type FormValues []*FormValue

func ParseForm(req *http.Request) (FormValues, error) {
	if err := req.ParseForm(); err != nil {
		return nil, errors.Wrap(err, "can't parse form")
	}

	frm := make(map[string]interface{}, len(req.Form))
	for k, v := range req.Form {
		frm[k] = strings.Join(v, "||")
	}

	ret := make(FormValues, 0, len(frm))
	for k, v := range req.Form {
		ret = append(ret, &FormValue{Key: k, Value: strings.Join(v, "||")})
	}
	sort.Slice(ret, func(i, j int) bool {
		return ret[i].Key < ret[j].Key
	})

	return ret, nil
}
