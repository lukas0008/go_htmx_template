package helper

import (
	"fmt"
	"html/template"
	"net/http"
	"strings"
)

type LayoutData struct {
	Page template.HTML
}

type Layout struct {
	temps []*template.Template
}

func (l1 Layout) Nest(l2 Layout) Layout {
	temps := l1.temps[:]
	temps = append(temps, l2.temps[:]...)
	return Layout {
		temps: temps,
	}
}


func (l Layout) Execute(w http.ResponseWriter, t *template.Template, data any) error {
	sBuilder := &strings.Builder{};
	t.Execute(sBuilder, data)
	page := sBuilder.String()
	for i := len(l.temps)-1; i >= 0; i-- {
		temp := l.temps[i]
		sBuilder := &strings.Builder{};
		err := temp.Execute(sBuilder, LayoutData {
			Page: template.HTML(page),
		})
		if err != nil {
			return err
		}
		page = sBuilder.String()
	}
	fmt.Println(page)
	w.Write([]byte(page))
	return nil
}

func CreateLayout(t string) Layout {
	temp, err := template.ParseFiles(fmt.Sprintf("./templates/layouts/%s", t))
	if err != nil {
		panic(fmt.Sprintf("The template file %s not found in templates/htmx", t))
	}
	return Layout {
		temps: []*template.Template{temp},
	}
}