package sqlite

import (
	"admini.dev/app/lib/database"
	"admini.dev/app/lib/schema"
)

func loadMetadata(db *database.Service) *schema.Metadata {
	return &schema.Metadata{
		Description: "",
		Comments:    nil,
		Origin:      nil,
		Source:      "sqlite:" + db.DatabaseName,
		Line:        0,
		Column:      0,
	}
}
