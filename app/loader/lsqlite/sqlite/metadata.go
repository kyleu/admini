package sqlite

import (
	"admini.dev/admini/app/lib/database"
	"admini.dev/admini/app/lib/schema"
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
