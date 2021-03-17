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
	"time"
)

var xmlGeneratorLogger = log.New(os.Stdout, "[XMLGenerator]", log.LstdFlags)

type XMLGenerator struct {
	C    *Configuration
	XML  *XML
	Body string
}

func (xg *XMLGenerator) Init(m *Mapper) {
	xg.XML = &XML{
		Mapper: m,
	}
}

func (xg *XMLGenerator) Generate() {
	xg.generateContent()
	xg.generateFile()
}

func (xg *XMLGenerator) generateContent() {
	xg.Body = ExecuteTpl(XMLTpl(), map[string]interface{}{
		"XML":    xg.XML,
		"Config": xg.C,
		"Extra": map[string]interface{}{
			"Date": time.Now().Format(xg.C.Global.DateLayout),
		},
	})
	if xg.C.Verbose {
		xmlGeneratorLogger.Println(fmt.Sprintf("[generateContent] for entity[%s]", xg.XML.Mapper.Entity.Name))
	}
}

func (xg *XMLGenerator) generateFile() {
	paths := make([]string, 0)
	paths = append(paths, xg.C.OutputDir)
	paths = append(paths, xg.C.XML.Dir)
	paths = append(paths, xg.XML.Mapper.Name)
	fileName := filepath.Join(paths...) + ".xml"
	dir := filepath.Dir(fileName)
	_ = os.MkdirAll(dir, 0700)
	_ = os.WriteFile(fileName, []byte(xg.Body), 0700)
	if xg.C.Verbose {
		xmlGeneratorLogger.Println(fmt.Sprintf("[generateFile] for entity[%s], saved as [%s]", xg.XML.Mapper.Entity.Name, fileName))
	}
}
