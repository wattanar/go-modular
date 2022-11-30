package main

import (
	"go-modular/internal/app/router"
	"log"
	"net/http"
)

func main() {
	r := router.NewRouter()

	log.Fatal(http.ListenAndServe(":8080", r))
}
