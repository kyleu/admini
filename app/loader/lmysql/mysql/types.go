package mysql

import (
	"strings"

	"github.com/kyleu/admini/app/lib/schema/types"
	"go.uber.org/zap"
)

func TypeForName(t string, logger *zap.SugaredLogger) *types.Wrapped {
	if strings.HasPrefix(t, "_") {
		return types.NewList(TypeForName(t[1:], logger))
	}
	return typeFor(t, nil, logger)
}

// nolint
func typeFor(t string, cr *columnResult, logger *zap.SugaredLogger) *types.Wrapped {
	if cr != nil && cr.Nullable == "YES" {
		cr.Nullable = "NO"
		return types.NewOption(typeFor(t, cr, logger))
	}
	if strings.HasPrefix(t, "_") {
		return types.NewList(typeFor(t[1:], cr, logger))
	}
	// TODO: limit to MySQL types
	switch strings.ToLower(t) {
	case "bit":
		if cr != nil && cr.CharLength.Valid {
			return types.NewListSized(types.NewBit(), int(cr.CharLength.Int64))
		}
		return types.NewBit()
	case "varbit", "bit varying":
		return types.NewList(types.NewBit())
	case "bool", "boolean":
		return types.NewBool()
	case "bpchar":
		return stringFor(cr)
	case "bytea":
		return types.NewList(types.NewByte())
	case "char", "character":
		return types.NewChar()
	case "character varying", "varchar":
		return stringFor(cr)
	case "date":
		return types.NewDate()
	case "daterange":
		return types.NewRange(types.NewDate())
	case "float4", "real", "float":
		return types.NewFloat(32)
	case "float8", "double precision", "double":
		return types.NewFloat(64)
	case "hstore":
		return types.NewMap(types.NewString(), types.NewString())
	case "int2", "smallint":
		return types.NewInt(16)
	case "int2range":
		return types.NewRange(types.NewInt(16))
	case "int4", "integer", "int", "mediumint":
		return types.NewInt(32)
	case "int4range":
		return types.NewRange(types.NewInt(32))
	case "int8", "bigint":
		return types.NewInt(64)
	case "int8range":
		return types.NewRange(types.NewInt(64))
	case "json":
		return types.NewJSON()
	case "jsonb":
		return types.NewJSON()
	case "name":
		return stringFor(cr)
	case "numeric", "decimal":
		// return schematypes.NewNumeric()
	case "numrange":
		return types.NewRange(types.NewFloat(64))
	case "oid":
		return types.NewInt(32)
	case "text":
		return stringFor(cr)
	case "time", "time without time zone":
		return types.NewTime()
	case "timetz", "time with time zone":
		// return schematypes.NewTimeTZ()
	case "timestamp", "timestamp without time zone", "datetime":
		return types.NewTimestamp()
	case "timestamptz", "timestamp with time zone":
		return types.NewTimestampZoned()
	case "uuid":
		return types.NewUUID()
	case "USER-DEFINED":
		return types.NewReference()
	}
	logger.Warn("unhandled MySQL type: [" + t + "]")
	return types.NewUnknown(t)
}

func stringFor(cr *columnResult) *types.Wrapped {
	max := 0
	if cr != nil && cr.CharLength.Valid {
		max = int(cr.CharLength.Int64)
	}
	return types.NewStringArgs(0, max, "")
}
