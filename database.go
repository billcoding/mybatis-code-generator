package main

import (
	. "github.com/billcoding/mybatis-code-generator/mapper"
	. "github.com/billcoding/mybatis-code-generator/model"
)

func tables(database string) []*Table {
	tableList := SelectTableListSelectMapper.Prepare(database).Exec().List(new(Table))
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
