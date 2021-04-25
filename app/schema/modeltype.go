package schema

import (
	"encoding/json"

	"github.com/kyleu/admini/app/util"
)

type ModelType struct {
	Key    string
	Title  string
	Plural string
}

var (
	ModelTypeEnum         = ModelType{Key: "enum", Title: "Enum", Plural: "Enums"}
	ModelTypeSequence     = ModelType{Key: "sequence", Title: "Sequence", Plural: "Sequences"}
	ModelTypeStruct       = ModelType{Key: "struct", Title: "Struct", Plural: "Structs"}
	ModelTypeInterface    = ModelType{Key: "interface", Title: "Interface", Plural: "Interfaces"}
	ModelTypeUnion        = ModelType{Key: "union", Title: "Union", Plural: "Unions"}
	ModelTypeIntersection = ModelType{Key: "intersection", Title: "Intersection", Plural: "Intersections"}
	ModelTypeUnknown      = ModelType{Key: "unknown", Title: "Unknown", Plural: "Unknowns"}
)

var AllModelTypes = []ModelType{
	ModelTypeEnum, ModelTypeSequence, ModelTypeStruct,
	ModelTypeInterface, ModelTypeUnion, ModelTypeIntersection,
}

func modelTypeFromString(s string) ModelType {
	for _, t := range AllModelTypes {
		if t.Key == s {
			return t
		}
	}
	util.LogWarn("unhandled model type [" + s + "]")
	return ModelTypeUnknown
}

func (t *ModelType) String() string {
	return t.Key
}

func (t *ModelType) MarshalJSON() ([]byte, error) {
	return json.Marshal(t.Key)
}

func (t *ModelType) UnmarshalJSON(data []byte) error {
	var s string
	err := json.Unmarshal(data, &s)
	if err != nil {
		return err
	}
	*t = modelTypeFromString(s)
	return nil
}
