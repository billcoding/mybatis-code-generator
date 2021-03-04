package main

import (
	"fmt"
	. "github.com/billcoding/mybatis-code-generator/generator"
)

func main() {
	tableList := tables("nterp_wms", CFG)
	columnList := columns("nterp_wms", CFG)
	tableMap := transformTables(tableList)
	columnMap := transformColumns(columnList)
	setTableColumns(tableMap, columnMap)
	for _, v := range tableMap {
		eg := EntityGenerator{
			C:     CFG,
			Table: v,
		}
		eg.Generate()
		fmt.Println(eg.Class)
		break
	}
}
