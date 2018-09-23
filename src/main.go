package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"

	"./logger"
	"github.com/gorilla/mux"
)

var log = logger.Log

var (
	port = flag.Int("port", 8080, "HTTP port")
	help = flag.Bool("help", false, "Help")
)

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
	fmt.Println("Endpoint Hit: homePage")
}

//HomeHandler HomeHandler
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	// vars := mux.Vars(r)
	// "username: %v\n", vars["username"]
	w.WriteHeader(http.StatusOK)
	// log.Info(fmt.Sprintf("%v", vars))
	// fmt.Fprintf(w, ": %v\n", vars["username"])
}

// PostsHandler PostsHandler
func PostsHandler(w http.ResponseWriter, r *http.Request) {}

// PostHandler PostHandler
func PostHandler(w http.ResponseWriter, r *http.Request) {}

// CommentsHandler CommentsHandler
func CommentsHandler(w http.ResponseWriter, r *http.Request) {}

// LikePostHandler LikePostHandler
func LikePostHandler(w http.ResponseWriter, r *http.Request) {}

//LikeCommentsHandler LikeCommentsHandler
func LikeCommentsHandler(w http.ResponseWriter, r *http.Request) {}

// Methods
const (
	GET    string = "GET"
	POST   string = "POST"
	DELETE string = "DELETE"
)

func initServer(port int) {
	router := mux.NewRouter()

	router.HandleFunc("/", HomeHandler)
	router.HandleFunc("/posts", PostsHandler).Methods(GET)
	router.HandleFunc("/posts/{id}", PostHandler).Methods(GET)
	router.HandleFunc("/posts/{id}/comments/{commentId}", CommentsHandler).Methods(GET, POST)
	router.HandleFunc("/posts/{id}/like", LikePostHandler).Methods(POST)
	router.HandleFunc("/posts/{id}/comments/{commentId}/like", LikeCommentsHandler).Methods(POST)

	// router.HandleFunc("/posts/{id}/comments/{commentId}").Methods(GET)

	http.Handle("/", router)

	addr := fmt.Sprintf("0.0.0.0:%d", port)
	log.Info(fmt.Sprintf("Iniciando no endereço: %s", addr))
	log.Fatal(http.ListenAndServe(addr, nil))
}

func main() {
	flag.Parse()

	if *help {
		fmt.Fprintf(os.Stderr, "Usage of %s:\n", os.Args[0])
		flag.PrintDefaults()
		return
	}

	initServer(*port)
}
