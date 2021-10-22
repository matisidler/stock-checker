package render

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)

var templateCache map[string]*template.Template

func ExecutingTemplate(w http.ResponseWriter, r *http.Request, tmpl string) {
	tpl, ok := templateCache[tmpl]
	if !ok {
		log.Fatal("template not found")
	}

	buf := new(bytes.Buffer)

	tpl.Execute(buf, nil)

	_, err := buf.WriteTo(w)
	if err != nil {
		log.Fatal("error writting template to browser")
	}

}

func CreatingTemplateCache() (map[string]*template.Template, error) {

	templateCache = map[string]*template.Template{}

	pages, err := filepath.Glob("./template/*.page.tpl")

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

		layouts, err := filepath.Glob("./template/*.layout.tpl")
		if err != nil {
			log.Println("error CreatingTemplateCache 2: ")
			return templateCache, err
		}

		if len(layouts) > 0 {
			fmt.Println("aca")
			tpl, err = tpl.ParseGlob("./template/*.layout.tpl")
			if err != nil {
				log.Println("error CreatingTemplateCache 3: ")
				return templateCache, err
			}
		}
		templateCache[namePage] = tpl
	}
	fmt.Println(templateCache)
	return templateCache, nil

}
