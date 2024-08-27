package server

import (
	"github.com/Jhon-Henkel/full_cycle_hexagonal_arc/application"
	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
	"log"
	"net/http"
	"os"
	"time"
)

type WebServer struct {
	Service application.IProductService
}

func NewWebServer(service application.IProductService) *WebServer {
	return &WebServer{Service: service}
}

func (w WebServer) Serve() {
	router := mux.NewRouter()
	middleware := negroni.New(negroni.NewLogger())

	server := &http.Server{
		ReadHeaderTimeout: 10 * time.Second,
		WriteTimeout:      10 * time.Second,
		Addr:              ":8080",
		Handler:           http.DefaultServeMux,
		ErrorLog:          log.New(os.Stderr, "log: ", log.LstdFlags),
	}
	err := server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
