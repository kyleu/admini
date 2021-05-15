package cutil

import (
	"net/http"
	"sort"
	"strings"

	"github.com/pkg/errors"
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

func (c FormValues) Get(k string) *FormValue {
	for _, x := range c {
		if x.Key == k {
			return x
		}
	}
	return nil
}

const sfx = "--selected"

func (c FormValues) AsChanges() (map[string]interface{}, error) {
	keys := []string{}
	vals := map[string]interface{}{}

	for _, f := range c {
		if strings.HasSuffix(f.Key, sfx) {
			k := strings.TrimSuffix(f.Key, sfx)
			keys = append(keys, k)
		} else {
			curr, ok := vals[f.Key]
			if ok {
				return nil, errors.Errorf("multiple values presented for [%v] (%v/%v)", f.Key, curr, f.Value)
			}
			vals[f.Key] = f.Value
		}
	}

	ret := make(map[string]interface{}, len(keys))
	for _, k := range keys {
		ret[k] = vals[k]
	}
	return ret, nil
}
