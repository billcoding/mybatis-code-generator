package generator

import (
	. "github.com/billcoding/mybatis-code-generator/config"
	. "github.com/billcoding/mybatis-code-generator/model"
	"github.com/billcoding/mybatis-code-generator/tpl"
	. "github.com/billcoding/mybatis-code-generator/util"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"
)

var logger = log.New(os.Stdout, "[EntityGenerator]", log.LstdFlags)

type EntityGenerator struct {
	C      *Configuration
	Table  *Table
	Entity *Entity
	Class  string
}

func (eg *EntityGenerator) Generate() {
	eg.generateEntity()
	eg.generateClass()
	eg.generateFile()
}

func (eg *EntityGenerator) generateEntity() {
	eg.Entity = &Entity{
		PKG:    eg.C.Entity.PKG,
		Table:  eg.Table,
		Fields: make([]*Field, 0),
	}
	eg.setEntityParams()
}

func (eg *EntityGenerator) setEntityParams() {
	eg.Entity.Name = ConvertString(eg.Table.Name, eg.C.Entity.TableToEntityStrategy)
	eg.Entity.PKGName = eg.Entity.PKG + eg.Entity.Name
	for _, column := range eg.Table.Columns {
		field := &Field{
			Name:   ConvertString(column.Name, eg.C.Entity.ColumnToFieldStrategy),
			Type:   MysqlToJavaTypes[column.Type],
			Column: column,
		}
		if column.ColumnKey == "PRI" {
			eg.Entity.HaveId = true
			eg.Entity.Id = field
		} else {
			eg.Entity.Fields = append(eg.Entity.Fields, field)
		}
	}
}

func (eg *EntityGenerator) generateClass() string {
	class := ExecuteTpl(tpl.EntityTpl(), map[string]interface{}{
		"Entity": eg.Entity,
		"Config": eg.C,
	})
	var buffer strings.Builder
	io.WriteString(&buffer, class)
	return class
}

func (eg *EntityGenerator) generateFile() {
	entityFileName := filepath.Join(strings.Split(eg.Entity.PKGName, ".")...)
	f, err := os.OpenFile(entityFileName, os.O_CREATE, 0700)
	if err != nil {
		panic(err)
	}
	io.WriteString(f, eg.Class)
}
