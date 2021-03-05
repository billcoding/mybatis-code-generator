package {{.Config.Repository.PKG}};

import {{.Repository.Entity.PKGName}};
import org.springframework.data.jpa.repository.JpaRepository;
{{if .Config.Repository.RepositoryAnnotation}}import org.springframework.stereotype.Repository;{{end}}

/**
 {{if .Config.Repository.Comment}}* {{.Repository.Entity.Table.Comment}} Repository 接口{{end}}
 * @author {{.Config.Global.Author}}
 {{if .Config.Global.Date}}* @since {{.Extra.Date}}{{end}}
 {{if .Config.Global.Copyright}}* @create by {{.Config.Global.CopyrightContent}}{{end}}
 {{if .Config.Global.Website}}* @see {{.Config.Global.WebsiteContent}}{{end}}
 */
{{if .Config.Repository.RepositoryAnnotation}}@Repository{{end}}
public interface {{.Repository.Name}} extends JpaRepository<{{.Repository.Entity.Name}}, {{.Repository.Entity.Id.Type}}>{
}
