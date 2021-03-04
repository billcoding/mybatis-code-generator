package config

type Configuration struct {
	Global GlobalConfiguration
	Entity EntityConfiguration
	Mapper MapperConfiguration
	XML    XMLConfiguration
}

type GlobalConfiguration struct {
	Author     string
	Date       bool
	DateLayout string
}

type EntityConfiguration struct {
	Entity                   bool
	PKG                      string
	TableToEntityStrategy    StrategyType
	ColumnToFieldStrategy    StrategyType
	Copyright                bool
	CopyrightContent         string
	Comment                  bool
	FieldComment             bool
	Lombok                   bool
	LombokData               bool
	LombokNoArgsConstructor  bool
	LombokAllArgsConstructor bool
	LombokBuilder            bool
	Implement                bool
	Implements               []string
	Extend                   bool
	Extends                  string
	EntityClassPrefixes      []string
	EntityClassSuffixes      []string
	EntityAnnotation         bool
	TableAnnotation          bool
	IdAnnotation             bool
	ColumnAnnotation         bool
}

type MapperConfiguration struct {
	Mapper                bool
	PKG                   string
	TableToMapperStrategy StrategyType
	MapperNamePrefix      string
	MapperNameSuffix      string
	Extend                bool
	Extends               []string
	MybatisPlus           bool
	Comment               bool
}

type XMLConfiguration struct {
	XML                bool
	Dir                string
	TableToXMLStrategy StrategyType
}
