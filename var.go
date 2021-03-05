package main

import (
	"flag"
	. "github.com/billcoding/mybatis-code-generator/config"
	"log"
	"os"
	"strings"
)

var (
	version       = flag.Bool("v", false, "The version info")
	help          = flag.Bool("h", false, "The help info")
	outputDir     = flag.String("o", "", "The output dir")
	dsn           = flag.String("dsn", "root:123@tcp(127.0.0.1:3306)/test", "The MySQL DSN")
	database      = flag.String("db", "", "The Database name")
	includeTables = flag.String("git", "", "The include table names[table_a,table_b]")
	excludeTables = flag.String("get", "", "The exclude table names[table_a,table_b]")
	author        = flag.String("au", "bill", "The file copyright author")
	verbose       = flag.Bool("vb", false, "The verbose detail show?")

	entity                         = flag.Bool("e", true, "The entity enable?")
	entityPKG                      = flag.String("ep", "entity", "The entity PKG")
	entityTableToEntityStrategy    = flag.Int("etes", 3, "The table to entity name strategy[0: None,1: OnlyFirstLetterUpper,2: UnderlineToCamel,3: UnderlineToUpper]")
	entityColumnToFieldStrategy    = flag.Int("ecfs", 2, "The column to field name strategy[0: None,1: OnlyFirstLetterUpper,2: UnderlineToCamel,3: UnderlineToUpper]")
	entityComment                  = flag.Bool("ec", true, "The entity comment generated?")
	entityFieldComment             = flag.Bool("efc", true, "The entity field comment generated?")
	entityLombok                   = flag.Bool("el", true, "The entity field lombok plugin generated?")
	entityLombokData               = flag.Bool("eld", true, "The entity field lombok plugin @Date generated?")
	entityLombokNoArgsConstructor  = flag.Bool("elnac", true, "The entity field lombok plugin @NoArgsConstructor generated?")
	entityLombokAllArgsConstructor = flag.Bool("elaac", true, "The entity field lombok plugin @AllArgsConstructor generated?")
	entityLombokBuilder            = flag.Bool("ellb", true, "The entity field lombok plugin @Builder generated?")
	entityImplements               = flag.String("ei", "", "The entity implements interfaces")
	entityExtends                  = flag.String("ee", "", "The entity extends class")
	entityClassPrefix              = flag.String("ecp", "", "The entity class body prefix")
	entityClassSuffix              = flag.String("ecs", "", "The entity class body suffix")
	entityAnnotation               = flag.Bool("ea", true, "The entity @Entity generated?")
	entityTableAnnotation          = flag.Bool("eta", true, "The entity @Table generated?")
	entityIdAnnotation             = flag.Bool("eia", true, "The entity @Id generated?")
	entityColumnAnnotation         = flag.Bool("eca", true, "The entity @Column generated?")

	mapper                 = flag.Bool("m", true, "The Mapper enable?")
	mapperPKG              = flag.String("mp", "mapper", "The Mapper PKG")
	mapperNamePrefix       = flag.String("mnp", "", "The Mapper name prefix")
	mapperNameSuffix       = flag.String("mns", "Mapper", "The Mapper name suffix")
	mapperMybatis          = flag.Bool("mmb", false, "The Mapper supports Mybatis?")
	mapperComment          = flag.Bool("mc", true, "The Mapper comment?")
	mapperMapperAnnotation = flag.Bool("mma", true, "The Mapper @Mapper generated?")

	repository                     = flag.Bool("r", false, "The Repository enable?")
	repositoryPKG                  = flag.String("rp", "repository", "The Repository PKG")
	repositoryNamePrefix           = flag.String("rnp", "", "The Repository name prefix")
	repositoryNameSuffix           = flag.String("rns", "Repository", "The Repository name suffix")
	repositoryComment              = flag.Bool("rc", true, "The Repository comment?")
	repositoryRepositoryAnnotation = flag.Bool("rra", true, "The Repository @Repository generated?")

	xml              = flag.Bool("x", true, "The mapper xml enable?")
	xmlDir           = flag.String("xd", "xml", "The XML Dir generated")
	xmlComment       = flag.Bool("xc", true, "The XML comment?")
	xmlContentPrefix = flag.String("xcp", "", "The XML body prefix")
	xmlContentSuffix = flag.String("xcs", "", "The XML body suffix")
)

var logger = log.New(os.Stdout, "[mybatis-code-generator]", log.LstdFlags)

var CFG = &Configuration{
	OutputDir:     "",
	IncludeTables: make([]string, 0),
	ExcludeTables: make([]string, 0),
	Global: &GlobalConfiguration{
		Author:           "bill",
		Date:             true,
		DateLayout:       "2006-01-02",
		Copyright:        true,
		CopyrightContent: "Mybatis code generator written by Golang",
		Website:          true,
		WebsiteContent:   "https://github.com/billcoding/mybatis-code-generator",
	},
	Entity: &EntityConfiguration{
		PKG:                      "entity",
		TableToEntityStrategy:    UnderlineToUpper,
		ColumnToFieldStrategy:    UnderlineToCamel,
		Comment:                  true,
		FieldComment:             true,
		Lombok:                   true,
		LombokData:               true,
		LombokNoArgsConstructor:  true,
		LombokAllArgsConstructor: true,
		LombokBuilder:            true,
		Implement:                false,
		Implements:               make([]string, 0),
		Extend:                   false,
		Extends:                  "",
		EntityClassPrefixes:      make([]string, 0),
		EntityClassSuffixes:      make([]string, 0),
		EntityAnnotation:         true,
		TableAnnotation:          true,
		IdAnnotation:             true,
		ColumnAnnotation:         true,
	},
	Mapper: &MapperConfiguration{
		PKG:              "mapper",
		MapperNamePrefix: "",
		MapperNameSuffix: "Mapper",
		MybatisPlus:      true,
		Comment:          true,
		MapperAnnotation: true,
	},
	Repository: &RepositoryConfiguration{
		PKG:                  "repository",
		RepositoryNamePrefix: "",
		RepositoryNameSuffix: "Repository",
		Comment:              true,
		RepositoryAnnotation: true,
	},
	XML: &XMLConfiguration{
		Dir:            "xml",
		MapperPrefixes: make([]string, 0),
		MapperSuffixes: make([]string, 0),
		Comment:        true,
	},
}

func setCFG() {
	if *outputDir != "" {
		CFG.OutputDir = *outputDir
	}
	if CFG.OutputDir == "" {
		exec, err := os.Getwd()
		if err != nil {
			panic(err)
		}
		CFG.OutputDir = exec
	}
	if *includeTables != "" {
		CFG.IncludeTables = strings.Split(*includeTables, ",")
	} else if *excludeTables != "" {
		CFG.ExcludeTables = strings.Split(*excludeTables, ",")
	}

	if *author != "" {
		CFG.Global.Author = *author
	}

	if *entityPKG != "" {
		CFG.Entity.PKG = *entityPKG
	}

	entityTableToEntityStrategyMap := map[int]StrategyType{
		0: None,
		1: OnlyFirstLetterUpper,
		2: UnderlineToCamel,
		3: UnderlineToUpper,
	}
	s1, have1 := entityTableToEntityStrategyMap[*entityTableToEntityStrategy]
	if have1 {
		CFG.Entity.TableToEntityStrategy = s1
	}
	s2, have2 := entityTableToEntityStrategyMap[*entityColumnToFieldStrategy]
	if have2 {
		CFG.Entity.ColumnToFieldStrategy = s2
	}
	CFG.Entity.Comment = *entityComment
	CFG.Entity.FieldComment = *entityFieldComment
	CFG.Entity.Lombok = *entityLombok
	CFG.Entity.LombokData = *entityLombokData
	CFG.Entity.LombokNoArgsConstructor = *entityLombokNoArgsConstructor
	CFG.Entity.LombokAllArgsConstructor = *entityLombokAllArgsConstructor
	CFG.Entity.LombokBuilder = *entityLombokBuilder

	if *entityImplements != "" {
		CFG.Entity.Implement = true
		CFG.Entity.Implements = strings.Split(*entityImplements, ",")
	}
	if *entityExtends != "" {
		CFG.Entity.Extend = true
		CFG.Entity.Extends = *entityExtends
	}
	if *entityClassPrefix != "" {
		CFG.Entity.EntityClassPrefixes = strings.Split(*entityClassPrefix, ",")
	}
	if *entityClassSuffix != "" {
		CFG.Entity.EntityClassSuffixes = strings.Split(*entityClassSuffix, ",")
	}
	CFG.Entity.EntityAnnotation = *entityAnnotation
	CFG.Entity.TableAnnotation = *entityTableAnnotation
	CFG.Entity.IdAnnotation = *entityIdAnnotation
	CFG.Entity.ColumnAnnotation = *entityColumnAnnotation

	if *mapperPKG != "" {
		CFG.Mapper.PKG = *mapperPKG
	}
	CFG.Mapper.MapperNamePrefix = *mapperNamePrefix
	CFG.Mapper.MapperNameSuffix = *mapperNameSuffix
	CFG.Mapper.MybatisPlus = *mapperMybatis
	CFG.Mapper.Comment = *mapperComment
	CFG.Mapper.MapperAnnotation = *mapperMapperAnnotation

	if *repositoryPKG != "" {
		CFG.Repository.PKG = *repositoryPKG
	}
	CFG.Repository.RepositoryNamePrefix = *repositoryNamePrefix
	CFG.Repository.RepositoryNameSuffix = *repositoryNameSuffix
	CFG.Repository.Comment = *repositoryComment
	CFG.Repository.RepositoryAnnotation = *repositoryRepositoryAnnotation

	if *xmlDir != "" {
		CFG.XML.Dir = *xmlDir
	}
	CFG.XML.Comment = *xmlComment
	if *xmlContentPrefix != "" {
		CFG.XML.MapperPrefixes = strings.Split(*xmlContentPrefix, ",")
	}
	if *xmlContentSuffix != "" {
		CFG.XML.MapperSuffixes = strings.Split(*xmlContentSuffix, ",")
	}
}
