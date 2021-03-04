package model

type Field struct {
	Name             string
	Type             string
	Comment          bool
	ColumnAnnotation bool
	Column           *Column
}
