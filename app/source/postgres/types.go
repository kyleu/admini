package postgres

import (
	"github.com/kyleu/admini/app/schema/schematypes"
	"github.com/kyleu/admini/app/util"
)

func typeFor(cr *columnResult) *schematypes.Wrapped {
	ret := simpleType(cr.DataType)
	if err := enhance(cr, ret); err != nil {
		return schematypes.Wrap(&schematypes.Unknown{X: "ERROR: " + err.Error()})
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
	var ret schematypes.Type
	switch t {
	case "aclitem":
		// TODO
		// ret = &schematypes.ACL
	case "array", "ARRAY":
		// TODO
		ret = &schematypes.List{}
	case "bit":
		// TODO
		ret = &schematypes.Bit{}
	case "varbit", "bit varying":
		// TODO
		ret = &schematypes.List{T: schematypes.Wrap(&schematypes.Bit{})}
	case "bool", "boolean":
		ret = &schematypes.Bool{}
	case "box":
		// TODO
		// ret = &schematypes.Box
	case "bpchar":
		// TODO
		// ret = &schematypes.Bpchar
	case "bytea":
		// TODO
		ret = &schematypes.List{T: schematypes.Wrap(&schematypes.Byte{})}
	case "char", "character":
		// TODO
		ret = &schematypes.Char{}
	case "character varying", "varchar":
		// TODO
		ret = &schematypes.String{}
	case "cid":
		// TODO
		// ret = &schematypes.CID
	case "cidr":
		// TODO
		// ret = &schematypes.CIDR
	case "circle":
		// TODO
		// ret = &schematypes.Circle
	case "date":
		// TODO
		ret = &schematypes.Date{}
	case "daterange":
		// TODO
		ret = &schematypes.List{T: schematypes.Wrap(&schematypes.Date{})}
	case "float4", "real", "float":
		// TODO
		ret = &schematypes.Float{}
	case "float8", "double precision", "double":
		// TODO
		ret = &schematypes.Float{}
	case "hstore":
		// TODO
		ret = &schematypes.Map{K: schematypes.Wrap(&schematypes.String{}), V: schematypes.Wrap(&schematypes.String{})}
	case "inet":
		// TODO
		// ret = &schematypes.Inet
	case "int1", "tinyint":
		// TODO
		ret = &schematypes.Int{}
	case "int2", "smallint":
		// TODO
		ret = &schematypes.Int{}
	case "int4", "integer", "int", "mediumint":
		// TODO
		ret = &schematypes.Int{}
	case "int4range":
		// TODO
		ret = &schematypes.List{T: schematypes.Wrap(&schematypes.Int{})}
	case "int8", "bigint":
		// TODO
		ret = &schematypes.Int{}
	case "int8range":
		// TODO
		ret = &schematypes.List{T: schematypes.Wrap(&schematypes.Int{})}
	case "interval":
		// TODO
		// ret = &schematypes.Interval
	case "json":
		ret = &schematypes.JSON{}
	case "jsonb":
		ret = &schematypes.JSON{}
	case "line":
		// TODO
		// ret = &schematypes.Line
	case "lseg":
		// TODO
		// ret = &schematypes.LineSegment
	case "macaddr":
		// TODO
		// ret = &schematypes.MacAddr
	case "money":
		// TODO
		// ret = &schematypes.Money
	case "name":
		// TODO
		// ret = &schematypes.Name
	case "numeric", "decimal":
		// TODO
		// ret = &schematypes.Numeric
	case "numrange":
		// TODO
		ret = &schematypes.List{T: schematypes.Wrap(&schematypes.Float{})}
	case "oid":
		// TODO
		// ret = &schematypes.OID
	case "path":
		// TODO
		// ret = &schematypes.Path
	case "point":
		// TODO
		// ret = &schematypes.Point
	case "polygon":
		// TODO
		// ret = &schematypes.Polygon
	case "record":
		// TODO
		// ret = &schematypes.Record
	case "text":
		ret = &schematypes.String{}
	case "tid":
		// TODO
		// ret = &schematypes.TID
	case "time", "time without time zone":
		// TODO
		ret = &schematypes.Time{}
	case "timetz", "time with time zone":
		// TODO
		// ret = &schematypes.TimeTZ
	case "timestamp", "timestamp without time zone", "datetime":
		// TODO
		ret = &schematypes.Timestamp{}
	case "timestamptz", "timestamp with time zone":
		// TODO
		ret = &schematypes.TimestampZoned{}
	case "tsrange":
		// TODO
		ret = &schematypes.Range{T: schematypes.Wrap(&schematypes.Timestamp{})}
	case "tsquery":
		// TODO
		// ret = &schematypes.TsQuery
	case "tsvector":
		// TODO
		ret = &schematypes.List{T: schematypes.Wrap(&schematypes.Timestamp{})}
	case "tstzrange":
		// TODO
		ret = &schematypes.Range{T: schematypes.Wrap(&schematypes.TimestampZoned{})}
	case "uuid":
		ret = &schematypes.UUID{}
	case "USER-DEFINED":
		// TODO
		ret = &schematypes.Reference{}
	case "xid":
		// TODO
		// ret = &schematypes.XID
	case "xml":
		// TODO
		ret = &schematypes.XML{}
	case "year":
		// TODO
		// ret = &schematypes.Year
	default:
		util.LogWarn("unhandled type: " + t)
		ret = &schematypes.Unknown{X: "unhandled: " + t}
	}
	if ret == nil {
		ret = &schematypes.Unknown{X: t}
	}
	return schematypes.Wrap(ret)
}
