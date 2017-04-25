package genangular

import "text/template"

type typeData struct {
	Name  string
	Desc  string
	Props []string
}

var typeT = `
{{ if .Desc }}// {{ .Desc }}{{ end }}
export class {{ .Name }} {
	{{ range .Props }}
	{{ . }}
	{{ end }}
}
`

var typeTmpl = template.Must(template.New("type").Parse(typeT))
