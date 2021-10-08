package mysql

import (
	"strings"

	"go.uber.org/zap"

	"github.com/kyleu/admini/app/schema/schematypes"
)

func TypeForName(t string, logger *zap.SugaredLogger) *schematypes.Wrapped {
	if strings.HasPrefix(t, "_") {
		return schematypes.NewList(TypeForName(t[1:], logger))
	}
	return typeFor(t, nil, logger)
}

// nolint
func typeFor(t string, cr *columnResult, logger *zap.SugaredLogger) *schematypes.Wrapped {
	if cr != nil && cr.Nullable == "YES" {
		cr.Nullable = "NO"
		return schematypes.NewOption(typeFor(t, cr, logger))
	}
	if strings.HasPrefix(t, "_") {
		return schematypes.NewList(typeFor(t[1:], cr, logger))
	}
	// TODO: limit to MySQL types
	switch strings.ToLower(t) {
	case "bit":
		if cr != nil && cr.CharLength.Valid {
			return schematypes.NewListSized(schematypes.NewBit(), int(cr.CharLength.Int64))
		}
		return schematypes.NewBit()
	case "varbit", "bit varying":
		return schematypes.NewList(schematypes.NewBit())
	case "bool", "boolean":
		return schematypes.NewBool()
	case "bpchar":
		return stringFor(cr)
	case "bytea":
		return schematypes.NewList(schematypes.NewByte())
	case "char", "character":
		return schematypes.NewChar()
	case "character varying", "varchar":
		return stringFor(cr)
	case "date":
		return schematypes.NewDate()
	case "daterange":
		return schematypes.NewRange(schematypes.NewDate())
	case "float4", "real", "float":
		return schematypes.NewFloat(32)
	case "float8", "double precision", "double":
		return schematypes.NewFloat(64)
	case "hstore":
		return schematypes.NewMap(schematypes.NewString(), schematypes.NewString())
	case "int2", "smallint":
		return schematypes.NewInt(16)
	case "int2range":
		return schematypes.NewRange(schematypes.NewInt(16))
	case "int4", "integer", "int", "mediumint":
		return schematypes.NewInt(32)
	case "int4range":
		return schematypes.NewRange(schematypes.NewInt(32))
	case "int8", "bigint":
		return schematypes.NewInt(64)
	case "int8range":
		return schematypes.NewRange(schematypes.NewInt(64))
	case "json":
		return schematypes.NewJSON()
	case "jsonb":
		return schematypes.NewJSON()
	case "name":
		return stringFor(cr)
	case "numeric", "decimal":
		// return schematypes.NewNumeric()
	case "numrange":
		return schematypes.NewRange(schematypes.NewFloat(64))
	case "oid":
		return schematypes.NewInt(32)
	case "text":
		return stringFor(cr)
	case "time", "time without time zone":
		return schematypes.NewTime()
	case "timetz", "time with time zone":
		// return schematypes.NewTimeTZ()
	case "timestamp", "timestamp without time zone", "datetime":
		return schematypes.NewTimestamp()
	case "timestamptz", "timestamp with time zone":
		return schematypes.NewTimestampZoned()
	case "uuid":
		return schematypes.NewUUID()
	case "USER-DEFINED":
		return schematypes.NewReference()
	}
	logger.Warn("unhandled MySQL type: [" + t + "]")
	return schematypes.NewUnknown(t)
}

func stringFor(cr *columnResult) *schematypes.Wrapped {
	max := 0
	if cr != nil && cr.CharLength.Valid {
		max = int(cr.CharLength.Int64)
	}
	return schematypes.NewStringArgs(0, max, "")
}
