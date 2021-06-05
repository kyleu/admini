package database

import (
	"fmt"
	"strings"
)

const whereSpaces = " where "

func SQLInsert(table string, columns []string, rows int) string {
	if rows <= 0 {
		rows = 1
	}
	colString := strings.Join(columns, ", ")
	var placeholders []string
	for i := 0; i < rows; i++ {
		var ph []string
		for idx := range columns {
			ph = append(ph, fmt.Sprintf("$%d", (i*len(columns))+idx+1))
		}
		placeholders = append(placeholders, "("+strings.Join(ph, ", ")+")")
	}
	return fmt.Sprintf("insert into %s (%s) values %s", table, colString, strings.Join(placeholders, ", "))
}

func SQLSelect(columns string, tables string, where string, orderBy string, limit int, offset int) string {
	if columns == "" {
		columns = "*"
	}

	whereClause := ""
	if len(where) > 0 {
		whereClause = whereSpaces + where
	}

	orderByClause := ""
	if len(orderBy) > 0 {
		orderByClause = " order by " + orderBy
	}

	limitClause := ""
	if limit > 0 {
		limitClause = fmt.Sprintf(" limit %d", limit)
	}

	offsetClause := ""
	if offset > 0 {
		offsetClause = fmt.Sprintf(" offset %d", offset)
	}

	return "select " + columns + " from " + tables + whereClause + orderByClause + limitClause + offsetClause
}

func SQLSelectSimple(columns string, tables string, where ...string) string {
	return SQLSelect(columns, tables, strings.Join(where, " and "), "", 0, 0)
}

func SQLUpdate(table string, columns []string, where string) string {
	whereClause := ""
	if len(where) > 0 {
		whereClause = whereSpaces + where
	}

	stmts := make([]string, 0, len(columns))
	for i, col := range columns {
		stmts = append(stmts, fmt.Sprintf("%s = $%d", col, i+1))
	}
	return fmt.Sprintf("update %s set %s%s", table, strings.Join(stmts, ", "), whereClause)
}

func SQLDelete(table string, where string) string {
	if strings.TrimSpace(where) == "" {
		return fmt.Sprintf("attempt to delete from [%s] with empty where clause", table)
	}
	return "delete from " + table + whereSpaces + where
}
