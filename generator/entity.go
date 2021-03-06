package generator

import (
	"fmt"
	. "github.com/billcoding/mybatis-code-generator/config"
	. "github.com/billcoding/mybatis-code-generator/model"
	"github.com/billcoding/mybatis-code-generator/tpl"
	. "github.com/billcoding/mybatis-code-generator/util"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"
)

var entityGeneratorLogger = log.New(os.Stdout, "[EntityGenerator]", log.LstdFlags)

type EntityGenerator struct {
	C       *Configuration
	Table   *Table
	Entity  *Entity
	Body    string
	Imports []string
}

func (eg *EntityGenerator) Generate() {
	eg.generateBody()
	eg.generateFile()
}

func (eg *EntityGenerator) Init() *EntityGenerator {
	eg.Entity = &Entity{
		PKG:     eg.C.Entity.PKG,
		Table:   eg.Table,
		Fields:  make([]*Field, 0),
		Comment: eg.C.Entity.Comment,
		Ids:     make([]*Field, 0),
	}
	eg.Entity.Name = ConvertString(eg.Table.Name, eg.C.Entity.TableToEntityStrategy)
	eg.Entity.PKGName = eg.Entity.PKG + "." + eg.Entity.Name
	importMap := make(map[string]struct{}, 0)
	for _, column := range eg.Table.Columns {
		field := &Field{
			Name:             ConvertString(column.Name, eg.C.Entity.ColumnToFieldStrategy),
			Type:             MysqlToJavaTypes[column.Type],
			Column:           column,
			Comment:          eg.C.Entity.FieldComment,
			ColumnAnnotation: eg.C.Entity.ColumnAnnotation,
			IdAnnotation:     eg.C.Entity.IdAnnotation,
		}
		if column.ColumnKey == "PRI" {
			eg.Entity.HaveId = true
			eg.Entity.Ids = append(eg.Entity.Ids, field)
		} else {
			eg.Entity.Fields = append(eg.Entity.Fields, field)
		}
		if importPKG, have := JavaTypePKGs[field.Type]; have {
			importMap[importPKG] = struct{}{}
		}
	}
	eg.Imports = make([]string, 0)
	for k := range importMap {
		eg.Imports = append(eg.Imports, k)
	}
	if eg.C.Entity.EntityAnnotation {
		eg.Imports = append(eg.Imports, "javax.persistence.Entity")
	}
	if eg.C.Entity.TableAnnotation {
		eg.Imports = append(eg.Imports, "javax.persistence.Table")
	}
	if eg.C.Entity.ColumnAnnotation {
		eg.Imports = append(eg.Imports, "javax.persistence.Column")
	}
	if eg.C.Entity.IdAnnotation {
		eg.Imports = append(eg.Imports, "javax.persistence.Id")
	}
	if eg.C.Entity.Lombok {
		if eg.C.Entity.LombokAllArgsConstructor {
			eg.Imports = append(eg.Imports, "lombok.AllArgsConstructor")
		}
		if eg.C.Entity.LombokBuilder {
			eg.Imports = append(eg.Imports, "lombok.Builder")
		}
		if eg.C.Entity.LombokData {
			eg.Imports = append(eg.Imports, "lombok.Data")
		}
		if eg.C.Entity.LombokNoArgsConstructor {
			eg.Imports = append(eg.Imports, "lombok.NoArgsConstructor")
		}
	}
	if !eg.Entity.HaveId {
		panic(fmt.Sprintf("Table of [%s] required at least one PRI column", eg.Entity.Table.Name))
	}
	return eg
}

func (eg *EntityGenerator) generateBody() {
	eg.Body = ExecuteTpl(tpl.EntityTpl(), map[string]interface{}{
		"Entity": eg.Entity,
		"Config": eg.C,
		"Extra": map[string]interface{}{
			"Date": time.Now().Format(eg.C.Global.DateLayout),
		},
		"Imports": eg.Imports,
	})
	if eg.C.Verbose {
		entityGeneratorLogger.Println(fmt.Sprintf("[generateBody] for entity[%s]", eg.Entity.Name))
	}
}

func (eg *EntityGenerator) generateFile() {
	paths := make([]string, 0)
	paths = append(paths, eg.C.OutputDir)
	paths = append(paths, strings.Split(eg.Entity.PKGName, ".")...)
	fileName := filepath.Join(paths...) + ".java"
	dir := filepath.Dir(fileName)
	_ = os.MkdirAll(dir, 0700)
	_ = os.WriteFile(fileName, []byte(eg.Body), 0700)
	if eg.C.Verbose {
		entityGeneratorLogger.Println(fmt.Sprintf("[generateFile] for entity[%s], saved as [%s]", eg.Entity.Name, fileName))
	}
}
