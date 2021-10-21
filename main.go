package main

import (
	"log"
	"net/http"

	"github.com/matisidler/stock-checker/pkg/router"
)

//url almacena la url donde se realiza la petición
var url = "https://yh-finance.p.rapidapi.com/stock/v2/get-summary?symbol=TSLA&region=US"

const (
	port = ":8080"
)

func main() {

	srv := &http.Server{
		Addr:    port,
		Handler: router.Routes(),
	}

	log.Println("initializing server on port", port)

	err := srv.ListenAndServe()
	if err != nil {
		log.Fatal("error initializing server:", err)
	}

	/* req, _ := http.NewRequest("GET", url, nil)

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
