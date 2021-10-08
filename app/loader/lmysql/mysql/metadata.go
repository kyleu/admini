package mysql

import (
	"context"

	"github.com/kyleu/admini/app/database"
	"github.com/kyleu/admini/app/schema"
)

func loadMetadata(ctx context.Context, db *database.Service) *schema.Metadata {
	return &schema.Metadata{
		Description: "",
		Comments:    nil,
		Origin:      nil,
		Source:      "mysql:" + db.DatabaseName,
		Line:        0,
		Column:      0,
	}
}