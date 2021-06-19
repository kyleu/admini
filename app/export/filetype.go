package export

import (
	"github.com/kyleu/admini/app/util"
)

type FileType struct {
	Key         string `json:"key"`
	Title       string `json:"title,omitempty"`
	Icon        string `json:"icon,omitempty"`
	Description string `json:"description,omitempty"`
}

var (
	FileTypeGo      = FileType{Key: "go", Title: "Go", Icon: "file", Description: "Golang source code"}
	FileTypeJSON    = FileType{Key: "json", Title: "JSON", Icon: "file", Description: "JavaScript object notation"}
	FileTypeText    = FileType{Key: "text", Title: "Text", Icon: "file", Description: "Plain text"}
	FileTypeUnknown = FileType{Key: "unknown", Title: "unknown", Icon: "file", Description: "Unknown file type"}
)

var AllFileTypes = []FileType{FileTypeGo, FileTypeText}

func FileTypeFromString(s string) FileType {
	for _, t := range AllFileTypes {
		if t.Key == s {
			return t
		}
	}
	return FileTypeUnknown
}

func (t *FileType) String() string {
	return t.Key
}

func (t *FileType) MarshalJSON() ([]byte, error) {
	return util.ToJSONBytes(t.Key, false), nil
}

func (t *FileType) UnmarshalJSON(data []byte) error {
	var s string
	err := util.FromJSON(data, &s)
	if err != nil {
		return err
	}
	*t = FileTypeFromString(s)
	return nil
}
