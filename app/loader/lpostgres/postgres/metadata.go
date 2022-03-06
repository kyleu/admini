package postgres

import (
	"context"

	"admini.dev/app/lib/database"
	"admini.dev/app/lib/schema"
)

func loadMetadata(ctx context.Context, db *database.Service) *schema.Metadata {
	return &schema.Metadata{
		Description: "",
		Comments:    nil,
		Origin:      nil,
		Source:      "postgres:" + db.DatabaseName,
		Line:        0,
		Column:      0,
	}
}
