package main

import (
	"html/template"
	"net/http"
)

func setMainHandlers() {
	Router.HandleFunc("/", MainPage)
}

func MainPage(writer http.ResponseWriter, request *http.Request) {
	t, _ := template.ParseFiles(
		"templates/base.html", "templates/navbar.html", "templates/public.navbar.html", "templates/index.html",
	)
	t.ExecuteTemplate(writer, "base", nil)
}
