package mysql

import (
	"strings"

	types2 "github.com/kyleu/admini/app/schema/types"
	"go.uber.org/zap"
)

func TypeForName(t string, logger *zap.SugaredLogger) *types2.Wrapped {
	if strings.HasPrefix(t, "_") {
		return types2.NewList(TypeForName(t[1:], logger))
	}
	return typeFor(t, nil, logger)
}

// nolint
func typeFor(t string, cr *columnResult, logger *zap.SugaredLogger) *types2.Wrapped {
	if cr != nil && cr.Nullable == "YES" {
		cr.Nullable = "NO"
		return types2.NewOption(typeFor(t, cr, logger))
	}
	if strings.HasPrefix(t, "_") {
		return types2.NewList(typeFor(t[1:], cr, logger))
	}
	// TODO: limit to MySQL types
	switch strings.ToLower(t) {
	case "bit":
		if cr != nil && cr.CharLength.Valid {
			return types2.NewListSized(types2.NewBit(), int(cr.CharLength.Int64))
		}
		return types2.NewBit()
	case "varbit", "bit varying":
		return types2.NewList(types2.NewBit())
	case "bool", "boolean":
		return types2.NewBool()
	case "bpchar":
		return stringFor(cr)
	case "bytea":
		return types2.NewList(types2.NewByte())
	case "char", "character":
		return types2.NewChar()
	case "character varying", "varchar":
		return stringFor(cr)
	case "date":
		return types2.NewDate()
	case "daterange":
		return types2.NewRange(types2.NewDate())
	case "float4", "real", "float":
		return types2.NewFloat(32)
	case "float8", "double precision", "double":
		return types2.NewFloat(64)
	case "hstore":
		return types2.NewMap(types2.NewString(), types2.NewString())
	case "int2", "smallint":
		return types2.NewInt(16)
	case "int2range":
		return types2.NewRange(types2.NewInt(16))
	case "int4", "integer", "int", "mediumint":
		return types2.NewInt(32)
	case "int4range":
		return types2.NewRange(types2.NewInt(32))
	case "int8", "bigint":
		return types2.NewInt(64)
	case "int8range":
		return types2.NewRange(types2.NewInt(64))
	case "json":
		return types2.NewJSON()
	case "jsonb":
		return types2.NewJSON()
	case "name":
		return stringFor(cr)
	case "numeric", "decimal":
		// return schematypes.NewNumeric()
	case "numrange":
		return types2.NewRange(types2.NewFloat(64))
	case "oid":
		return types2.NewInt(32)
	case "text":
		return stringFor(cr)
	case "time", "time without time zone":
		return types2.NewTime()
	case "timetz", "time with time zone":
		// return schematypes.NewTimeTZ()
	case "timestamp", "timestamp without time zone", "datetime":
		return types2.NewTimestamp()
	case "timestamptz", "timestamp with time zone":
		return types2.NewTimestampZoned()
	case "uuid":
		return types2.NewUUID()
	case "USER-DEFINED":
		return types2.NewReference()
	}
	logger.Warn("unhandled MySQL type: [" + t + "]")
	return types2.NewUnknown(t)
}

func stringFor(cr *columnResult) *types2.Wrapped {
	max := 0
	if cr != nil && cr.CharLength.Valid {
		max = int(cr.CharLength.Int64)
	}
	return types2.NewStringArgs(0, max, "")
}
