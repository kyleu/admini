package sqlite

import (
	"strconv"
	"strings"

	"go.uber.org/zap"

	"github.com/kyleu/admini/app/schema/schematypes"
)

// nolint
func typeFor(t string, cr *columnResult, logger *zap.SugaredLogger) *schematypes.Wrapped {
	if cr != nil && cr.NotNull == 0 {
		cr.NotNull = 1
		return schematypes.NewOption(typeFor(t, cr, logger))
	}
	lt := strings.ToLower(t)
	switch {
	case lt == "int" || lt == "integer":
		return schematypes.NewInt(0)
	case lt == "datetime":
		return schematypes.NewTimestamp()
	case strings.HasPrefix(lt, "nvarchar"):
		s := lt[strings.LastIndex(lt, "(")+1 : strings.LastIndex(lt, ")")]
		max, _ := strconv.Atoi(s)
		return schematypes.NewStringArgs(0, max, "")
	}
	logger.Warn("unhandled SQLite type: " + t)
	return schematypes.NewUnknown(t)
}
