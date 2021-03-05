package {{.Config.Entity.PKG}};

{{range $i, $e := .Imports}}import {{$e}};
{{end}}
/**
 {{if .Config.Entity.Comment}}* {{.Entity.Table.Comment}}{{end}}
 * @author {{.Config.Global.Author}}
 {{if .Config.Global.Date}}* @since {{.Extra.Date}}{{end}}
 {{if .Config.Global.Copyright}}* @create by {{.Config.Global.CopyrightContent}}{{end}}
 {{if .Config.Global.Website}}* @see {{.Config.Global.WebsiteContent}}{{end}}
*/
{{if .Config.Entity.EntityAnnotation}}@Entity{{end}}
{{if .Config.Entity.TableAnnotation}}@Table(name = "{{.Entity.Table.Name}}"){{end}}
{{if .Config.Entity.Lombok}}{{if .Config.Entity.LombokData}}@Data{{end}}
{{if .Config.Entity.LombokNoArgsConstructor}}@NoArgsConstructor{{end}}
{{if .Config.Entity.LombokAllArgsConstructor}}@AllArgsConstructor{{end}}
{{if .Config.Entity.LombokBuilder}}@Builder{{end}}{{end}}
public class {{.Entity.Name}}{{if .Config.Entity.Extend}} extends {{.Config.Entity.Extends}}{{end}}{{if .Config.Entity.Implement}} implements {{range $i, $e := .Config.Entity.Implements}}{{if gt $i 0}}, {{end}}{{$e}}{{end}}{{end}}{
    {{range $i, $e := .Config.Entity.EntityClassPrefixes}}{{$e}}
    {{end}}
    {{if .Entity.HaveId}}
    {{if .Config.Entity.FieldComment}}/**
     * {{.Entity.Id.Column.Comment}}
     */{{end}}
    {{if .Config.Entity.IdAnnotation}}@Id{{end}}
    {{if .Config.Entity.ColumnAnnotation}}@Column(name = "{{.Entity.Id.Column.Name}}"){{end}}
    {{if .Config.Entity.IdAnnotation}}@GeneratedValue(strategy = GenerationType.AUTO){{end}}
    private {{.Entity.Id.Type}} {{.Entity.Id.Name}};
    {{end}}
    {{range $i, $e := .Entity.Fields}}
    {{if $e.Comment}}/**
     * {{$e.Column.Comment}}
     */{{end}}
    {{if $e.ColumnAnnotation}}@Column(name = "{{$e.Column.Name}}"){{end}}
    private {{$e.Type}} {{$e.Name}};
    {{end}}
    {{range $i, $e := .Config.Entity.EntityClassSuffixes}}{{$e}}
    {{end}}
}