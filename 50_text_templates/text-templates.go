package main

import (
	"os"
	"text/template"
)

func main() {
	t1 := template.New("t1")
	t1, err := t1.Parse("value is {{.}}\n")
	if err != nil {
		panic(err)
	}
	t1 = template.Must(t1.Parse("value: {{.}}\n"))

	t1.Execute(os.Stdout, "some text")
	t1.Execute(os.Stdout, 5)
	t1.Execute(os.Stdout, []string{"Go", "Rust", "C++", "C#"})

	Create := func(name, t string) *template.Template {
		return template.Must(template.New(name).Parse(t))
	}
	t2 := Create("t2", "name: {{.name}}")

	t2.Execute(os.Stdout, struct {
		name string
	}{"jane doe"})
	t2.Execute(os.Stdout, map[string]string{
		"name": "mickey mouse",
	})
	t3 := Create("t3", "{{if . -}} yes {{else -}} no {{end}}\n")
	t3.Execute(os.Stdout, "not empty")
	t3.Execute(os.Stdout, "")

	t4 := Create("t4", "range: {{range .}}{{.}} {{end}}\n")
	t4.Execute(os.Stdout, []string{"go", "rust", "cpp", "csharp"})
}
