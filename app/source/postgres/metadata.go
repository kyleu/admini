package postgres

import (
	"github.com/kyleu/admini/app/database"
	"github.com/kyleu/admini/app/schema"
)

func loadMetadata(key string, db *database.Service) (*schema.Metadata, error) {
	return &schema.Metadata{
		Description: "asdfsdf",
		Comments:    nil,
		Origin:      nil,
		Source:      "postgres:" + key,
		Line:        0,
		Column:      0,
	}, nil
}
