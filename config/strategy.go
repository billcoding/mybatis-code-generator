package config

type StrategyType string

var (
	// 无，原名称copy					abc => abc
	None StrategyType
	// 仅首字母大写					abc => Abc
	OnlyFirstLetterUpper StrategyType
	// 下划线转驼峰（首字母小写）		a_b_c => aBC
	UnderlineToCamel StrategyType
	// 下划线转大写					a_b_c => ABC
	UnderlineToUpper StrategyType
)
