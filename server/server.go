package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"

	rest "github.com/ma-he-sh/postman-api-hack/server/rest"
)

func main() {
	r := mux.NewRouter().StrictSlash(true)

	mainRouter := r.Host("localhost").Subrouter()
	rest.RestRoutes(mainRouter)

	fs := http.FileServer(http.Dir("./payload"))
	r.PathPrefix("/payload/").Handler(http.StripPrefix("/payload/", fs))

	cors := handlers.CORS(
		handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type"}),
		handlers.AllowedOrigins([]string{"*"}),
	)

	r.Use(cors)

	server := &http.Server{
		Addr:         ":8080",
		Handler:      r,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
	}
	log.Fatal(server.ListenAndServe())
}
