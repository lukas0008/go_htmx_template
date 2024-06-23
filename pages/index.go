package pages

import (
	"html/template"
	"net/http"

	"github.com/lukas0008/go_htmx_template/helper"
)

type indexData struct {
	MyHtmxElem template.HTML
}

func RegisterIndex(layout helper.Layout) {
	myelem := registerMyElem()
	helper.CreatePage("GET /", "index.html", func(w http.ResponseWriter, r *http.Request, t *template.Template) {
		elemHtml, err := myelem(w,r)
		if err != nil {
			panic("Will an error really happen here?")
		}
		layout.Execute(w, t, indexData {
			MyHtmxElem: elemHtml,
		})
	})
}