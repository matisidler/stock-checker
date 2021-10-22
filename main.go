package main

import (
	"log"
	"net/http"

	"github.com/matisidler/stock-checker/pkg/config"
	"github.com/matisidler/stock-checker/pkg/handlers"
	"github.com/matisidler/stock-checker/pkg/render"
	"github.com/matisidler/stock-checker/pkg/router"
)

//url almacena la url donde se realiza la petición
var (
	configuration = config.AppConfig{}
	port          = ":8080"
)

func main() {
	run()

	srv := &http.Server{
		Addr:    port,
		Handler: router.Routes(&configuration),
	}

	log.Println("initializing server on port", port)

	err := srv.ListenAndServe()
	if err != nil {
		log.Fatal("error initializing server:", err)
	}

	/*
		url         = "https://yh-finance.p.rapidapi.com/stock/v2/get-summary?symbol=TSLA&region=US"

		req, _ := http.NewRequest("GET", url, nil)

		req.Header.Add("x-rapidapi-host", "yh-finance.p.rapidapi.com")
		req.Header.Add("x-rapidapi-key", "b71d688d90msheeb554bbd71d876p11f0cbjsna25c43d11625")

		res, err := http.DefaultClient.Do(req)
		if err != nil {
			log.Println("error obteniendo la respuesta")
		}

		defer res.Body.Close()
		var m structs.Model
		err = json.NewDecoder(res.Body).Decode(&m)
		if err != nil {
			log.Fatal("error decoding body", err)
		}

		fmt.Println(m.Price["quoteType"]) */

}

func run() {
	cache, err := render.CreatingTemplateCache()
	if err != nil {
		log.Fatal(err)
	}
	configuration = config.AppConfig{
		TplCache: cache,
	}

	handlers.Repo = handlers.NewRepo(configuration)

}
