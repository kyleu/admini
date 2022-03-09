package postgres

import (
	"strings"

	types2 "admini.dev/admini/app/lib/types"
	"go.uber.org/zap"

	"admini.dev/admini/app/lib/schema/model"
	"admini.dev/admini/app/util"
)

func TypeForName(t string, enums model.Models, logger *zap.SugaredLogger) *types2.Wrapped {
	if strings.HasPrefix(t, "_") {
		return types2.NewList(TypeForName(t[1:], enums, logger))
	}
	return typeFor(t, nil, enums, logger)
}

// nolint
func typeFor(t string, cr *columnResult, enums model.Models, logger *zap.SugaredLogger) *types2.Wrapped {
	if cr != nil && cr.Nullable == pgYes {
		cr.Nullable = pgNo
		return types2.NewOption(typeFor(t, cr, enums, logger))
	}
	if strings.HasPrefix(t, "_") {
		return types2.NewList(typeFor(t[1:], cr, enums, logger))
	}
	if t == "ARRAY" && cr != nil && cr.ArrayType.Valid {
		return types2.NewList(typeFor(cr.ArrayType.String, cr, enums, logger))
	}
	switch strings.ToLower(t) {
	case "aclitem":
		// return types.NewACL()
	case "bit":
		if cr != nil && cr.CharLength.Valid {
			return types2.NewListSized(types2.NewBit(), int(cr.CharLength.Int32))
		}
		return types2.NewBit()
	case "varbit", "bit varying":
		return types2.NewList(types2.NewBit())
	case "bool", "boolean":
		return types2.NewBool()
	case "box":
		// return types.NewBox()
	case "bpchar":
		return stringFor(cr)
	case "bytea":
		return types2.NewList(types2.NewByte())
	case "char", "character":
		return types2.NewChar()
	case "character varying", "varchar":
		return stringFor(cr)
	case "cid":
		// return types.NewCID()
	case "cidr":
		// return types.NewCIDR()
	case "circle":
		// return types.NewCircle()
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
	case "inet":
		// return types.NewInet()
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
	case "interval":
		// return types.NewInterval()
	case "json":
		return types2.NewJSON()
	case "jsonb":
		return types2.NewJSON()
	case "line":
		// return types.NewLine()
	case "lseg":
		// return types.NewLineSegment()
	case "macaddr":
		// return types.NewMacAddr()
	case "money":
		// return types.NewMoney()
	case "name":
		return stringFor(cr)
	case "numeric", "decimal":
		// return types.NewNumeric()
	case "numrange":
		return types2.NewRange(types2.NewFloat(64))
	case "oid":
		return types2.NewInt(32)
	case "path":
		// return types.NewPath()
	case "point":
		// return types.NewPoint()
	case "polygon":
		// return types.NewPolygon()
	case "record":
		// return types.NewRecord()
	case "text":
		return stringFor(cr)
	case "tid":
		// return types.NewTID()
	case "time", "time without time zone":
		return types2.NewTime()
	case "timetz", "time with time zone":
		// return types.NewTimeTZ()
	case "timestamp", "timestamp without time zone", "datetime":
		return types2.NewTimestamp()
	case "timestamptz", "timestamp with time zone":
		return types2.NewTimestampZoned()
	case "tsrange":
		return types2.NewRange(types2.NewTimestamp())
	case "tsquery":
		// return types.NewTsQuery()
	case "tsvector":
		return types2.NewList(types2.NewTimestamp())
	case "tstzrange":
		return types2.NewRange(types2.NewTimestampZoned())
	case "uuid":
		return types2.NewUUID()
	case "USER-DEFINED":
		return types2.NewReference()
	case "xid":
		// return types.NewXID()
	case "xml":
		return types2.NewXML()
	case "year":
		// return types.NewYear()
	}
	if e := enums.Get(util.Pkg{cr.Schema}, t); e != nil {
		return types2.NewEnum(cr.UDTName)
	}
	logger.Warn("unhandled PostgreSQL type: " + t)
	return types2.NewUnknown(t)
}

func stringFor(cr *columnResult) *types2.Wrapped {
	max := 0
	if cr != nil && cr.CharLength.Valid {
		max = int(cr.CharLength.Int32)
	}
	return types2.NewStringArgs(0, max, "")
}
