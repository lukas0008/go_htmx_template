# My personal Go+Htmx template

This is a very simple template that I made for myself to make working with htmx easy from the start.

## Running
Clone the repo
Run:
```go
go run .
```
It will run on port 80

## Requirements
- None afaik.
- Was made on go version 1.22.4

## Usage

### Static
Static files go to the /static folder, then are served under the GET /static/pathtoresource endpoint.

### Templates
All templates are meant to be stored under /templates

- htmx under /templates/htmx
- pages under /templates/pages
- layouts under /templates/layouts

Note that whenever you create a layout htmx element or page, you will use the path relative to their correspoding directory f.e. `mylayout.html` instead of `./templates/layouts/mylayout.html`.

### Layouts
Layouts can be created using `helper.CreateLayout(templatename)`.

You can nest layouts using `nestedLayout := myLayout.Nest(anotherLayout)`

Then to use a layout, run `layout.Execute(responseWriter, myPageTemplate, myPageData)`, instead of `myPageTemplate.Execute(responseWriter, myPageData`.

### Pages

Pages do not have much behaviour, the usage is exactly the same as `http.HandleFunc`, except having an extra, second param which asks for the template path. In the callback function you also get a `*template.Template`.

Its just convenience to not have to manually call `template.ParseFiles`

### Htmx

To create an Htmx element, you need to call `helper.CreateHtmxElement(urlPattern, htmxTemplatePath, callback)`.
> PS: Check the example usage of callback in /pages/myelem.htmx.go (its easier to understand that way)

This htmx is then hosted on the provided urlPattern, but there is also a callback returned from this. The callback allows you to get the html string from the element by calling it, this allows you to have an htmx element that is both rendered on request, and can be updated with htmx.

## Apologies
To anyone who actually wants to use this, I apologise. This is the first time I have used go in my life, so it is very likely I haven't followed good conventions.