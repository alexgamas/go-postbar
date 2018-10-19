/*
Autor: Alex Gamas (@alexgamas)

Based on ...

Soham Kamani article.
https://www.sohamkamani.com/blog/2017/09/13/how-to-build-a-web-application-in-golang/

Building a RESTful API with Golang
https://www.codementor.io/codehakase/building-a-restful-api-with-golang-a6yivzqdo

Effective Go
https://golang.org/doc/effective_go.html#methods

*/

package main

import (
	"flag"
	"fmt"
	"go-postbar/database"
	hnd "go-postbar/handlers"
	"go-postbar/logger"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

var (
	Log      = logger.Log
	address  = flag.String("address", "0.0.0.0", "Bind Address")
	port     = flag.Int("port", 8080, "HTTP port")
	help     = flag.Bool("help", false, "Help")
	mongoURL = flag.String("mongo_url", "mongodb://localhost/postbar", "MongoDB connection string")
)

// PostsRouter PostsRouter
func PostsRouter() *mux.Router {

	r := mux.NewRouter()

	r.HandleFunc("/health", hnd.HealthHandler).Methods(hnd.GET)

	r.HandleFunc("/posts", hnd.PostsHandler).Methods(hnd.GET, hnd.POST)
	r.HandleFunc("/posts/{id}", hnd.PostHandler).Methods(hnd.GET, hnd.DELETE)

	r.NotFoundHandler = http.HandlerFunc(hnd.NotFoundHandler)
	return r
}

// StartServer Inicia o servidor com os parametros passados
func StartServer(address string, port int) {

	database.Dial(*mongoURL)
	defer database.Close()

	r := PostsRouter()

	addr := fmt.Sprintf("%s:%d", address, port)

	Log.Info(fmt.Sprintf("Service started, address: %s", addr))

	//log.Fatal(http.ListenAndServe(addr, handlers.LoggingHandler(os.Stdout, r)))

	Log.Fatal(http.ListenAndServe(addr, handlers.LoggingHandler(logger.DebugLogWriter, r)))

}

func main() {
	flag.Parse()

	if *help {
		fmt.Fprintf(os.Stderr, "Usage of %s:\n", os.Args[0])
		flag.PrintDefaults()
		return
	}

	StartServer(*address, *port)
}
