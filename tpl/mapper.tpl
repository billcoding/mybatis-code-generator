package {{.Config.Mapper.PKG}};

{{if .Config.Mapper.MybatisPlus}}import com.baomidou.mybatisplus.core.mapper.BaseMapper;
import {{.Mapper.Entity.PKGName}};
{{end}}{{if .Config.Mapper.MapperAnnotation}}import org.apache.ibatis.annotations.Mapper;{{end}}

/**
 {{if .Config.Mapper.Comment}}* {{.Mapper.Entity.Table.Comment}} Mapper 接口{{end}}
 * @author {{.Config.Global.Author}}
 {{if .Config.Global.Date}}* @since {{.Extra.Date}}{{end}}
 {{if .Config.Global.Copyright}}* @create by {{.Config.Global.CopyrightContent}}{{end}}
 {{if .Config.Global.Website}}* @see {{.Config.Global.WebsiteContent}}{{end}}
 */
{{if .Config.Mapper.MapperAnnotation}}@Mapper{{end}}
public interface {{.Mapper.Name}}{{if .Config.Mapper.MybatisPlus}} extends BaseMapper<{{.Mapper.Entity.Name}}>{{end}}{
}