package model

type Entity struct {
	PKG     string
	PKGName string
	Name    string
	Comment bool
	Table   *Table
	Fields  []*Field
	HaveId  bool
	Id      *Field
}
