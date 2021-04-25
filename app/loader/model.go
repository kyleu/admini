package loader

import (
	"encoding/json"

	"github.com/kyleu/admini/app/schema"
)

type Loader interface {
	GetSchema(config json.RawMessage) (*schema.Schema, error)
}
