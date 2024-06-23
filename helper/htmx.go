package helper

import (
	"fmt"
	"html/template"
	"net/http"
	"strings"
)

func CreateHtmxElement(urlPattern string, t string, cb func(*http.Request, http.ResponseWriter) (any, error)) func(http.ResponseWriter, *http.Request) (template.HTML, error) {
	pt, err := template.ParseFiles(fmt.Sprintf("./templates/htmx/%s", t))
	if err != nil {
		panic(fmt.Sprintf("The template file %s not found in templates/htmx", t))
	}
	http.HandleFunc(urlPattern, func(w http.ResponseWriter, r *http.Request) {
		data, err := cb(r, w)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		} else {
			pt.Execute(w, data)
		}
	})
	return func(w http.ResponseWriter, r *http.Request) (template.HTML, error) {
		data, err := cb(r, w)
		if err != nil {
			return "", err
		}
		builder := &strings.Builder{}
		if err := pt.Execute(builder, data); err != nil {
			return "", err
		}
		return template.HTML(builder.String()), nil
	}
}