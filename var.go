package main

import (
	. "github.com/billcoding/mybatis-code-generator/config"
	"os"
)

var CFG = &Configuration{
	OutputDir:     "",
	IncludeTables: []string{"purchase_in_log"},
	ExcludeTables: make([]string, 0),
	Global: &GlobalConfiguration{
		Author:           "bill",
		Date:             true,
		DateLayout:       "2006-01-02",
		Copyright:        true,
		CopyrightContent: "Mybatis code generator written by Golang",
	},
	Entity: &EntityConfiguration{
		PKG:                      "entityxx",
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
		PKG:                   "mapperxx",
		TableToMapperStrategy: UnderlineToUpper,
		MapperNamePrefix:      "",
		MapperNameSuffix:      "Mapper",
		Extend:                false,
		Extends:               make([]string, 0),
		MybatisPlus:           false,
		Comment:               true,
		MapperAnnotation:      true,
	},
	XML: &XMLConfiguration{
		Dir:                "xmlxx",
		TableToXMLStrategy: UnderlineToUpper,
		MapperPrefixes:     make([]string, 0),
		MapperSuffixes:     make([]string, 0),
	},
}

func init() {
	if CFG.OutputDir == "" {
		exec, err := os.Getwd()
		if err != nil {
			panic(err)
		}
		CFG.OutputDir = exec
	}
}
