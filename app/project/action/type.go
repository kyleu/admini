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
	TypeFolder    = Type{Key: "", Title: "Folder", Description: "holds other actions, like a folder"}
	TypeStatic    = Type{Key: "static", Title: "Static", Description: "returns HTML for rendering"}
	TypeSeparator = Type{Key: "separator", Title: "Separator", Description: "a separator, used between other items"}

	TypeAll      = Type{Key: "all", Title: "All Sources", Description: "provides actions for each source in the system"}
	TypeSource   = Type{Key: "source", Title: "Source", Description: "provides actions for each model in the source"}
	TypePackage  = Type{Key: "package", Title: "Package", Description: "provides actions for a package contained in a source"}
	TypeModel    = Type{Key: "model", Title: "Model", Description: "provides actions for a model contained in a source"}
	TypeActivity = Type{Key: "activity", Title: "Activity", Description: "provides actions for a specific activity"}

	TypeTest    = Type{Key: "test", Title: "Test", Description: "a test action, who knows what it'll do"}
	TypeUnknown = Type{Key: "unknown", Title: "Unknown", Description: "an unknown action type"}
)

var AllTypes = []Type{
	TypeFolder, TypeStatic, TypeSeparator, TypeAll, TypeSource, TypePackage, TypeModel, TypeActivity, TypeTest,
}

func TypeFromString(s string) (Type, error) {
	for _, t := range AllTypes {
		if t.Key == s {
			return t, nil
		}
	}
	if s == "folder" {
		return TypeFolder, nil
	}
	return TypeUnknown, errors.New("unhandled action type from string [" + s + "]")
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
	x, err := TypeFromString(s)
	if err != nil {
		return err
	}
	*t = x
	return nil
}
