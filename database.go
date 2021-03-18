package main

import (
	"fmt"
	. "github.com/billcoding/mybatis-code-generator/config"
	. "github.com/billcoding/mybatis-code-generator/model"
	"strings"
)

func tables(database string, c *Configuration) []*Table {
	whereSql := ""
	if c.IncludeTables != nil && len(c.IncludeTables) > 0 {
		whereSql = fmt.Sprintf("AND t.`TABLE_NAME` IN('%s')", strings.Join(c.IncludeTables, "','"))
	} else if c.ExcludeTables != nil && len(c.ExcludeTables) > 0 {
		whereSql = fmt.Sprintf("AND t.`TABLE_NAME` NOT IN('%s')", strings.Join(c.ExcludeTables, "','"))
	}
	tableList := SelectTableListSelectMapper.Prepare(map[string]interface{}{
		"DBName": database,
		"Where":  whereSql,
	}).Exec().List(new(Table))
	ts := make([]*Table, len(tableList))
	for i, t := range tableList {
		tt := t.(*Table)
		tt.Columns = make([]*Column, 0)
		ts[i] = tt
	}
	return ts
}

func columns(database string) []*Column {
	columnList := SelectTableColumnListSelectMapper.Prepare(database).Exec().List(new(Column))
	cs := make([]*Column, len(columnList))
	for i, c := range columnList {
		cs[i] = c.(*Column)
		switch cs[i].Type {
		case "int":
			cs[i].Type = "integer"
			cs[i].UpperType = "INTEGER"
		case "datetime":
			cs[i].UpperType = "DATE"
		case "text":
			cs[i].Type = "varchar"
			cs[i].UpperType = "VARCHAR"
		}
	}
	return cs
}

func transformTables(tables []*Table) map[string]*Table {
	tableMap := make(map[string]*Table, len(tables))
	for _, t := range tables {
		tableMap[t.Name] = t
	}
	return tableMap
}

func transformColumns(columns []*Column) map[string]*[]*Column {
	columnMap := make(map[string]*[]*Column, 0)
	for _, c := range columns {
		cs, have := columnMap[c.Table]
		if have {
			*cs = append(*cs, c)
		} else {
			csp := make([]*Column, 1)
			csp[0] = c
			columnMap[c.Table] = &csp
		}
	}
	return columnMap
}

func setTableColumns(tableMap map[string]*Table, columnMap map[string]*[]*Column) {
	for k, v := range tableMap {
		v.Columns = append(v.Columns, *columnMap[k]...)
	}
}
