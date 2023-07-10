package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	env "github.com/monstrasitix/swa/config"
)

func main() {
	env.LoadDev()

	addr := ":" + env.GetPort()

	router := mux.NewRouter()
	server := http.Server{
		Addr:    addr,
		Handler: router,
	}

	router.PathPrefix("/").
		Handler(http.FileServer(http.Dir("./public/")))

	fmt.Println("Listening on: http://localhost" + addr)
	log.Fatal(server.ListenAndServe())
}
