package handlers

import (
	"encoding/json"
	"go-postbar/database"
	"go-postbar/logger"
	"go-postbar/model"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

// Collection Posts collection
const collection = "posts"

var log = logger.Log

func getAll(w http.ResponseWriter) {
	results := database.GetAll(collection)

	if results == nil {
		w.WriteHeader(http.StatusNotFound)
	} else {
		json.NewEncoder(w).Encode(results)
	}
}

func save(w http.ResponseWriter, r *http.Request) {
	var post model.Post

	_ = json.NewDecoder(r.Body).Decode(&post)

	if &post != nil {
		post.Created = time.Now()
		database.Save(collection, post)
		json.NewEncoder(w).Encode(post)
	}

}

// PostsHandler PostsHandler
func PostsHandler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	if r.Method == GET {
		getAll(w)
	} else if r.Method == POST {
		save(w, r)
	}

}

// PostHandler PostHandler
func PostHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	var id = params["id"]

	if r.Method == GET {

		result := database.GetOne("posts", id)

		if result == nil {
			w.WriteHeader(http.StatusNotFound)
		} else {
			json.NewEncoder(w).Encode(result)
		}
	} else if r.Method == DELETE {

		result := database.DeleteOne("posts", id)

		if !result {
			w.WriteHeader(http.StatusNotFound)
		}
	}
}

// DeleteHandler PostHandler
func DeleteHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var id = params["id"]
	result := database.GetOne("posts", id)

	if result == nil {
		w.WriteHeader(http.StatusNotFound)
	} else {
		json.NewEncoder(w).Encode(result)
	}
}
