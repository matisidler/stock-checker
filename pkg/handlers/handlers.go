package handlers

import (
	"encoding/json"
	"log"
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request) {
	prueba, err := json.Marshal("Hola")
	if err != nil {
		log.Println("error marshalling prueba", err)
	}
	w.Write(prueba)
}
