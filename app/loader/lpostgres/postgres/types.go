package postgres

import (
	"strings"

	"admini.dev/admini/app/lib/schema/model"
	"admini.dev/admini/app/lib/types"
	"admini.dev/admini/app/util"
)

func TypeForName(t string, enums model.Models, logger util.Logger) *types.Wrapped {
	if strings.HasPrefix(t, "_") {
		return types.NewList(TypeForName(t[1:], enums, logger))
	}
	return typeFor(t, nil, enums, logger)
}

//nolint:cyclop, funlen, gocyclo
func typeFor(t string, cr *columnResult, enums model.Models, logger util.Logger) *types.Wrapped {
	if cr != nil && cr.Nullable == pgYes {
		cr.Nullable = pgNo
		return types.NewOption(typeFor(t, cr, enums, logger))
	}
	if strings.HasPrefix(t, "_") {
		return types.NewList(typeFor(t[1:], cr, enums, logger))
	}
	if t == "ARRAY" && cr != nil && cr.ArrayType.Valid {
		return types.NewList(typeFor(cr.ArrayType.String, cr, enums, logger))
	}
	switch strings.ToLower(t) {
	case "aclitem":
		// return types.NewACL()
	case "bit":
		if cr != nil && cr.CharLength.Valid {
			return types.NewListSized(types.NewBit(), int(cr.CharLength.Int32))
		}
		return types.NewBit()
	case "varbit", "bit varying":
		return types.NewList(types.NewBit())
	case "bool", "boolean":
		return types.NewBool()
	case "box":
		// return types.NewBox()
	case "bpchar":
		return stringFor(cr)
	case "bytea":
		return types.NewList(types.NewByte())
	case "char", "character":
		return types.NewChar()
	case "character varying", "varchar":
		return stringFor(cr)
	case "cid":
		// return types.NewCID()
	case "cidr":
		// return types.NewCIDR()
	case "circle":
		// return types.NewCircle()
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
	case "inet":
		// return types.NewInet()
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
	case "interval":
		// return types.NewInterval()
	case "json":
		return types.NewJSON()
	case "jsonb":
		return types.NewJSON()
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
		return types.NewRange(types.NewFloat(64))
	case "oid":
		return types.NewInt(32)
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
		return types.NewTime()
	case "timetz", "time with time zone":
		// return types.NewTimeTZ()
	case "timestamp", "timestamp without time zone", "datetime":
		return types.NewTimestamp()
	case "timestamptz", "timestamp with time zone":
		return types.NewTimestampZoned()
	case "tsrange":
		return types.NewRange(types.NewTimestamp())
	case "tsquery":
		// return types.NewTsQuery()
	case "tsvector":
		return types.NewList(types.NewTimestamp())
	case "tstzrange":
		return types.NewRange(types.NewTimestampZoned())
	case "uuid":
		return types.NewUUID()
	case "USER-DEFINED":
		return types.NewReference()
	case "xid":
		// return types.NewXID()
	case "xml":
		return types.NewXML()
	case "year":
		// return types.NewYear()
	}
	if e := enums.Get(util.Pkg{cr.Schema}, t); e != nil {
		return types.NewEnum(cr.UDTName)
	}
	logger.Warn("unhandled PostgreSQL type: " + t)
	return types.NewUnknown(t)
}

func stringFor(cr *columnResult) *types.Wrapped {
	max := 0
	if cr != nil && cr.CharLength.Valid {
		max = int(cr.CharLength.Int32)
	}
	return types.NewStringArgs(0, max, "")
}
