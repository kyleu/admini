package action

import (
	"encoding/json"

	"github.com/kyleu/admini/app/util"
)

type Type struct {
	Key   string
	Title string
}

var (
	ActionTypeFoo = Type{Key: "foo", Title: "Foo"}
	ActionTypeBar = Type{Key: "bar", Title: "Bar"}
	ActionTypeUnknown = Type{Key: "unknown", Title: "Unknown"}
)

var AllActionTypes = []Type{ActionTypeFoo, ActionTypeBar}

func actionTypeFromString(s string) Type {
	for _, t := range AllActionTypes {
		if t.Key == s {
			return t
		}
	}
	util.LogWarn("unhandled action type [" + s + "]")
	return ActionTypeUnknown
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
	*t = actionTypeFromString(s)
	return nil
}
