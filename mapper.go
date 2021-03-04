package main

var tableXML = `<?xml version="1.0" encoding="UTF-8"?>
<batis-mapper binding="table">
    <select id="SelectTableList">
		SELECT 
		  t.TABLE_NAME,
		  IFNULL(t.TABLE_COMMENT, '') as TABLE_COMMENT
		FROM
		  information_schema.TABLES AS t 
		WHERE t.TABLE_SCHEMA = '{{.DBName}}'
		{{.Where}}
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
			IFNULL(t.COLUMN_COMMENT, '') as COLUMN_COMMENT
		  FROM
			information_schema.COLUMNS AS t 
		  WHERE t.TABLE_SCHEMA = '{{.}}' 
		  ORDER BY t.TABLE_NAME ASC,
			t.ORDINAL_POSITION ASC) AS t 
    </select>
</batis-mapper>`
