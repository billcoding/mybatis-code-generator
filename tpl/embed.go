package tpl

import "embed"

//go:embed entity.tpl mapper.tpl repository.tpl xml.tpl
var FS embed.FS
var entityTpl = `entity.tpl`
var mapperTpl = `mapper.tpl`
var repositoryTpl = `repository.tpl`
var xmlTpl = `xml.tpl`
var entityTplContent = ""
var mapperTplContent = ""
var repositoryTplContent = ""
var xmlTplContent = ""

func EntityTpl() string {
	if entityTplContent == "" {
		file, err := FS.ReadFile(entityTpl)
		if err != nil {
			panic(err)
		}
		entityTplContent = string(file)
	}
	return entityTplContent
}

func MapperTpl() string {
	if mapperTplContent == "" {
		file, err := FS.ReadFile(mapperTpl)
		if err != nil {
			panic(err)
		}
		mapperTplContent = string(file)
	}
	return mapperTplContent
}

func RepositoryTpl() string {
	if repositoryTplContent == "" {
		file, err := FS.ReadFile(repositoryTpl)
		if err != nil {
			panic(err)
		}
		repositoryTplContent = string(file)
	}
	return repositoryTplContent
}

func XMLTpl() string {
	if xmlTplContent == "" {
		file, err := FS.ReadFile(xmlTpl)
		if err != nil {
			panic(err)
		}
		xmlTplContent = string(file)
	}
	return xmlTplContent
}
