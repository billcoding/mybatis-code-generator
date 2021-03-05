package generator

import (
	"fmt"
	. "github.com/billcoding/mybatis-code-generator/config"
	. "github.com/billcoding/mybatis-code-generator/model"
	. "github.com/billcoding/mybatis-code-generator/tpl"
	. "github.com/billcoding/mybatis-code-generator/util"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"
)

var xmlGeneratorLogger = log.New(os.Stdout, "[XMLGenerator]", log.LstdFlags)

type XMLGenerator struct {
	C       *Configuration
	XML     *XML
	Content string
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
	xg.generateContent()
	xg.generateFile()
}

func (xg *XMLGenerator) generateContent() {
	xml := ExecuteTpl(XMLTpl(), map[string]interface{}{
		"XML":    xg.XML,
		"Config": xg.C,
		"Extra": map[string]interface{}{
			"Date": time.Now().Format(xg.C.Global.DateLayout),
		},
	})
	var buffer strings.Builder
	_, _ = io.WriteString(&buffer, xml)
	xg.Content = buffer.String()
	if xg.C.Verbose {
		xmlGeneratorLogger.Println(fmt.Sprintf("[generateContent] for entity[%s]", xg.XML.Mapper.Entity.Name))
	}
}

func (xg *XMLGenerator) generateFile() {
	paths := make([]string, 0)
	paths = append(paths, xg.C.OutputDir)
	paths = append(paths, xg.C.XML.Dir)
	paths = append(paths, xg.XML.Mapper.Name)
	xmlFileName := filepath.Join(paths...) + ".xml"
	dir := filepath.Dir(xmlFileName)
	_ = os.MkdirAll(dir, 0700)
	_ = os.WriteFile(xmlFileName, []byte(xg.Content), 0700)
	if xg.C.Verbose {
		xmlGeneratorLogger.Println(fmt.Sprintf("[generateFile] for entity[%s], saved as [%s]", xg.XML.Mapper.Entity.Name, xmlFileName))
	}
}
