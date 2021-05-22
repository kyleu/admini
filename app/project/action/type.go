package action

import (
	"encoding/json"

	"github.com/pkg/errors"
)

type Type struct {
	Key         string
	Title       string
	Description string
}

var (
	ActionTypeAll     = Type{Key: "all", Title: "All Sources", Description: "provides actions for each source in the system"}
	ActionTypeSource  = Type{Key: "source", Title: "Source", Description: "provides actions for each model in the source"}
	ActionTypePackage = Type{Key: "package", Title: "Package", Description: "provides actions for a package contained in a source"}
	ActionTypeModel   = Type{Key: "model", Title: "Model", Description: "provides actions for a model contained in a source"}
	ActionTypeStatic  = Type{Key: "static", Title: "Static", Description: "returns HTML for rendering"}
	ActionTypeTest    = Type{Key: "test", Title: "Test", Description: "a test action, who knows what it'll do"}
	ActionTypeUnknown = Type{Key: "unknown", Title: "Unknown", Description: "an unknown action type"}
)

var AllActionTypes = []Type{
	ActionTypeAll, ActionTypeSource, ActionTypePackage, ActionTypeModel, ActionTypeStatic, ActionTypeTest,
}

func actionTypeFromString(s string) (Type, error) {
	for _, t := range AllActionTypes {
		if t.Key == s {
			return t, nil
		}
	}
	return ActionTypeUnknown, errors.New("unhandled action type [" + s + "]")
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
	x, err := actionTypeFromString(s)
	if err != nil {
		return err
	}
	*t = x
	return nil
}
