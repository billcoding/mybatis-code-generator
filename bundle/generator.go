package bundle

import (
	. "github.com/billcoding/mybatis-code-generator/config"
	. "github.com/billcoding/mybatis-code-generator/generator"
	. "github.com/billcoding/mybatis-code-generator/model"
)

func GetEntityGenerators(CFG *Configuration, tableMap map[string]*Table) []Generator {
	egs := make([]Generator, 0)
	for _, v := range tableMap {
		eg := &EntityGenerator{
			C:     CFG,
			Table: v,
		}
		eg.Init()
		egs = append(egs, eg)
	}
	return egs
}

func GetMapperGenerators(CFG *Configuration, entityGenerators []Generator) []Generator {
	egs := make([]Generator, 0)
	for _, eg := range entityGenerators {
		mg := &MapperGenerator{
			C: CFG,
		}
		mg.Init(eg.(*EntityGenerator).Entity)
		egs = append(egs, mg)
	}
	return egs
}

func GetRepositoryGenerators(CFG *Configuration, entityGenerators []Generator) []Generator {
	rgs := make([]Generator, 0)
	for _, eg := range entityGenerators {
		rg := &RepositoryGenerator{
			C: CFG,
		}
		rg.Init(eg.(*EntityGenerator).Entity)
		rgs = append(rgs, rg)
	}
	return rgs
}

func GetXMLGenerators(CFG *Configuration, mapperGenerators []Generator) []Generator {
	xgs := make([]Generator, 0)
	for _, mg := range mapperGenerators {
		xg := &XMLGenerator{
			C: CFG,
		}
		xg.Init(mg.(*MapperGenerator).Mapper)
		xgs = append(xgs, xg)
	}
	return xgs
}
