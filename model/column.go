package model

type Column struct {
	Table     string `db:"TABLE_NAME"`
	Name      string `db:"COLUMN_NAME"`
	Type      string `db:"COLUMN_TYPE"`
	UpperType string `db:"UPPER_COLUMN_TYPE"`
	ColumnKey string `db:"COLUMN_KEY"`
	Comment   string `db:"COLUMN_COMMENT"`
}
