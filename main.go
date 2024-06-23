package main

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/lukas0008/go_htmx_template/helper"
	"github.com/lukas0008/go_htmx_template/pages"
)

type Test struct {
	Test template.HTML
}

func main() {
	mainLayout := helper.CreateLayout("main.html")
	pages.RegisterIndex(mainLayout)
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("GET /static/", http.StripPrefix("/static/", fs))
	fmt.Println("Starting")
	http.ListenAndServe("127.0.0.1:80", nil)
}