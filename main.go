package main

import (
	_ "golang-todolist/database"
	"golang-todolist/routes"
	"log"
	"net/http"
)

func main() {

	router := routes.NewRouter()

	server := http.Server{
		Addr:    "localhost:3000",
		Handler: router,
	}

	err := server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}

}
