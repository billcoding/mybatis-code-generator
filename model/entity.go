package model

type Entity struct {
	PKG     string
	PKGName string
	Name    string
	Table   *Table
	Fields  []*Field
	HaveId  bool
	Id      *Field
}
