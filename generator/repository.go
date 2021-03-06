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

var repositoryGeneratorLogger = log.New(os.Stdout, "[RepositoryGenerator]", log.LstdFlags)

type RepositoryGenerator struct {
	C          *Configuration
	Repository *Repository
	Body       string
}

func (rg *RepositoryGenerator) Init(e *Entity) {
	rg.Repository = &Repository{
		PKG:    rg.C.Mapper.PKG,
		Entity: e,
	}
	rg.Repository.Name = rg.C.Repository.NamePrefix + rg.Repository.Entity.Name + rg.C.Repository.NameSuffix
	rg.Repository.PKGName = rg.C.Repository.PKG + "." + rg.Repository.Name
}

func (rg *RepositoryGenerator) Generate() {
	rg.generateBody()
	rg.generateFile()
}

func (rg *RepositoryGenerator) generateBody() {
	rg.Body = ExecuteTpl(RepositoryTpl(), map[string]interface{}{
		"Repository": rg.Repository,
		"Config":     rg.C,
		"Extra": map[string]interface{}{
			"Date": time.Now().Format(rg.C.Global.DateLayout),
		},
	})
	if rg.C.Verbose {
		repositoryGeneratorLogger.Println(fmt.Sprintf("[generateBody] for entity[%s]", rg.Repository.Entity.Name))
	}
}

func (rg *RepositoryGenerator) generateFile() {
	paths := make([]string, 0)
	paths = append(paths, rg.C.OutputDir)
	paths = append(paths, strings.Split(rg.Repository.PKGName, ".")...)
	fileName := filepath.Join(paths...) + ".java"
	dir := filepath.Dir(fileName)
	_ = os.MkdirAll(dir, 0700)
	_ = os.WriteFile(fileName, []byte(rg.Body), 0700)
	if rg.C.Verbose {
		repositoryGeneratorLogger.Println(fmt.Sprintf("[generateFile] for entity[%s], saved as [%s]", rg.Repository.Entity.Name, fileName))
	}
}
