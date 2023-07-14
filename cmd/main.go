package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/monstrasitix/swa/config"
)

func init() {
	config.LoadDev()
}

func main() {
	addr := ":" + config.GetPort()
	tmpl := *config.NewTemplate()

	router := mux.NewRouter()
	server := http.Server{
		Addr:    addr,
		Handler: router,
	}

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmpl.Render(w, "home.html", nil)
	}).Methods(http.MethodGet)

	router.HandleFunc("/contact", func(w http.ResponseWriter, r *http.Request) {
		tmpl.Render(w, "contact.html", nil)
	}).Methods(http.MethodGet)

	fmt.Println("Listening on: http://localhost" + addr)
	log.Fatal(server.ListenAndServe())
}
