package util

import (
	"fmt"
	"strings"

	"github.com/pkg/errors"
)

type ValueMap map[string]interface{}

func ValueMapFor(kvs ...interface{}) ValueMap {
	ret := make(ValueMap, len(kvs)/2)
	for i := 0; i < len(kvs); i += 2 {
		k, ok := kvs[i].(string)
		if !ok {
			k = fmt.Sprintf("error-invalid-type:%T", kvs[i])
		}
		ret[k] = kvs[i+1]
	}
	return ret
}

func (c ValueMap) GetRequired(k string) (interface{}, error) {
	v, ok := c[k]
	if !ok {
		msg := "no form value [%v] among candidates [%v]"
		return nil, errors.Errorf(msg, k, strings.Join(c.Keys(), ", "))
	}
	return v, nil
}

func (c ValueMap) GetString(k string, allowEmpty bool) (string, error) {
	v, err := c.GetRequired(k)
	if err != nil {
		return "", err
	}

	ret := ""
	switch t := v.(type) {
	case []string:
		ret = strings.Join(t, "|")
	case string:
		ret = t
	default:
		return "", errors.Errorf("expected string or array of strings, encountered %T: %v", t, t)
	}
	if !allowEmpty && ret == "" {
		return "", errors.Errorf("field [%v] may not be empty", k)
	}
	return ret, nil
}

func (c ValueMap) GetStringOpt(k string) string {
	ret, _ := c.GetString(k, true)
	return ret
}

func (c ValueMap) GetStringArray(k string, allowMissing bool) ([]string, error) {
	v, err := c.GetRequired(k)
	if err != nil {
		if allowMissing {
			return nil, nil
		}
		return nil, err
	}

	switch t := v.(type) {
	case []string:
		return t, nil
	default:
		return nil, errors.Errorf("expected array of strings, encountered %T: %v", t, t)
	}
}

const selectedSuffix = "--selected"

func (c ValueMap) AsChanges() (ValueMap, error) {
	keys := []string{}
	vals := ValueMap{}

	for k, v := range c {
		if strings.HasSuffix(k, selectedSuffix) {
			key := strings.TrimSuffix(k, selectedSuffix)
			keys = append(keys, key)
		} else {
			curr, ok := vals[k]
			if ok {
				return nil, errors.Errorf("multiple values presented for [%v] (%v/%v)", k, curr, v)
			}
			vals[k] = v
		}
	}

	ret := make(ValueMap, len(keys))
	for _, k := range keys {
		ret[k] = vals[k]
	}
	return ret, nil
}

func (c ValueMap) Keys() []string {
	ret := make([]string, 0, len(c))
	for k := range c {
		ret = append(ret, k)
	}
	return ret
}