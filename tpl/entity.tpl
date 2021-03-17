package {{.Config.Entity.PKG}};

{{range $i, $e := .Imports}}import {{$e}};
{{end}}
/**
 {{if .Config.Entity.Comment}}* {{.Entity.Table.Comment}}{{end}}
 * @author {{.Config.Global.Author}}
 {{if .Config.Global.Date}}* @since {{.Extra.Date}}{{end}}
 {{if .Config.Global.Copyright}}* @created by {{.Config.Global.CopyrightContent}}{{end}}
 {{if .Config.Global.Website}}* @repo {{.Config.Global.WebsiteContent}}{{end}}
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
    {{range $i, $e := .Entity.Fields}}
    {{if $e.Comment}}/**
     * {{$e.Column.Comment}}
     */{{end}}
    {{if $e.IdAnnotation}}@Id{{end}}
    {{if $e.ColumnAnnotation}}@Column(name = "{{$e.Column.Name}}"){{end}}
    private {{$e.Type}} {{$e.Name}};
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