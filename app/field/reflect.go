package field

import (
	"fmt"
	"github.com/kyleu/admini/app/schema/schematypes"
	"reflect"
)

func NewFieldByType(key string, t reflect.Type, ro bool, md *Metadata) *Field {
	return &Field{Key: key, Type: fromReflect(t), ReadOnly: ro, Metadata: md}
}

func fromReflect(t reflect.Type) *schematypes.Wrapped {
	switch t.Kind() {
	case reflect.Invalid:
		return schematypes.NewError("can't reflect invalid")
	case reflect.Bool:
		return schematypes.NewBool()
	case reflect.Int:
		return schematypes.NewInt(0)
	case reflect.Int8:
		return schematypes.NewInt(8)
	case reflect.Int16:
		return schematypes.NewInt(16)
	case reflect.Int32:
		return schematypes.NewInt(32)
	case reflect.Int64:
		return schematypes.NewInt(64)
	case reflect.Uint:
		return schematypes.NewUnsignedInt(0)
	case reflect.Uint8:
		return schematypes.NewUnsignedInt(8)
	case reflect.Uint16:
		return schematypes.NewUnsignedInt(16)
	case reflect.Uint32:
		return schematypes.NewUnsignedInt(32)
	case reflect.Uint64:
		return schematypes.NewUnsignedInt(64)
	case reflect.Uintptr:
		return schematypes.NewError("can't reflect uint ponters")
	case reflect.Float32:
		return schematypes.NewFloat(32)
	case reflect.Float64:
		return schematypes.NewFloat(64)
	case reflect.Complex64:
		return schematypes.NewError("can't reflect complex")
	case reflect.Complex128:
		return schematypes.NewError("can't reflect complex")
	case reflect.Array:
		return schematypes.NewList(fromReflect(t.Elem()))
	case reflect.Chan:
		return schematypes.NewError("can't reflect channels")
	case reflect.Func:
		return schematypes.NewError("can't reflect functions")
	case reflect.Interface:
		return schematypes.NewError("can't reflect interfaces")
	case reflect.Map:
		return schematypes.NewMap(fromReflect(t.Key()), fromReflect(t.Elem()))
	case reflect.Ptr:
		return schematypes.NewOption(fromReflect(t.Elem()))
	case reflect.Slice:
		return schematypes.NewList(fromReflect(t.Elem()))
	case reflect.String:
		return schematypes.NewString()
	case reflect.Struct:
		return schematypes.NewError("can't reflect structs")
	case reflect.UnsafePointer:
		return schematypes.NewError("can't reflect unsafe pointers")
	default:
	  return schematypes.NewUnknown(fmt.Sprintf("%v", t.Kind()))
	}
}
