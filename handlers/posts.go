package handlers

import (
	"encoding/json"
	"go-postbar/database"
	"go-postbar/logger"
	"net/http"

	"github.com/gorilla/mux"
)

var log = logger.Log

// PostsHandler PostsHandler
func PostsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	results := database.GetAll("posts")
	json.NewEncoder(w).Encode(results)
}

// PostHandler PostHandler
func PostHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var id = params["id"]
	result := database.GetOne("posts", id)
	json.NewEncoder(w).Encode(result)
}

// CommentsHandler CommentsHandler
func CommentsHandler(w http.ResponseWriter, r *http.Request) {}

// LikePostHandler LikePostHandler
func LikePostHandler(w http.ResponseWriter, r *http.Request) {}

// LikeCommentsHandler LikeCommentsHandler
func LikeCommentsHandler(w http.ResponseWriter, r *http.Request) {}
