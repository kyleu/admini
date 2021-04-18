package schema

import "github.com/kyleu/admini/app/schema/schematypes"

func GetExample() Fields {
	return Fields{
		{Key: "id", Type: schematypes.Wrap(&schematypes.UUID{}), Metadata: &Metadata{Description: "id of the model"}},
		{Key: "name", Type: schematypes.Wrap(&schematypes.String{MaxLength: 1024}), Metadata: &Metadata{Description: "name of the model"}},
	}
}
