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

	CFG.Verbose = *print

	if *dsn == "" {
		if CFG.Verbose {
			logger.Println("The DSN is required")
		}
		printUsage()
		return
	}

	if *database == "" {
		if CFG.Verbose {
			logger.Println("The Database name is required")
		}
		printUsage()
		return
	}

	initBatis()

	setCFG()

	tableList := tables(*database, CFG)
	columnList := columns(*database, CFG)
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

	entityGenerators := GetEntityGenerators(tableMap)
	generators = append(generators, entityGenerators...)

	if *mapper {
		mapperGenerators := GetMapperGenerators(entityGenerators)
		generators = append(generators, mapperGenerators...)
		if *xml {
			xmlGenerators := GetXMLGenerators(mapperGenerators)
			generators = append(generators, xmlGenerators...)
		}
	}

	if *repository {
		repositoryGenerators := GetRepositoryGenerators(entityGenerators)
		generators = append(generators, repositoryGenerators...)
	}

	for _, generator := range generators {
		generator.Generate()
	}

}
