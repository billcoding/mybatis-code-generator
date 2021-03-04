package mapper

import (
	. "github.com/billcoding/gobatis"
	. "github.com/billcoding/mybatis-code-generator/config"
)

var (
	SelectTableListSelectMapper       *SelectMapper
	SelectTableColumnListSelectMapper *SelectMapper
)

func init() {
	Default().AddRaw(tableXML)

	SelectTableListSelectMapper = NewHelperWithBatis(Ba, "table", "SelectTableList").Select()
	SelectTableColumnListSelectMapper = NewHelperWithBatis(Ba, "table", "SelectTableColumnList").Select()
}

var tableXML = `<?xml version="1.0" encoding="UTF-8"?>
<batis-mapper binding="table">
    <select id="SelectTableList">
		SELECT 
		  t.TABLE_NAME,
		  t.TABLE_COMMENT 
		FROM
		  information_schema.TABLES AS t 
		WHERE t.TABLE_SCHEMA = '{{.}}' 
    </select>
    <select id="SelectTableColumnList">
		SELECT 
		  t.*,
		  UPPER(t.COLUMN_TYPE) AS UPPER_COLUMN_TYPE 
		FROM
		  (SELECT 
			t.TABLE_NAME,
			t.COLUMN_KEY,
			t.COLUMN_NAME,
			IF(
			  LOCATE('(', t.COLUMN_TYPE) = 0,
			  t.COLUMN_TYPE,
			  SUBSTRING(
				t.COLUMN_TYPE,
				1,
				LOCATE('(', t.COLUMN_TYPE) - 1
			  )
			) AS COLUMN_TYPE,
			t.COLUMN_COMMENT 
		  FROM
			information_schema.COLUMNS AS t 
		  WHERE t.TABLE_SCHEMA = '{{.}}' 
		  ORDER BY t.TABLE_NAME ASC,
			t.ORDINAL_POSITION ASC) AS t 
    </select>
</batis-mapper>`
