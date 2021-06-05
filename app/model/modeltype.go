package model

import (
	"encoding/json"

	"github.com/pkg/errors"
)

type Type struct {
	Key    string
	Title  string
	Plural string
}

var (
	TypeEnum         = Type{Key: "enum", Title: "Enum", Plural: "Enums"}
	TypeSequence     = Type{Key: "sequence", Title: "Sequence", Plural: "Sequences"}
	TypeStruct       = Type{Key: "struct", Title: "Struct", Plural: "Structs"}
	TypeInterface    = Type{Key: "interface", Title: "Interface", Plural: "Interfaces"}
	TypeUnion        = Type{Key: "union", Title: "Union", Plural: "Unions"}
	TypeIntersection = Type{Key: "intersection", Title: "Intersection", Plural: "Intersections"}
	TypeUnknown      = Type{Key: "unknown", Title: "Unknown", Plural: "Unknowns"}
)

var AllModelTypes = []Type{
	TypeEnum, TypeSequence, TypeStruct,
	TypeInterface, TypeUnion, TypeIntersection,
}

func modelTypeFromString(s string) (Type, error) {
	for _, t := range AllModelTypes {
		if t.Key == s {
			return t, nil
		}
	}
	return TypeUnknown, errors.New("\"unhandled model type [\" + s + \"]\"")
}

func (t *Type) String() string {
	return t.Key
}

func (t *Type) MarshalJSON() ([]byte, error) {
	return json.Marshal(t.Key)
}

func (t *Type) UnmarshalJSON(data []byte) error {
	var s string
	err := json.Unmarshal(data, &s)
	if err != nil {
		return err
	}
	x, err := modelTypeFromString(s)
	if err != nil {
		return err
	}
	*t = x
	return nil
}
