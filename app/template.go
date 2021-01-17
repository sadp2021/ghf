package app

const buildTemplate = `{{range .Field}}//Set{{.Title}} 设置 {{.Name}}
func (b *{{$.Name}}) Set{{.Title}}({{.Name}} {{.TypeName}}) *{{$.Name}} {
	b.{{.Name}} = {{.Name}}

	return b
}
{{end}}`
