package config

var JavaTypePKGs = map[string]string{
	"BigDecimal": "java.math.BigDecimal",
	"Date":       "java.util.Date",
}

var MysqlToJavaTypes = map[string]string{
	"bit":        "Boolean",
	"tinyint":    "Byte",
	"smallint":   "Short",
	"mediumint":  "Integer",
	"int":        "Integer",
	"bigint":     "Long",
	"float":      "Float",
	"double":     "Double",
	"decimal":    "BigDecimal",
	"date":       "String",
	"time":       "String",
	"year":       "Short",
	"datetime":   "Date",
	"timestamp":  "Date",
	"char":       "String",
	"varchar":    "String",
	"tinytext":   "String",
	"mediumtext": "String",
	"text":       "String",
	"longtext":   "String",
	"tinyblob":   "Byte[]",
	"mediumblob": "Byte[]",
	"blob":       "Byte[]",
	"longblob":   "Byte[]",
}
