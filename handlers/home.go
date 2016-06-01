package handlers

import (
	"github.com/ignatov/boot-test/libhttp"
	"html/template"
	"net/http"
	"github.com/GeertJohan/go.rice"
)

func GetHome(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")

	templateBox, err := rice.FindBox("../templates")
	dash, err := templateBox.String("dashboard.html.tmpl")
	home, err := templateBox.String("home.html.tmpl")
	tmpl, err := template.New("message").Parse(dash)
	tmpl, err = tmpl.Parse(home)

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
