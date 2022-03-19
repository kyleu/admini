package result

import (
	"fmt"
	"reflect"

	"github.com/pkg/errors"

	"admini.dev/admini/app/lib/schema/field"
)

func FromReflection(title string, t ...any) (*Result, error) {
	if len(t) == 0 {
		return nil, errors.New("empty input when building result")
	}
	first := t[0]
	q := fmt.Sprintf("built from [%T]", first)

	fields, err := getFields(first)
	if err != nil {
		return nil, errors.Wrapf(err, "unable to calculate fields for [%T]", first)
	}

	data := make([][]any, 0, len(t))
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

func getValues(x any) ([]any, error) {
	return valuesOf(reflect.ValueOf(x))
}

func valuesOf(v reflect.Value) ([]any, error) {
	if v.Kind() == reflect.Ptr {
		return valuesOf(v.Elem())
	}

	t := v.Type()

	ret := make([]any, 0, t.NumField())
	for i := 0; i < t.NumField(); i++ {
		f := v.Field(i)
		ret = append(ret, f.Interface())
	}
	return ret, nil
}

func getFields(x any) (field.Fields, error) {
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
