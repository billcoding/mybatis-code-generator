package {{.Config.Mapper.PKG}};

{{if .Config.Mapper.MybatisPlus}}
import com.baomidou.mybatisplus.core.mapper.BaseMapper;
import {{.Config.Entity.PKG}}.{{.Mapper.Entity.Name}};
{{end}}

/**
 {{if .Config.Mapper.Comment}}* {{.Entity.Table.Name}} Mapper 接口{{end}}
 *
 * @author {{.Config.Global.Author}}
 {{if .Config.Global.Date}}* @since {{.Extra.DATE}}{{end}}
 {{if .Config.Entity.Copyright}}@create by {{.Config.Entity.CopyrightContent}}{{end}}
 */
public interface {{.Mapper.Name}}{{if .Config.Mapper.Extend}}extends{{.end}}{{range $i, $e := .Config.Mapper.Extends}}{{ if gt $index 0 }}, {{end}}{{$e}}{{end}} {

}
