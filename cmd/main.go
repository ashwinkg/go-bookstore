package main

import (
	"github.com/ashwinkg/go-bookstore/pkg/models"
	"github.com/ashwinkg/go-bookstore/pkg/routes"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	err := models.InitDB()
	if err != nil {
		log.Fatalf("Could not connect to the database: %v", err)
	}
	router := mux.NewRouter()
	routes.RegisterBookStoreRoutes(router)
	http.Handle("/", router)
	http.ListenAndServe("localhost:9010", router)
}
