package model

import (
	"encoding/json"

	"github.com/kyleu/admini/app/util"
)

type Type struct {
	Key    string
	Title  string
	Plural string
}

var (
	ModelTypeEnum         = Type{Key: "enum", Title: "Enum", Plural: "Enums"}
	ModelTypeSequence     = Type{Key: "sequence", Title: "Sequence", Plural: "Sequences"}
	ModelTypeStruct       = Type{Key: "struct", Title: "Struct", Plural: "Structs"}
	ModelTypeInterface    = Type{Key: "interface", Title: "Interface", Plural: "Interfaces"}
	ModelTypeUnion        = Type{Key: "union", Title: "Union", Plural: "Unions"}
	ModelTypeIntersection = Type{Key: "intersection", Title: "Intersection", Plural: "Intersections"}
	ModelTypeUnknown      = Type{Key: "unknown", Title: "Unknown", Plural: "Unknowns"}
)

var AllModelTypes = []Type{
	ModelTypeEnum, ModelTypeSequence, ModelTypeStruct,
	ModelTypeInterface, ModelTypeUnion, ModelTypeIntersection,
}

func modelTypeFromString(s string) Type {
	for _, t := range AllModelTypes {
		if t.Key == s {
			return t
		}
	}
	util.LogWarn("unhandled model type [" + s + "]")
	return ModelTypeUnknown
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
	*t = modelTypeFromString(s)
	return nil
}
