package cutil

import (
	"github.com/kyleu/admini/app/util"
	"net/http"
	"sort"
	"strings"

	"github.com/pkg/errors"
)

func ParseForm(req *http.Request) (FormValues, error) {
	ct := GetContentType(req)
	if IsContentTypeJSON(ct) {
		return parseJSONForm(req)
	}
	return parseHTTPForm(req)
}

func parseJSONForm(req *http.Request) (FormValues, error) {
	m := map[string]interface{}{}
	err := util.FromJSONReader(req.Body, &m)
	if err != nil {
		return nil, errors.Wrap(err, "can't parse JSON body")
	}
	ret := make(FormValues, 0, len(m))
	for k, v := range m {
		ret = append(ret, &FormValue{Key: k, Value: v})
	}
	ret.Sort()
	return ret, nil
}

func parseHTTPForm(req *http.Request) (FormValues, error) {
	if err := req.ParseForm(); err != nil {
		return nil, errors.Wrap(err, "can't parse form")
	}

	ret := make(FormValues, 0, len(req.Form))
	for k, v := range req.Form {
		ret = append(ret, &FormValue{Key: k, Value: v})
	}
	ret.Sort()
	return ret, nil
}

type FormValue struct {
	Key   string      `json:"key"`
	Value interface{} `json:"value"`
}

type FormValues []*FormValue

func (c FormValues) Get(k string) *FormValue {
	for _, x := range c {
		if x.Key == k {
			return x
		}
	}
	return nil
}

func (c FormValues) GetRequired(k string) (*FormValue, error) {
	fv := c.Get(k)
	if fv == nil {
		msg := "no form value [%v] among candidates [%v]"
		return nil, errors.Errorf(msg, k, strings.Join(c.Keys(), ", "))
	}
	return fv, nil
}

func (c FormValues) GetString(k string, allowEmpty bool) (string, error) {
	fv, err := c.GetRequired(k)
	if err != nil {
		return "", err
	}

	ret := ""
	switch t := fv.Value.(type) {
	case []string:
		ret = strings.Join(t, "|")
	case string:
		ret = t
	default:
		return "", errors.Errorf("unhandled field value type %T: %v", t, t)
	}
	if !allowEmpty && ret == "" {
		return "", errors.Errorf("field [%v] is required", fv.Key)
	}
	return ret, nil
}

func (c FormValues) GetStringArray(k string, allowMissing bool) ([]string, error) {
	fv, err := c.GetRequired(k)
	if err != nil {
		if allowMissing {
			return nil, nil
		}
		return nil, err
	}

	switch t := fv.Value.(type) {
	case []string:
		return t, nil
	default:
		return nil, errors.Errorf("unhandled field value type %T: %v", t, t)
	}
}

func (c FormValues) Sort() {
	sort.Slice(c, func(i, j int) bool {
		return c[i].Key < c[j].Key
	})
}

const selectedSuffix = "--selected"

func (c FormValues) AsChanges() (map[string]interface{}, error) {
	keys := []string{}
	vals := map[string]interface{}{}

	for _, f := range c {
		if strings.HasSuffix(f.Key, selectedSuffix) {
			k := strings.TrimSuffix(f.Key, selectedSuffix)
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

func (c FormValues) Keys() []string {
	ret := make([]string, 0, len(c))
	for _, fv := range c {
		ret = append(ret, fv.Key)
	}
	return ret
}
