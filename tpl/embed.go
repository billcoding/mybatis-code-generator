package tpl

import "embed"

//go:embed entity.tpl mapper.tpl xml.tpl
var FS embed.FS
