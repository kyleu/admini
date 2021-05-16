package result

import (
	"fmt"
	"github.com/kyleu/admini/app/field"
	"github.com/pkg/errors"
	"reflect"
)

func ResultFromReflection(title string, t ...interface{}) (*Result, error) {
	if len(t) == 0 {
		return nil, errors.New("empty input when building result")
	}
	first := t[0]
	q := fmt.Sprintf("built from [%T]", first)

	fields, err := getFields(first)
	if err != nil {
		return nil, errors.Wrapf(err, "unable to calculate fields for [%T]", first)
	}

	data := make([][]interface{}, 0, len(t))
	for _, x := range t {
		v, err := getValues(x)
		if err != nil {
			return nil, errors.Wrap(err, "")
		}
		data = append(data, v)
	}

	timing := &Timing{}

	ret := &Result{Title: title, Count: len(t), Query: q, Fields: fields, Data: data, Timing: timing}
	return ret, nil
}

func getValues(x interface{}) ([]interface{}, error) {
	return valuesOf(reflect.ValueOf(x))
}

func valuesOf(v reflect.Value) ([]interface{}, error) {
	if v.Kind() == reflect.Ptr {
		return valuesOf(v.Elem())
	}

	t := v.Type()

	ret := make([]interface{}, 0, t.NumField())
	for i := 0; i < t.NumField(); i++ {
		f := v.Field(i)
		ret = append(ret, f.Interface())
	}
	return ret, nil
}

func getFields(x interface{}) (field.Fields, error) {
	return fieldsOf(reflect.ValueOf(x))
}

func fieldsOf(v reflect.Value) (field.Fields, error) {
	if v.Kind() == reflect.Ptr {
		return fieldsOf(v.Elem())
	}

	t := v.Type()

	ret := make(field.Fields, 0, t.NumField())
	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		ret = append(ret, field.NewFieldByType(f.Name, f.Type, !v.CanSet(), nil))
	}

	return ret, nil
}
