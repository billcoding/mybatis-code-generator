package {{.Config.Entity.PKG}};
/**
 {{if .Config.Entity.Comment}}* {{.Entity.Comment}}{{end}}
 * @author {{.Config.Global.Author}}
 {{if .Config.Global.Date}}* @since {{.Extra.DATE}}{{end}}
 {{if .Config.Global.Copyright}}@create by {{.Config.Global.CopyrightContent}}{{end}}
*/
{{if .Config.Entity.EntityAnnotation}}@Entity{{end}}
{{if .Config.Entity.TableAnnotation}}@Table(name = "{{.Table.Name}}"){{end}}
{{if .Config.Entity.Lombok}}
{{if .Config.Entity.LombokData}}@Data{{end}}
{{if .Config.Entity.LombokNoArgsConstructor}}@NoArgsConstructor{{end}}
{{if .Config.Entity.LombokAllArgsConstructor}}@AllArgsConstructor{{end}}
{{if .Config.Entity.LombokBuilder}}@Builder{{end}}
{{end}}
public class {{.Entity.Name}}{{if .Config.Entity.Extend}} extends {{end}}{{.Config.Entity.Extends}}{{if .Config.Entity.Implement}} implements {{.end}}{{range $i, $e := .Config.Entity.Implements}}{{ if gt $index 0 }}, {{end}}{{$e}}{{end}}{
    {{range $i, $e := .Config.Entity.EntityClassPrefixes}}
    {{$e}}
    {{end}}

    {{if .Entity.HaveId}}
    {{if .Entity.IdAnnotation}}@Id{{end}}
    {{if .Entity.ColumnAnnotation}}@Column(name = "{{.Entity.Id.Column}}"){{end}}
    {{if .Entity.IdAnnotation}}@GeneratedValue(strategy = GenerationType.AUTO){{end}}
    private {{.Entity.Id.JavaType}} {{.Entity.Id.Name}};
    {{end}}

    {{range $i, $e := .Entity.Fields}}
    {{if .Config.Entity.FieldComment}}/**
     * {{$e.Comment}}
     */{{end}}
    {{if .Entity.ColumnAnnotation}}@Column(name = "{{$e.Column}}"){{end}}
    private {{$e.JavaType}} {{$e.Name}};
    {{end}}
    {{end}}

    {{range $i, $e := .Config.Entity.EntityClassSuffixes}}
    {{$e}}
    {{end}}
}