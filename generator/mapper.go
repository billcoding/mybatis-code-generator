package generator

import (
	. "github.com/billcoding/mybatis-code-generator/config"
	. "github.com/billcoding/mybatis-code-generator/model"
	. "github.com/billcoding/mybatis-code-generator/tpl"
	. "github.com/billcoding/mybatis-code-generator/util"
	"io"
	"os"
	"path/filepath"
	"strings"
	"time"
)

type MapperGenerator struct {
	C             *Configuration
	Mapper        *Mapper
	MapperContent string
}

func (mg *MapperGenerator) Init(e *Entity) {
	mg.Mapper = &Mapper{
		PKG:    mg.C.Mapper.PKG,
		Entity: e,
	}
	mg.Mapper.Name = mg.C.Mapper.MapperNamePrefix + mg.Mapper.Entity.Name + mg.C.Mapper.MapperNameSuffix
	mg.Mapper.PKGName = mg.C.Mapper.PKG + "." + mg.Mapper.Name
}

func (mg *MapperGenerator) Generate() {
	mg.generateMapperContent()
	mg.generateFile()
}

func (mg *MapperGenerator) generateMapperContent() {
	class := ExecuteTpl(MapperTpl(), map[string]interface{}{
		"Mapper": mg.Mapper,
		"Config": mg.C,
		"Extra": map[string]interface{}{
			"Date": time.Now().Format(mg.C.Global.DateLayout),
		},
	})
	var buffer strings.Builder
	_, _ = io.WriteString(&buffer, class)
	mg.MapperContent = buffer.String()
}

func (mg *MapperGenerator) generateFile() {
	paths := make([]string, 0)
	paths = append(paths, mg.C.OutputDir)
	paths = append(paths, strings.Split(mg.Mapper.PKGName, ".")...)
	mapperFileName := filepath.Join(paths...) + ".java"
	dir := filepath.Dir(mapperFileName)
	_ = os.MkdirAll(dir, 0700)
	_ = os.WriteFile(mapperFileName, []byte(mg.MapperContent), 0700)
}
