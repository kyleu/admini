package sqlite

import (
	"strconv"
	"strings"

	types2 "admini.dev/app/lib/types"
	"go.uber.org/zap"
)

// nolint
func typeFor(t string, cr *columnResult, logger *zap.SugaredLogger) *types2.Wrapped {
	if cr != nil && cr.NotNull == 0 {
		cr.NotNull = 1
		return types2.NewOption(typeFor(t, cr, logger))
	}
	lt := strings.ToLower(t)
	switch {
	case lt == "int" || lt == "integer":
		return types2.NewInt(0)
	case lt == "datetime":
		return types2.NewTimestamp()
	case strings.HasPrefix(lt, "nvarchar"):
		s := lt[strings.LastIndex(lt, "(")+1 : strings.LastIndex(lt, ")")]
		max, _ := strconv.Atoi(s)
		return types2.NewStringArgs(0, max, "")
	}
	logger.Warn("unhandled SQLite type: " + t)
	return types2.NewUnknown(t)
}
