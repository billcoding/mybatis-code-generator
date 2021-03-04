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
)

type XMLGenerator struct {
	C          *Configuration
	XML        *XML
	XMLContent string
}

func (xg *XMLGenerator) Init(m *Mapper) {
	xg.XML = &XML{
		Mapper: m,
		ResultMap: &ResultMap{
			Items: make([]*Field, 0),
		},
	}
	xg.XML.ResultMap.HaveId = xg.XML.Mapper.Entity.HaveId
	xg.XML.ResultMap.Id = xg.XML.Mapper.Entity.Id
	for _, field := range xg.XML.Mapper.Entity.Fields {
		xg.XML.ResultMap.Items = append(xg.XML.ResultMap.Items, field)
	}
}

func (xg *XMLGenerator) Generate() {
	xg.generateXMLContent()
	xg.generateFile()
}

func (xg *XMLGenerator) generateXMLContent() {
	xml := ExecuteTpl(XMLTpl(), map[string]interface{}{
		"XML":    xg.XML,
		"Config": xg.C,
	})
	var buffer strings.Builder
	_, _ = io.WriteString(&buffer, xml)
	xg.XMLContent = buffer.String()
}

func (xg *XMLGenerator) generateFile() {
	paths := make([]string, 0)
	paths = append(paths, xg.C.OutputDir)
	paths = append(paths, xg.C.XML.Dir)
	paths = append(paths, xg.XML.Mapper.Name)
	xmlFileName := filepath.Join(paths...) + ".xml"
	dir := filepath.Dir(xmlFileName)
	_ = os.MkdirAll(dir, 0700)
	_ = os.WriteFile(xmlFileName, []byte(xg.XMLContent), 0700)
}
