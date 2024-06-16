package main

import (
	"github.com/ilayzen/swagger-golang/pkg/delivery/rest"
	"log"
	"net/http"
)

func main() {

	h := rest.NewHandlers()
	router := rest.NewRouter(h)

	err := http.ListenAndServe(":5020", router)
	if err != nil {
		log.Fatalf("Cannot run the server, err: %v", err)
	}
}
