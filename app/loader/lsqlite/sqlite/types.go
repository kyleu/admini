package sqlite

import (
	"strconv"
	"strings"

	"admini.dev/admini/app/lib/types"
	"admini.dev/admini/app/util"
)

func typeFor(t string, cr *columnResult, logger util.Logger) *types.Wrapped {
	if cr != nil && cr.NotNull == 0 {
		cr.NotNull = 1
		return types.NewOption(typeFor(t, cr, logger))
	}
	lt := strings.ToLower(t)
	switch {
	case lt == "int" || lt == "integer":
		return types.NewInt(0)
	case lt == "datetime":
		return types.NewTimestamp()
	case strings.HasPrefix(lt, "nvarchar"):
		s := lt[strings.LastIndex(lt, "(")+1 : strings.LastIndex(lt, ")")]
		max, _ := strconv.Atoi(s)
		return types.NewStringArgs(0, max, "")
	}
	logger.Warn("unhandled SQLite type: " + t)
	return types.NewUnknown(t)
}
