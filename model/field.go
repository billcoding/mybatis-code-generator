package model

type Field struct {
	Name             string
	Type             string
	Comment          bool
	ColumnAnnotation bool
	IdAnnotation     bool
	Column           *Column
}
