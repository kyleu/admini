package action

import (
	"encoding/json"

	"github.com/pkg/errors"
)

type Type struct {
	Key         string
	Title       string
	Icon        string
	Description string
}

var (
	TypeFolder    = Type{Key: "", Title: "Folder", Icon: "folder", Description: "holds other actions, like a folder"}
	TypeStatic    = Type{Key: "static", Title: "Static", Icon: "image", Description: "returns HTML for rendering"}
	TypeSeparator = Type{Key: "separator", Title: "Separator", Icon: "handle", Description: "a separator, used between other items"}

	TypeAll      = Type{Key: "all", Title: "All Sources", Icon: "world", Description: "provides actions for each source in the system"}
	TypeSource   = Type{Key: "source", Title: "Source", Icon: "location", Description: "provides actions for each model in the source"}
	TypePackage  = Type{Key: "package", Title: "Package", Icon: "tag", Description: "provides actions for a package contained in a source"}
	TypeModel    = Type{Key: "model", Title: "Model", Icon: "file-text", Description: "provides actions for a model contained in a source"}
	TypeActivity = Type{Key: "activity", Title: "Activity", Icon: "happy", Description: "provides actions for a specific activity"}

	TypeTest    = Type{Key: "test", Title: "Test", Icon: "star", Description: "a test action, who knows what it'll do"}
	TypeUnknown = Type{Key: "unknown", Title: "Unknown", Icon: "star", Description: "an unknown action type"}
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
	return TypeUnknown, errors.Errorf("unhandled action type from string [%s]", s)
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
