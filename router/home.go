package router

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

// func homePage(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintf(w, "Welcome to the HomePage!")
// 	fmt.Println("Endpoint Hit: homePage")
// }

// HomeHandler HomeHandler
// func HomeHandler(w http.ResponseWriter, r *http.Request) {
// 	// 	// vars := mux.Vars(r)
// 	// 	// "username: %v\n", vars["username"]
// 	w.WriteHeader(http.StatusOK)
// 	// 	// log.Info(fmt.Sprintf("%v", vars))
// 	fmt.Fprintf(w, "ok")
// }
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	// 	// vars := mux.Vars(r)
	// 	// "username: %v\n", vars["username"]
	w.WriteHeader(http.StatusOK)
	// 	// log.Info(fmt.Sprintf("%v", vars))
	fmt.Fprintf(w, "ok")
}

// HomeRouter HomeRouter
func HomeRouter() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/", HomeHandler)
	router.HandleFunc("/home", HomeHandler)

	return router
}
