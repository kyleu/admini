package sqlite

import (
	"github.com/kyleu/admini/app/database"
	"github.com/kyleu/admini/app/schema"
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