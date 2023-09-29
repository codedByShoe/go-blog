package utils

import (
	"fmt"
	"net/http"
	"text/template"
)

func RenderTemplate(w http.ResponseWriter, tmpl string, data interface{}) {
	layout := "../web/templates/layout.html"
	path := fmt.Sprintf("../web/templates/%s", tmpl)
	t, err := template.ParseFiles(path, layout)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = t.ExecuteTemplate(w, "layout", data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
