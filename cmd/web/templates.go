package main

import (
	"github.com/agung96tm/go-simple-web-application/internal/data"
	"html/template"
	"path/filepath"
)

type templateData struct {
	Users *[]data.User
	User  *data.User

	CurrentYear int
}

func newTemplateCache() (map[string]*template.Template, error) {
	cache := map[string]*template.Template{}

	pages, err := filepath.Glob("./internal/ui/html/pages/*.gohtml")
	if err != nil {
		return nil, err
	}

	for _, page := range pages {
		name := filepath.Base(page)
		files := []string{
			"./internal/ui/html/pages/base.gohtml",
			"./internal/ui/html/partials/nav.gohtml",
			page,
		}

		ts, err := template.ParseFiles(files...)
		if err != nil {
			return nil, err
		}

		cache[name] = ts
	}

	return cache, nil
}
