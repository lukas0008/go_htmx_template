package pages

import (
	"html/template"
	"net/http"
	"time"

	"github.com/lukas0008/go_htmx_template/helper"
)

type myelemData struct {
	Now string
}

func registerMyElem() func(http.ResponseWriter, *http.Request) (template.HTML, error) {
	return helper.CreateHtmxElement("GET /myelem.htmx", "myelem.html", func(r *http.Request, w http.ResponseWriter) (any, error) {
			return myelemData {
				Now: time.Now().String(),
			}, nil
	})
}