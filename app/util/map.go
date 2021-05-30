package util

import (
	"sort"
	"strings"

	"github.com/pkg/errors"
)

type ValueMapEntry struct {
	Key   string      `json:"key"`
	Value interface{} `json:"value"`
}

type ValueMap []*ValueMapEntry

func ValueMapFor(kvs ...interface{}) ValueMap {
	ret := make(ValueMap, 0, len(kvs) / 2)
	for i := 0; i < len(kvs); i += 2 {
		ret = append(ret, &ValueMapEntry{Key: kvs[i].(string), Value: kvs[i+1]})
	}
	return ret
}

func ValueMapFrom(m map[string]interface{}) ValueMap {
	ret := make(ValueMap, 0, len(m))
	for k, v := range m {
		ret = append(ret, &ValueMapEntry{Key: k, Value: v})
	}
	return ret
}

func (c ValueMap) Get(k string) *ValueMapEntry {
	for _, x := range c {
		if x.Key == k {
			return x
		}
	}
	return nil
}

func (c ValueMap) ToMap() map[string]interface{} {
	ret := make(map[string]interface{}, len(c))
	for _, v := range c {
		ret[v.Key] = v.Value
	}
	return ret
}

func (c ValueMap) GetRequired(k string) (*ValueMapEntry, error) {
	fv := c.Get(k)
	if fv == nil {
		msg := "no form value [%v] among candidates [%v]"
		return nil, errors.Errorf(msg, k, strings.Join(c.Keys(), ", "))
	}
	return fv, nil
}

func (c ValueMap) GetString(k string, allowEmpty bool) (string, error) {
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
		return "", errors.Errorf("field [%v] may not be empty", fv.Key)
	}
	return ret, nil
}

func (c ValueMap) GetStringOpt(k string) string {
	ret, _ := c.GetString(k, true)
	return ret
}
func (c ValueMap) GetStringArray(k string, allowMissing bool) ([]string, error) {
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

func (c ValueMap) Sort() {
	sort.Slice(c, func(i, j int) bool {
		return c[i].Key < c[j].Key
	})
}

const selectedSuffix = "--selected"

func (c ValueMap) AsChanges() (map[string]interface{}, error) {
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

func (c ValueMap) Keys() []string {
	ret := make([]string, 0, len(c))
	for _, fv := range c {
		ret = append(ret, fv.Key)
	}
	return ret
}

func (c ValueMap) Set(k string, v interface{}) ValueMap {
	for _, v := range c {
		if v.Key == k {
			v.Value = v
			return c
		}
	}
	c = append(c, &ValueMapEntry{Key: k, Value: v})
	return c
}
