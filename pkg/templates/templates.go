package templates

import (
	"fmt"
	"html/template"
	"net/http"
)

type Service interface {
	Render(w http.ResponseWriter, tmpl string, layout string, data interface{}) error
}

type TemplateService struct {
	htmlTpl *template.Template
}

func NewTemplateService() (*TemplateService, error) {
	templates, err := template.ParseGlob("pkg/templates/*.html")
	if err != nil {
		return nil, err
	}
	return &TemplateService{htmlTpl: templates}, nil
}

func (ts *TemplateService) Render(w http.ResponseWriter, tmpl string, layout string, data interface{}) error {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	t, err := template.ParseFiles(fmt.Sprintf("pkg/templates/%s", layout), fmt.Sprintf("pkg/templates/%s", tmpl))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return err
	}
	err = t.ExecuteTemplate(w, "layout", data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	return err
}
