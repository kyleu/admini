package sqlserver

import (
	"context"

	"admini.dev/admini/app/lib/database"
	"admini.dev/admini/app/lib/schema"
)

func loadMetadata(_ context.Context, db *database.Service) *schema.Metadata {
	return &schema.Metadata{
		Description: "",
		Comments:    nil,
		Origin:      nil,
		Source:      "sqlserver:" + db.DatabaseName,
		Line:        0,
		Column:      0,
	}
}
