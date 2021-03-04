package {{.Config.Mapper.PKG}};

{{if .Config.Mapper.MybatisPlus}}
import com.baomidou.mybatisplus.core.mapper.BaseMapper;
import {{.Mapper.Entity.PKGName}};
{{end}}

/**
 {{if .Config.Mapper.Comment}}* {{.Mapper.Entity.Table.Comment}} Mapper 接口{{end}}
 *
 * @author {{.Config.Global.Author}}
 {{if .Config.Global.Date}}* @since {{.Extra.DATE}}{{end}}
 {{if .Config.Global.Copyright}}@create by {{.Config.Global.CopyrightContent}}{{end}}
 */
public interface {{.Mapper.Name}}{{if .Config.Mapper.Extend}}extends{{.end}}{{range $i, $e := .Config.Mapper.Extends}}{{if gt $index 0}}, {{end}}{{$e}}{{end}} {

}
