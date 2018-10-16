package router

import (
	"fmt"
	"go-postbar/model"
	"net/http"

	"github.com/gorilla/mux"
)

// var log = logger.Log

// Methods
const (
	GET    string = "GET"
	POST   string = "POST"
	DELETE string = "DELETE"
)

// type Person struct {
// 	name string
// 	age  int
// }

// PostsHandler PostsHandler
func PostsHandler(w http.ResponseWriter, r *http.Request) {
	p := model.Post{}

	fmt.Println(p)
}

// PostHandler PostHandler
func PostHandler(w http.ResponseWriter, r *http.Request) {}

// CommentsHandler CommentsHandler
func CommentsHandler(w http.ResponseWriter, r *http.Request) {}

// LikePostHandler LikePostHandler
func LikePostHandler(w http.ResponseWriter, r *http.Request) {}

// LikeCommentsHandler LikeCommentsHandler
func LikeCommentsHandler(w http.ResponseWriter, r *http.Request) {}

// PostsRouter PostsRouter
func PostsRouter() *mux.Router {

	router := mux.NewRouter()

	router.HandleFunc("/", PostsHandler).Methods(GET)
	router.HandleFunc("/{id}", PostHandler).Methods(GET)
	router.HandleFunc("/{id}/comments/{commentId}", CommentsHandler).Methods(GET, POST)
	router.HandleFunc("/{id}/like", LikePostHandler).Methods(POST)
	router.HandleFunc("/{id}/comments/{commentId}/like", LikeCommentsHandler).Methods(POST)

	return router
}
