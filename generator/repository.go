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

var repositoryGeneratorLogger = log.New(os.Stdout, "[RepositoryGenerator]", log.LstdFlags)

type RepositoryGenerator struct {
	C                 *Configuration
	Repository        *Repository
	RepositoryContent string
}

func (rg *RepositoryGenerator) Init(e *Entity) {
	rg.Repository = &Repository{
		PKG:    rg.C.Mapper.PKG,
		Entity: e,
	}
	rg.Repository.Name = rg.C.Repository.RepositoryNamePrefix + rg.Repository.Entity.Name + rg.C.Repository.RepositoryNameSuffix
	rg.Repository.PKGName = rg.C.Repository.PKG + "." + rg.Repository.Name
}

func (rg *RepositoryGenerator) Generate() {
	rg.generateInterface()
	rg.generateFile()
}

func (rg *RepositoryGenerator) generateInterface() {
	class := ExecuteTpl(RepositoryTpl(), map[string]interface{}{
		"Repository": rg.Repository,
		"Config":     rg.C,
		"Extra": map[string]interface{}{
			"Date": time.Now().Format(rg.C.Global.DateLayout),
		},
	})
	var buffer strings.Builder
	_, _ = io.WriteString(&buffer, class)
	rg.RepositoryContent = buffer.String()
	if rg.C.Verbose {
		repositoryGeneratorLogger.Println(fmt.Sprintf("[generateInterface] for entity[%s]", rg.Repository.Entity.Name))
	}
}

func (rg *RepositoryGenerator) generateFile() {
	paths := make([]string, 0)
	paths = append(paths, rg.C.OutputDir)
	paths = append(paths, strings.Split(rg.Repository.PKGName, ".")...)
	repositoryFileName := filepath.Join(paths...) + ".java"
	dir := filepath.Dir(repositoryFileName)
	_ = os.MkdirAll(dir, 0700)
	_ = os.WriteFile(repositoryFileName, []byte(rg.RepositoryContent), 0700)
	if rg.C.Verbose {
		repositoryGeneratorLogger.Println(fmt.Sprintf("[generateFile] for entity[%s], saved as [%s]", rg.Repository.Entity.Name, repositoryFileName))
	}
}
