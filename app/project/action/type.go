package action

import (
	"encoding/json"

	"github.com/kyleu/admini/app/util"
)

type Type struct {
	Key         string
	Title       string
	Description string
}

var (
	ActionTypeSource  = Type{Key: "source", Title: "Source", Description: "provides actions for each model in the source"}
	ActionTypeStatic  = Type{Key: "static", Title: "Static", Description: "returns HTML for rendering"}
	ActionTypeTest    = Type{Key: "test", Title: "Test", Description: "a test action, who knows what it'll do"}
	ActionTypeUnknown = Type{Key: "unknown", Title: "Unknown", Description: "an unknown action type"}
)

var AllActionTypes = []Type{
	ActionTypeSource, ActionTypeStatic, ActionTypeTest,
}

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
