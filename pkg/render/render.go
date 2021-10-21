package render

import (
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)

func ExecutingTemplate(w http.ResponseWriter, r *http.Request, tmpl string) {

}

func CreatingTemplateCache() (map[string]*template.Template, error) {

	templateCache := map[string]*template.Template{}

	pages, err := filepath.Glob("../../templates/*.page.tpl")
	if err != nil {
		log.Println("error CreatingTemplateCache 0: ")
		return templateCache, err
	}

	for _, page := range pages {
		namePage := filepath.Base(page)

		tpl, err := template.New(namePage).ParseFiles(page)
		if err != nil {
			log.Println("error CreatingTemplateCache 1: ")
			return templateCache, err
		}

		layouts, err := filepath.Glob("../../templates/*.layout.tpl")
		if err != nil {
			log.Println("error CreatingTemplateCache 2: ")
			return templateCache, err
		}

		if len(layouts) > 0 {
			tpl, err = tpl.ParseGlob("../../templates/*.layout.tpl")
			if err != nil {
				log.Println("error CreatingTemplateCache 3: ")
				return templateCache, err
			}
		}
		templateCache[namePage] = tpl
	}

	return templateCache, nil

}
