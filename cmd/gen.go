package cmd

import (
	"fmt"
	. "github.com/billcoding/mybatis-code-generator/bundle"
	. "github.com/billcoding/mybatis-code-generator/config"
	. "github.com/billcoding/mybatis-code-generator/generator"
	"github.com/spf13/cobra"
	"os"
	"strings"
)

var genCmd = &cobra.Command{
	Use:     "gen",
	Aliases: []string{"g", "generate"},
	Short:   "Generate MyBatis Java files",
	Long: `Generate MyBatis Java files.
Simply type mybatis-code-generator help gen for full details.`,
	Example: `mybatis-code-generator gen -d "root:123@tcp(127.0.0.1:3306)/test" -D "database"
mybatis-code-generator gen -d "root:123@tcp(127.0.0.1:3306)/test" -D "database" -o "/to/path" 
mybatis-code-generator gen -d "root:123@tcp(127.0.0.1:3306)/test" -D "database" -au "bigboss" -o "/to/path"`,
	Run: func(cmd *cobra.Command, args []string) {
		CFG.Verbose = verbose

		if dsn == "" {
			_, _ = fmt.Fprintln(os.Stderr, "The DSN is required")
			return
		}

		if database == "" {
			_, _ = fmt.Fprintln(os.Stderr, "The Database name is required")
			return
		}

		if !entity {
			_, _ = fmt.Fprintln(os.Stderr, "Nothing do...")
			return
		}

		Init(dsn)
		setCFG()

		tableList := Tables(database, CFG)
		columnList := Columns(database)
		tableMap := TransformTables(tableList)
		columnMap := TransformColumns(columnList)
		SetTableColumns(tableMap, columnMap)
		generators := make([]Generator, 0)

		entityGenerators := GetEntityGenerators(CFG, tableMap)
		generators = append(generators, entityGenerators...)

		if mapper {
			mapperGenerators := GetMapperGenerators(CFG, entityGenerators)
			generators = append(generators, mapperGenerators...)
			if xml {
				xmlGenerators := GetXMLGenerators(CFG, mapperGenerators)
				generators = append(generators, xmlGenerators...)
			}
		}

		if repository {
			repositoryGenerators := GetRepositoryGenerators(CFG, entityGenerators)
			generators = append(generators, repositoryGenerators...)
		}

		for _, g := range generators {
			g.Generate()
		}

	},
}

func init() {
	genCmd.PersistentFlags().StringVarP(&outputDir, "output-dir", "o", "", "The output dir")
	genCmd.PersistentFlags().StringVarP(&dsn, "dsn", "d", "", "The MySQL DSN")
	genCmd.PersistentFlags().StringVarP(&database, "db", "D", "", "The Database name")
	genCmd.PersistentFlags().StringVarP(&includeTable, "include-table", "I", "", "The include table names[table_a,table_b]")
	genCmd.PersistentFlags().StringVar(&author, "author", "bill", "The file copyright author")
	genCmd.PersistentFlags().BoolVarP(&verbose, "verbose", "V", false, "Print verbose output")

	genCmd.PersistentFlags().BoolVarP(&entity, "entity", "e", true, "Generate Entity Java file")
	genCmd.PersistentFlags().StringVarP(&entityPKG, "entity-pkg", "E", "entity", "The Entity package")
	genCmd.PersistentFlags().BoolVar(&entityTable2EntityDefault, "table2entity-default", false, "The Table to Entity name strategy: default")
	genCmd.PersistentFlags().BoolVar(&entityTable2EntityFirstLetterUpper, "table2entity-first-letter-upper", false, "The Table to Entity name strategy: FirstLetterUpper")
	genCmd.PersistentFlags().BoolVar(&entityTable2EntityUnderlineToCamel, "table2entity-underline-to-camel", false, "The Table to Entity name strategy: UnderlineToCamel")
	genCmd.PersistentFlags().BoolVar(&entityTable2EntityUnderlineToUpper, "table2entity-underline-to-upper", true, "The Table to Entity name strategy: UnderlineToUpper")

	genCmd.PersistentFlags().BoolVar(&entityColumn2FieldDefault, "column2field-default", false, "The column to field name strategy: default")
	genCmd.PersistentFlags().BoolVar(&entityColumn2FieldFirstLetterUpper, "column2field-first-letter-upper", false, "The column to field name strategy: FirstLetterUpper")
	genCmd.PersistentFlags().BoolVar(&entityColumn2FieldUnderlineToCamel, "column2field-underline-to-camel", true, "The column to field name strategy: UnderlineToCamel")
	genCmd.PersistentFlags().BoolVar(&entityColumn2FieldUnderlineToUpper, "column2field-underline-to-upper", false, "The column to field name strategy: UnderlineToUpper")

	genCmd.PersistentFlags().BoolVar(&entityComment, "entity-comment", true, "Generate Entity comment")
	genCmd.PersistentFlags().BoolVar(&entityFieldComment, "entity-field-comment", true, "Generate Entity field comment")
	genCmd.PersistentFlags().BoolVar(&entityLombok, "entity-lombok", true, "Generate Entity lombok")
	genCmd.PersistentFlags().BoolVar(&entityLombokData, "entity-lombok-data", true, "Generate lombok @Data for Entity")
	genCmd.PersistentFlags().BoolVar(&entityLombokNoArgsCtr, "entity-lombok-no-args-ctr", true, "Generate lombok @NoArgsConstructor for Entity")
	genCmd.PersistentFlags().BoolVar(&entityLombokAllArgsCtr, "entity-lombok-all-args-ctr", true, "Generate lombok @AllArgsConstructor for Entity")
	genCmd.PersistentFlags().BoolVar(&entityLombokBuilder, "entity-lombok-builder", true, "Generate lombok @Builder for Entity")

	genCmd.PersistentFlags().StringVar(&entityImplements, "entity-implement", "", "The Entity implements interfaces")
	genCmd.PersistentFlags().StringVar(&entityExtends, "entity-extends", "", "The Entity extends class")
	genCmd.PersistentFlags().StringVar(&entityClassPrefix, "entity-class-prefix", "", "The Entity class body prefix")
	genCmd.PersistentFlags().StringVar(&entityClassSuffix, "entity-class-suffix", "", "The Entity class body suffix")

	genCmd.PersistentFlags().BoolVar(&entityAnnotation, "entity-annotation", true, "Generate @Entity for Entity")
	genCmd.PersistentFlags().BoolVar(&entityTableAnnotation, "entity-table-annotation", true, "Generate @Table for Entity")
	genCmd.PersistentFlags().BoolVar(&entityIdAnnotation, "entity-id-annotation", true, "Generate @Id for Entity field")
	genCmd.PersistentFlags().BoolVar(&entityColumnAnnotation, "entity-column-annotation", true, "Generate @Column for Entity field")

	genCmd.PersistentFlags().BoolVarP(&mapper, "mapper", "m", true, "Generate Mapper interface")
	genCmd.PersistentFlags().StringVarP(&mapperPKG, "mapper-pkg", "M", "mapper", "The Mapper interface package")
	genCmd.PersistentFlags().StringVar(&mapperNamePrefix, "mapper-name-prefix", "", "The Mapper name prefix")
	genCmd.PersistentFlags().StringVar(&mapperNameSuffix, "mapper-name-suffix", "Mapper", "The Mapper name suffix")
	genCmd.PersistentFlags().BoolVar(&mapperMybatis, "mapper-mybatis", true, "Generate Mapper Mybatis support")
	genCmd.PersistentFlags().BoolVar(&mapperTK, "mapper-tk", false, "Generate Mapper TK support")
	genCmd.PersistentFlags().BoolVar(&mapperComment, "mapper-comment", true, "Generate Mapper comment")
	genCmd.PersistentFlags().BoolVar(&mapperAnnotation, "mapper-annotation", true, "Generate @Mapper for Mapper")

	genCmd.PersistentFlags().BoolVarP(&repository, "repository", "r", false, "Generate Repository interface")
	genCmd.PersistentFlags().StringVarP(&repositoryPKG, "repository-pkg", "R", "repository", "The Repository interface package")
	genCmd.PersistentFlags().StringVar(&repositoryNamePrefix, "repository-name-prefix", "", "The Repository name prefix")
	genCmd.PersistentFlags().StringVar(&repositoryNameSuffix, "repository-name-suffix", "Repository", "The Repository name suffix")
	genCmd.PersistentFlags().BoolVar(&repositoryComment, "repository-comment", true, "Generate Repository comment")
	genCmd.PersistentFlags().BoolVar(&repositoryAnnotation, "repository-annotation", true, "Generate @Repository for Repository")

	genCmd.PersistentFlags().BoolVarP(&xml, "xml", "x", true, "Generate Mapper XML files")
	genCmd.PersistentFlags().StringVarP(&xmlDir, "xml-dir", "X", "xml", "The Mapper XML Dir")
	genCmd.PersistentFlags().BoolVar(&xmlComment, "xml-comment", true, "Generate Mapper XML comment")

	rootCmd.AddCommand(genCmd)
}

var (
	outputDir    = ""
	dsn          = ""
	database     = ""
	includeTable = ""
	author       = ""
	verbose      = false

	entity                             = true
	entityPKG                          = "entity"
	entityTable2EntityDefault          = false
	entityTable2EntityFirstLetterUpper = false
	entityTable2EntityUnderlineToCamel = false
	entityTable2EntityUnderlineToUpper = true

	entityColumn2FieldDefault          = false
	entityColumn2FieldFirstLetterUpper = false
	entityColumn2FieldUnderlineToCamel = true
	entityColumn2FieldUnderlineToUpper = false

	entityComment          = true
	entityFieldComment     = true
	entityLombok           = true
	entityLombokData       = true
	entityLombokNoArgsCtr  = true
	entityLombokAllArgsCtr = true
	entityLombokBuilder    = true

	entityImplements       = ""
	entityExtends          = ""
	entityClassPrefix      = ""
	entityClassSuffix      = ""
	entityAnnotation       = true
	entityTableAnnotation  = true
	entityIdAnnotation     = true
	entityColumnAnnotation = true

	mapper           = true
	mapperPKG        = "mapper"
	mapperNamePrefix = ""
	mapperNameSuffix = "Mapper"
	mapperMybatis    = true
	mapperTK         = false
	mapperComment    = true
	mapperAnnotation = true

	repository           = true
	repositoryPKG        = "repository"
	repositoryNamePrefix = ""
	repositoryNameSuffix = "Repository"
	repositoryComment    = true
	repositoryAnnotation = true

	xml        = true
	xmlDir     = "xml"
	xmlComment = true
)

var CFG = &Configuration{
	OutputDir:     "",
	Verbose:       false,
	IncludeTables: make([]string, 0),
	Global: &GlobalConfiguration{
		Author:           "bill",
		Date:             true,
		DateLayout:       "2006-01-02",
		Copyright:        true,
		CopyrightContent: "Mybatis code generator written in Golang",
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
		NamePrefix:       "",
		NameSuffix:       "Mapper",
		MybatisPlus:      true,
		Comment:          true,
		MapperAnnotation: true,
	},
	Repository: &RepositoryConfiguration{
		PKG:                  "repository",
		NamePrefix:           "",
		NameSuffix:           "Repository",
		Comment:              true,
		RepositoryAnnotation: true,
	},
	XML: &XMLConfiguration{
		Dir:     "xml",
		Comment: true,
	},
}

func setCFG() {
	if outputDir != "" {
		CFG.OutputDir = outputDir
	}
	if CFG.OutputDir == "" {
		exec, err := os.Getwd()
		if err != nil {
			panic(err)
		}
		CFG.OutputDir = exec
	}
	if includeTable != "" {
		CFG.IncludeTables = strings.Split(includeTable, ",")
	}

	if author != "" {
		CFG.Global.Author = author
	}

	{
		if entityPKG != "" {
			CFG.Entity.PKG = entityPKG
		}

		switch {
		case entityTable2EntityUnderlineToUpper:
			CFG.Entity.TableToEntityStrategy = UnderlineToUpper
		case entityTable2EntityUnderlineToCamel:
			CFG.Entity.TableToEntityStrategy = UnderlineToCamel
		case entityTable2EntityFirstLetterUpper:
			CFG.Entity.TableToEntityStrategy = FirstLetterUpper
		case entityTable2EntityDefault:
			CFG.Entity.TableToEntityStrategy = Default
		}

		switch {
		case entityColumn2FieldUnderlineToUpper:
			CFG.Entity.ColumnToFieldStrategy = UnderlineToUpper
		case entityColumn2FieldUnderlineToCamel:
			CFG.Entity.ColumnToFieldStrategy = UnderlineToCamel
		case entityColumn2FieldFirstLetterUpper:
			CFG.Entity.ColumnToFieldStrategy = FirstLetterUpper
		case entityColumn2FieldDefault:
			CFG.Entity.ColumnToFieldStrategy = Default
		}

		CFG.Entity.Comment = entityComment
		CFG.Entity.FieldComment = entityFieldComment
		CFG.Entity.Lombok = entityLombok
		CFG.Entity.LombokData = entityLombokData
		CFG.Entity.LombokNoArgsConstructor = entityLombokNoArgsCtr
		CFG.Entity.LombokAllArgsConstructor = entityLombokAllArgsCtr
		CFG.Entity.LombokBuilder = entityLombokBuilder

		if entityImplements != "" {
			CFG.Entity.Implement = true
			CFG.Entity.Implements = strings.Split(entityImplements, ",")
		}
		if entityExtends != "" {
			CFG.Entity.Extend = true
			CFG.Entity.Extends = entityExtends
		}
		if entityClassPrefix != "" {
			CFG.Entity.EntityClassPrefixes = strings.Split(entityClassPrefix, ",")
		}
		if entityClassSuffix != "" {
			CFG.Entity.EntityClassSuffixes = strings.Split(entityClassSuffix, ",")
		}
		CFG.Entity.EntityAnnotation = entityAnnotation
		CFG.Entity.TableAnnotation = entityTableAnnotation
		CFG.Entity.IdAnnotation = entityIdAnnotation
		CFG.Entity.ColumnAnnotation = entityColumnAnnotation
	}

	{
		if mapperPKG != "" {
			CFG.Mapper.PKG = mapperPKG
		}
		CFG.Mapper.NamePrefix = mapperNamePrefix
		CFG.Mapper.NameSuffix = mapperNameSuffix
		CFG.Mapper.MybatisPlus = mapperMybatis
		CFG.Mapper.TK = mapperTK
		if mapperTK {
			CFG.Mapper.MybatisPlus = false
		} else if mapperMybatis {
			CFG.Mapper.TK = false
		}
		CFG.Mapper.Comment = mapperComment
		CFG.Mapper.MapperAnnotation = mapperAnnotation
	}

	{
		if repositoryPKG != "" {
			CFG.Repository.PKG = repositoryPKG
		}
		CFG.Repository.NamePrefix = repositoryNamePrefix
		CFG.Repository.NameSuffix = repositoryNameSuffix
		CFG.Repository.Comment = repositoryComment
		CFG.Repository.RepositoryAnnotation = repositoryAnnotation
	}

	{
		if xmlDir != "" {
			CFG.XML.Dir = xmlDir
		}
		CFG.XML.Comment = xmlComment
	}
}
