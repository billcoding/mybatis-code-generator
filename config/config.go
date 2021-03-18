package config

type Configuration struct {
	OutputDir     string
	Verbose       bool
	IncludeTables []string
	ExcludeTables []string
	Global        *GlobalConfiguration
	Entity        *EntityConfiguration
	Mapper        *MapperConfiguration
	Repository    *RepositoryConfiguration
	XML           *XMLConfiguration
}

type GlobalConfiguration struct {
	Author           string
	Date             bool
	DateLayout       string
	Copyright        bool
	CopyrightContent string
	Website          bool
	WebsiteContent   string
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
	PKG              string
	NamePrefix       string
	NameSuffix       string
	Comment          bool
	MapperAnnotation bool
	MybatisPlus      bool
	TK               bool
}

type RepositoryConfiguration struct {
	PKG                  string
	NamePrefix           string
	NameSuffix           string
	Comment              bool
	RepositoryAnnotation bool
}

type XMLConfiguration struct {
	Dir     string
	Comment bool
}
