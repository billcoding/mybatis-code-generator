package config

type Configuration struct {
	OutputDir     string
	IncludeTables []string
	ExcludeTables []string
	Global        *GlobalConfiguration
	Entity        *EntityConfiguration
	Mapper        *MapperConfiguration
	XML           *XMLConfiguration
}

type GlobalConfiguration struct {
	Author           string
	Date             bool
	DateLayout       string
	Copyright        bool
	CopyrightContent string
}

type EntityConfiguration struct {
	PKG                      string
	TableToEntityStrategy    StrategyType
	ColumnToFieldStrategy    StrategyType
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
	PKG                   string
	TableToMapperStrategy StrategyType
	MapperNamePrefix      string
	MapperNameSuffix      string
	Extend                bool
	Extends               []string
	MybatisPlus           bool
	Comment               bool
	MapperAnnotation      bool
}

type XMLConfiguration struct {
	Dir                string
	TableToXMLStrategy StrategyType
	MapperPrefixes     []string
	MapperSuffixes     []string
}
