package server

import (
	"github.com/alyssonvitor500/go-hexagonal/adapters/web/handler"
	"github.com/alyssonvitor500/go-hexagonal/application"
	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
	"time"
)

type Webserver struct {
	Service application.ProductServiceInterface
}

func MakeNewWebserver(service application.ProductServiceInterface) *Webserver {
	return &Webserver{Service: service}
}

func (webserver *Webserver) Serve() {

	r := mux.NewRouter() // Responsible for routes
	n := negroni.New( // Middleware
		negroni.NewLogger(),
	)

	handler.MakeProductHandler(r, n, webserver.Service)
	http.Handle("/", r)

	server := &http.Server{
		ReadHeaderTimeout: 10 * time.Second,
		WriteTimeout:      10 * time.Second,
		Addr:              ":9000",
		Handler:           http.DefaultServeMux,
		ErrorLog:          log.New(os.Stderr, "log: ", log.Lshortfile),
	}

	err := server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
