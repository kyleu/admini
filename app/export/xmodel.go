package export

import (
	"github.com/pkg/errors"

	"github.com/kyleu/admini/app/model"
	"go.uber.org/zap"
)

func Model(m *model.Model, t *Format, logger *zap.SugaredLogger) (Results, error) {
	switch t.Language {
	case "go":
		return Results{goModelFile(m, t, logger), goServiceFile(m, t)}, nil
	case "json":
		return Results{jsonFile(m, t, logger)}, nil
	default:
		return nil, errors.Errorf("unhandled language [%s]", t.Language)
	}
}
