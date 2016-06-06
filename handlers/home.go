package handlers

import (
	"github.com/ignatov/boot-test/libhttp"
	"html/template"
	"net/http"
	"github.com/GeertJohan/go.rice"
	"github.com/docker/engine-api/client"
	"github.com/docker/engine-api/types"
	"golang.org/x/net/context"
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

	cli, err := client.NewClient("http://shstack.labs.intellij.net:2375", "v1.22", nil, nil)
	if err != nil {
		libhttp.HandleErrorJson(w, err)
		return
	}

	options := types.ContainerListOptions{All: true}
	containers, err := cli.ContainerList(context.Background(), options)
	if err != nil {
		libhttp.HandleErrorJson(w, err)
		return
	}

	tmpl.Execute(w, containers)
}
