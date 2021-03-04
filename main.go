package main

import (
	"flag"
	"fmt"
	. "github.com/billcoding/gobatis"
	. "github.com/billcoding/mybatis-code-generator/generator"
	. "github.com/billcoding/mybatis-code-generator/model"
)

var (
	help      = flag.Bool("h", false, "The help info")
	outputDir = flag.String("o", "", "The output dir")

	dsn           = flag.String("dsn", "", "The MySQL DSN")
	includeTables = flag.String("git", "", "The include table names")
	excludeTables = flag.String("get", "", "The exclude table names")

	author    = flag.String("au", "bill", "The file copyright author")
	entityPKG = flag.String("ep", "entity", "The entity PKG")

	entityTableToEntityStrategy    = flag.Int("etes", 3, "The table to entity name strategy[0: None,1: OnlyFirstLetterUpper,2: UnderlineToCamel,3: UnderlineToUpper]")
	entityColumnToFieldStrategy    = flag.Int("ecfs", 2, "The column to field name strategy[0: None,1: OnlyFirstLetterUpper,2: UnderlineToCamel,3: UnderlineToUpper]")
	entityComment                  = flag.Bool("ec", true, "The entity comment generated?")
	entityFieldComment             = flag.Bool("efc", true, "The entity field comment generated?")
	entityLombok                   = flag.Bool("el", true, "The entity field lombok plugin generated?")
	entityLombokData               = flag.Bool("eld", true, "The entity field lombok plugin @Date generated?")
	entityLombokNoArgsConstructor  = flag.Bool("elnac", true, "The entity field lombok plugin @NoArgsConstructor generated?")
	entityLombokAllArgsConstructor = flag.Bool("elaac", true, "The entity field lombok plugin @AllArgsConstructor generated?")
	entityLombokBuilder            = flag.Bool("ellb", true, "The entity field lombok plugin @Builder generated?")

	entityImplements  = flag.String("ei", "", "The entity implements interfaces")
	entityExtends     = flag.String("ee", "", "The entity extends class")
	entityClassPrefix = flag.String("ecp", "", "The entity class body prefix")
	entityClassSuffix = flag.String("ecs", "", "The entity class body suffix")

	entityAnnotation       = flag.Bool("ea", true, "The entity @Entity generated?")
	entityTableAnnotation  = flag.Bool("eta", true, "The entity @Table generated?")
	entityIdAnnotation     = flag.Bool("eia", true, "The entity @Id generated?")
	entityColumnAnnotation = flag.Bool("eca", true, "The entity @Column generated?")
)

var (
	SelectTableListSelectMapper       *SelectMapper
	SelectTableColumnListSelectMapper *SelectMapper
)

func initBatis() {
	Ba.DSN(*dsn)

	Default().AddRaw(tableXML)

	SelectTableListSelectMapper = NewHelperWithBatis(Ba, "table", "SelectTableList").Select()
	SelectTableColumnListSelectMapper = NewHelperWithBatis(Ba, "table", "SelectTableColumnList").Select()
}

func printUsage() {
	fmt.Printf(`Usage:

mybatis-code-generator -OPTIONS

Examples:

mybatis-code-generator -dsn "root:123@tcp(127.0.0.1:3306)/test" -au "bigboss" -o "/to/path"

Supports options:
`)
	flag.PrintDefaults()
}
func main() {
	flag.Parse()
	if *help {
		printUsage()
		return
	}

	if *dsn == "" {
		printUsage()
		return
	}

	initBatis()

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
