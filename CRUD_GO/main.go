package main

import (
	"log"
	"net/http"

	routes "../CRUD_GO/pkgs/routes"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterBookStoreRoutes(r)
	http.Handle("/", r)
	log.Println("Server listening at port 8080")
	log.Fatal(http.ListenAndServe("localhost:8080", r))
}
