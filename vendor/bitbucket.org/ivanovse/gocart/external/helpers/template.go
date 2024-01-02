package helpers

import (
	"html/template"
	"strings"
)

func NewTemplate(name string) *template.Template {
	return template.New("hello.gohtml").Funcs(template.FuncMap{
		"inc": func(n int) int {
			return n + 1
		},
		"tohtml": func(feature string) template.HTML {
			return template.HTML(feature)
		},
		"tojs": func(feature string) template.JS {
			return template.JS(feature)
		},
		"joinmapstrings": func(feature map[string]int) string {
			push := []string{}
			for k, _ := range feature {
				push = append(push, k)
			}
			return strings.Join(push, ",")
		},
	})
}

