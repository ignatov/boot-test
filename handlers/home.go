package handlers

import (
	"github.com/ignatov/boot-test/libhttp"
	"html/template"
	"net/http"
	"github.com/GeertJohan/go.rice"
	docker_client "github.com/docker/engine-api/client"
	"github.com/docker/engine-api/types"
	"golang.org/x/net/context"
	registry_client "github.com/docker/distribution/registry/client"
	"io"
)

func GetRepositories(registry registry_client.Registry) ([]string, error) {
	entriesBuf := make([]string, 10000)
	numFilled, err := registry.Repositories(context.Background(), entriesBuf, "")
	if err != io.EOF {
		return nil, err
	}

	entries := make([]string, numFilled)
	copy(entries, entriesBuf)

	return entries, nil
}

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

	registry, err := registry_client.NewRegistry(context.Background(), "http://docker-registry.labs.intellij.net/", nil)
	if err != nil {
		libhttp.HandleErrorJson(w, err)
		return
	}

	entries, err := GetRepositories(registry)
	if err != nil {
		libhttp.HandleErrorJson(w, err)
		return
	}

	cli, err := docker_client.NewClient("http://shstack.labs.intellij.net:2375", "v1.22", nil, nil)
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

	type HomePageData struct {
		Containers []types.Container
		Images []string
	}

	tmpl.Execute(w, HomePageData{containers, entries})
}
