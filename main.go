package main

import (
	. "github.com/billcoding/mybatis-code-generator/generator"
	. "github.com/billcoding/mybatis-code-generator/model"
)

func main() {
	tableList := tables("nterp_wms", CFG)
	columnList := columns("nterp_wms", CFG)
	tableMap := transformTables(tableList)
	columnMap := transformColumns(columnList)
	setTableColumns(tableMap, columnMap)
	generators := make([]Generator, 0)

	entityGenerators := GetEntityGenerators(tableMap)
	generators = append(generators, entityGenerators...)

	mapperGenerators := GetMapperGenerators(entityGenerators)
	generators = append(generators, mapperGenerators...)

	xmlGenerators := GetXMLGenerators(mapperGenerators)
	generators = append(generators, xmlGenerators...)

	for _, generator := range generators {
		generator.Generate()
	}
}

func GetEntityGenerators(tableMap map[string]*Table) []Generator {
	egs := make([]Generator, 0)
	for _, v := range tableMap {
		eg := &EntityGenerator{
			C:     CFG,
			Table: v,
		}
		eg.Init()
		egs = append(egs, eg)
	}
	return egs
}

func GetMapperGenerators(entityGenerators []Generator) []Generator {
	egs := make([]Generator, 0)
	for _, eg := range entityGenerators {
		mg := &MapperGenerator{
			C: CFG,
		}
		mg.Init(eg.(*EntityGenerator).Entity)
		egs = append(egs, mg)
	}
	return egs
}

func GetXMLGenerators(mapperGenerators []Generator) []Generator {
	xgs := make([]Generator, 0)
	for _, mg := range mapperGenerators {
		xml := &XMLGenerator{
			C: CFG,
		}
		xml.Init(mg.(*MapperGenerator).Mapper)
		xgs = append(xgs, xml)
	}
	return xgs
}
