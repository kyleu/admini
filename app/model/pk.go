package model

import (
	"fmt"

	"github.com/kyleu/admini/app/field"
	"github.com/kyleu/admini/app/util"
)

func (m *Model) GetPK() []string {
	if m.pk == nil {
		for _, idx := range m.Indexes {
			if idx.Primary {
				if m.pk != nil {
					util.LogError("multiple primary keys?!")
				}
				m.pk = idx.Fields
			}
		}
	}
	return m.pk
}

func (m *Model) IsPK(key string) bool {
	pk := m.GetPK()
	for _, col := range pk {
		if col == key {
			return true
		}
	}
	return false
}

func GetValues(src field.Fields, tgt []string, vals []interface{}) ([]interface{}, error) {
	if len(src) != len(vals) {
		return nil, fmt.Errorf("[%d] fields provided, but [%d] values provided", len(src), len(vals))
	}
	ret := make([]interface{}, 0, len(tgt))
	for _, t := range tgt {
		for idx, f := range src {
			if f.Key == t {
				ret = append(ret, vals[idx])
				break
			}
		}
	}
	return ret, nil
}

func GetStrings(src field.Fields, tgt []string, vals []interface{}) ([]string, error) {
	is, err := GetValues(src, tgt, vals)
	if err != nil {
		return nil, err
	}
	ret := make([]string, 0, len(is))
	for _, x := range is {
		ret = append(ret, fmt.Sprintf("%v", x))
	}
	return ret, nil
}
