package mysql

import (
	"context"

	"github.com/kyleu/admini/app/model"
	"go.uber.org/zap"

	"github.com/pkg/errors"

	"github.com/kyleu/admini/app/database"
	"github.com/kyleu/admini/app/schema"
)

func LoadDatabaseSchema(ctx context.Context, db *database.Service, logger *zap.SugaredLogger) (*schema.Schema, error) {
	metadata := loadMetadata(ctx, db)

	tables, err := loadTables(ctx, nil, db, logger)
	if err != nil {
		return nil, errors.Wrap(err, "can't load tables")
	}

	err = loadColumns(ctx, tables, db, logger)
	if err != nil {
		return nil, errors.Wrap(err, "can't load columns")
	}

	models := make(model.Models, 0, len(tables))
	models = append(models, tables...)
	models.Sort()

	ret := &schema.Schema{
		Paths:    []string{"mysql:" + db.DatabaseName},
		Models:   models,
		Metadata: metadata,
	}

	return ret, nil
}
