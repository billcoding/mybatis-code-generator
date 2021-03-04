package main

import (
	. "github.com/billcoding/mybatis-code-generator/generator"
)

func main() {
	tableList := tables("nterp_wms")
	columnList := columns("nterp_wms")
	tableMap := transformTables(tableList)
	columnMap := transformColumns(columnList)
	setTableColumns(tableMap, columnMap)

	for _, v := range tableMap {
		eg := EntityGenerator{
			C:     CFG,
			Table: v,
		}
		eg.Generate()
	}
}
