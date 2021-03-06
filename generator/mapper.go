package generator

import (
	"fmt"
	. "github.com/billcoding/mybatis-code-generator/config"
	. "github.com/billcoding/mybatis-code-generator/model"
	. "github.com/billcoding/mybatis-code-generator/tpl"
	. "github.com/billcoding/mybatis-code-generator/util"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"
)

var mapperGeneratorLogger = log.New(os.Stdout, "[MapperGenerator]", log.LstdFlags)

type MapperGenerator struct {
	C      *Configuration
	Mapper *Mapper
	Body   string
}

func (mg *MapperGenerator) Init(e *Entity) {
	mg.Mapper = &Mapper{
		PKG:    mg.C.Mapper.PKG,
		Entity: e,
	}
	mg.Mapper.Name = mg.C.Mapper.NamePrefix + mg.Mapper.Entity.Name + mg.C.Mapper.NameSuffix
	mg.Mapper.PKGName = mg.C.Mapper.PKG + "." + mg.Mapper.Name
}

func (mg *MapperGenerator) Generate() {
	mg.generateBody()
	mg.generateFile()
}

func (mg *MapperGenerator) generateBody() {
	mg.Body = ExecuteTpl(MapperTpl(), map[string]interface{}{
		"Mapper": mg.Mapper,
		"Config": mg.C,
		"Extra": map[string]interface{}{
			"Date": time.Now().Format(mg.C.Global.DateLayout),
		},
	})
	if mg.C.Verbose {
		mapperGeneratorLogger.Println(fmt.Sprintf("[generateBody] for entity[%s]", mg.Mapper.Entity.Name))
	}
}

func (mg *MapperGenerator) generateFile() {
	paths := make([]string, 0)
	paths = append(paths, mg.C.OutputDir)
	paths = append(paths, strings.Split(mg.Mapper.PKGName, ".")...)
	fileName := filepath.Join(paths...) + ".java"
	dir := filepath.Dir(fileName)
	_ = os.MkdirAll(dir, 0700)
	_ = os.WriteFile(fileName, []byte(mg.Body), 0700)
	if mg.C.Verbose {
		mapperGeneratorLogger.Println(fmt.Sprintf("[generateFile] for entity[%s], saved as [%s]", mg.Mapper.Entity.Name, fileName))
	}
}
