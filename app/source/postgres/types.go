package postgres

import (
	"strings"

	"github.com/kyleu/admini/app/schema/schematypes"
	"github.com/kyleu/admini/app/util"
)

func TypeForName(t string) *schematypes.Wrapped {
	if strings.HasPrefix(t, "_") {
		return schematypes.NewList(TypeForName(t[1:]))
	}
	return simpleType(strings.ToLower(t))
}

func typeFor(cr *columnResult) *schematypes.Wrapped {
	ret := simpleType(cr.DataType)
	if err := enhance(cr, ret); err != nil {
		return schematypes.NewUnknown("ERROR: " + err.Error())
	}

	return ret
}

func enhance(cr *columnResult, ret *schematypes.Wrapped) error {
	switch t := ret.T.(type) {
	case *schematypes.String:
		t.MaxLength = int(cr.CharLength.Int32)
	case *schematypes.List:
		if t.T == nil && cr.ArrayType.Valid {
			t.T = simpleType(cr.ArrayType.String)
		}
	}
	return nil
}

func simpleType(t string) *schematypes.Wrapped {
	switch t {
	case "aclitem":
		// return schematypes.NewACL()
	case "array", "ARRAY":
		return schematypes.NewList(nil)
	case "bit":
		return schematypes.NewBit()
	case "varbit", "bit varying":
		return schematypes.NewList(schematypes.NewBit())
	case "bool", "boolean":
		return schematypes.NewBool()
	case "box":
		// return schematypes.NewBox()
	case "bpchar":
		// return schematypes.NewBpchar()
	case "bytea":
		return schematypes.NewList(schematypes.NewByte())
	case "char", "character":
		return schematypes.NewChar()
	case "character varying", "varchar":
		return schematypes.NewString()
	case "cid":
		// return schematypes.NewCID()
	case "cidr":
		// return schematypes.NewCIDR()
	case "circle":
		// return schematypes.NewCircle()
	case "date":
		return schematypes.NewDate()
	case "daterange":
		return schematypes.NewList(schematypes.NewDate())
	case "float4", "real", "float":
		return schematypes.NewFloat()
	case "float8", "double precision", "double":
		return schematypes.NewFloat()
	case "hstore":
		return schematypes.NewMap(schematypes.NewString(), schematypes.NewString())
	case "inet":
		// return schematypes.NewInet()
	case "int1", "tinyint":
		return schematypes.NewInt()
	case "int2", "smallint":
		return schematypes.NewInt()
	case "int4", "integer", "int", "mediumint":
		return schematypes.NewInt()
	case "int4range":
		return schematypes.NewList(schematypes.NewInt())
	case "int8", "bigint":
		return schematypes.NewInt()
	case "int8range":
		return schematypes.NewList(schematypes.NewInt())
	case "interval":
		// return schematypes.NewInterval()
	case "json":
		return schematypes.NewJSON()
	case "jsonb":
		return schematypes.NewJSON()
	case "line":
		// return schematypes.NewLine()
	case "lseg":
		// return schematypes.NewLineSegment()
	case "macaddr":
		// return schematypes.NewMacAddr()
	case "money":
		// return schematypes.NewMoney()
	case "name":
		return schematypes.NewString()
	case "numeric", "decimal":
		// return schematypes.NewNumeric()
	case "numrange":
		return schematypes.NewList(schematypes.NewFloat())
	case "oid":
		// return schematypes.NewOID()
	case "path":
		// return schematypes.NewPath()
	case "point":
		// return schematypes.NewPoint()
	case "polygon":
		// return schematypes.NewPolygon()
	case "record":
		// return schematypes.NewRecord()
	case "text":
		return schematypes.NewString()
	case "tid":
		// return schematypes.NewTID()
	case "time", "time without time zone":
		return schematypes.NewTime()
	case "timetz", "time with time zone":
		// return schematypes.NewTimeTZ()
	case "timestamp", "timestamp without time zone", "datetime":
		return schematypes.NewTimestamp()
	case "timestamptz", "timestamp with time zone":
		return schematypes.NewTimestampZoned()
	case "tsrange":
		return schematypes.NewRange(schematypes.NewTimestamp())
	case "tsquery":
		// return schematypes.NewTsQuery()
	case "tsvector":
		return schematypes.NewList(schematypes.NewTimestamp())
	case "tstzrange":
		return schematypes.NewRange(schematypes.NewTimestampZoned())
	case "uuid":
		return schematypes.NewUUID()
	case "USER-DEFINED":
		return schematypes.NewReference()
	case "xid":
		// return schematypes.NewXID()
	case "xml":
		return schematypes.NewXML()
	case "year":
		// return schematypes.NewYear()
	}
	util.LogWarn("unhandled type: " + t)
	return schematypes.NewUnknown(t)
}
