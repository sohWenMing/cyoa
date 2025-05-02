package testTemplate

import "html/template"

// import "text/template"
// ...
// t, err := template.New("foo").Parse(`{{define "T"}}Hello, {{.}}!{{end}}`)
// err = t.ExecuteTemplate(out, "T", "<script>alert('you have been pwned')</script>")

func initTemplate(templateName string) (returnedTemplate *template.Template, err error) {
	newTemplate := template.New(templateName)
	t, err := newTemplate.Parse(`{{define "T"}} <b>This is a template to hold {{.}}!</b>{{end}}`)
	if err != nil {
		return nil, err
	}
	return t, nil
}
