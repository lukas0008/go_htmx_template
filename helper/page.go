package helper

import (
	"fmt"
	"html/template"
	"net/http"
)

func CreatePage(urlPattern string, t string, cb func(http.ResponseWriter,*http.Request,*template.Template)) {
	temp, err := template.ParseFiles(fmt.Sprintf("./templates/pages/%s", t))
	if err != nil {
		panic(fmt.Sprintf("Template %s not found in templates/pages", t))
	}
	http.HandleFunc(urlPattern, func(w http.ResponseWriter, r *http.Request) {
		cb(w, r, temp)
	})
}