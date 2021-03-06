package main

import (
	"flag"
	. "github.com/billcoding/mybatis-code-generator/generator"
)

func main() {
	flag.Parse()
	if *help {
		printUsage()
		return
	}

	if *version {
		printVersion()
		return
	}

	CFG.Verbose = *verbose

	if *dsn == "" {
		logger.Println("The DSN is required")
		return
	}

	if *database == "" {
		logger.Println("The Database name is required")
		return
	}

	initBatis()

	setCFG()

	tableList := tables(*database, CFG)
	columnList := columns(*database)
	tableMap := transformTables(tableList)
	columnMap := transformColumns(columnList)
	setTableColumns(tableMap, columnMap)
	generators := make([]Generator, 0)

	if !*entity {
		if CFG.Verbose {
			logger.Println("Nothing do...")
		}
		return
	}

	entityGenerators := getEntityGenerators(tableMap)
	generators = append(generators, entityGenerators...)

	if *mapper {
		mapperGenerators := getMapperGenerators(entityGenerators)
		generators = append(generators, mapperGenerators...)
		if *xml {
			xmlGenerators := getXMLGenerators(mapperGenerators)
			generators = append(generators, xmlGenerators...)
		}
	}

	if *repository {
		repositoryGenerators := getRepositoryGenerators(entityGenerators)
		generators = append(generators, repositoryGenerators...)
	}

	for _, generator := range generators {
		generator.Generate()
	}

}
