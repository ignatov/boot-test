package handlers

import (
	"github.com/ignatov/boot-test/libhttp"
	"html/template"
	"net/http"
)

func GetHome(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")

	tmpl, err := template.ParseFiles("templates/dashboard.html.tmpl", "templates/home.html.tmpl")
	if err != nil {
		libhttp.HandleErrorJson(w, err)
		return
	}

	type Container struct {
		Name  string
		Count int
	}

	sweaters := Container{
		"wool123",
		17,
	}

	cts := []Container{
		sweaters,
		sweaters,
		sweaters,
	}

	tmpl.Execute(w, cts)
}
