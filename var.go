package main

import (
	. "github.com/billcoding/mybatis-code-generator/config"
)

var CFG = &Configuration{
	IncludeTables: make([]string, 0),
	ExcludeTables: make([]string, 0),
	Global: &GlobalConfiguration{
		Author:           "bill",
		Date:             true,
		DateLayout:       "2006-01-02",
		Copyright:        true,
		CopyrightContent: "Mybatis code generator written by Golang",
	},
	Entity: &EntityConfiguration{
		Entity:                   false,
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
		Implement:                true,
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
		Mapper:                false,
		PKG:                   "mapper",
		TableToMapperStrategy: UnderlineToUpper,
		MapperNamePrefix:      "",
		MapperNameSuffix:      "Mapper",
		Extend:                false,
		Extends:               make([]string, 0),
		MybatisPlus:           false,
		Comment:               true,
	},
	XML: &XMLConfiguration{
		XML:                false,
		Dir:                "xml",
		TableToXMLStrategy: UnderlineToUpper,
	},
}
