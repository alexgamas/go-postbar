package handlers

import (
	"encoding/json"
	"go-postbar/database"
	"go-postbar/logger"
	"go-postbar/model"
	"net/http"
)

var log = logger.Log

// PostsHandler PostsHandler
func PostsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	database.GetAll("posts")
	json.NewEncoder(w).Encode("")

}

// PostHandler PostHandler
func PostHandler(w http.ResponseWriter, r *http.Request) {
	// w.WriteHeader(http.StatusOK)
	// log.Info(fmt.Sprintf("%v", vars))
	// fmt.Fprintf(w, "ok")
	w.Header().Set("Content-Type", "application/json")
	status := model.Health{Status: "UP"}
	json.NewEncoder(w).Encode(status)
}

// CommentsHandler CommentsHandler
func CommentsHandler(w http.ResponseWriter, r *http.Request) {}

// LikePostHandler LikePostHandler
func LikePostHandler(w http.ResponseWriter, r *http.Request) {}

// LikeCommentsHandler LikeCommentsHandler
func LikeCommentsHandler(w http.ResponseWriter, r *http.Request) {}
