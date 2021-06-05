package postgres

import (
	"github.com/kyleu/admini/app/database"
	"github.com/kyleu/admini/app/schema"
)

func loadMetadata(db *database.Service) *schema.Metadata {
	return &schema.Metadata{
		Description: "asdfsdf",
		Comments:    nil,
		Origin:      nil,
		Source:      "postgres:" + db.DatabaseName,
		Line:        0,
		Column:      0,
	}
}
