package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"./logger"
	"github.com/gorilla/mux"
)

var (
	port = flag.Int("port", 8181, "HTTP port")
	help = flag.Bool("help", false, "Help")
)

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
	fmt.Println("Endpoint Hit: homePage")
}

func CartHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "username: %v\n", vars["username"])
}

func handleRequests(port int) {
	router := mux.NewRouter()
	router.HandleFunc("/cart", CartHandler)
	addr := fmt.Sprintf("0.0.0.0:%d", port)
	logger.Info(fmt.Sprintf("Iniciando no endereço: %s", addr))

	log.Fatal(http.ListenAndServe(addr, nil))
}

func main() {
	flag.Parse()

	if *help {
		fmt.Fprintf(os.Stderr, "Usage of %s:\n", os.Args[0])
		flag.PrintDefaults()
		return
	}

	handleRequests(*port)
}
