package router

import (
	"fmt"
	"go-postbar/logger"
	"net/http"
)

var log = logger.Log

// StartServer (address string, port int)
func StartServer(address string, port int) {

	http.Handle("/posts", PostsRouter())
	// http.Handle("/h", HealthRouter())

	addr := fmt.Sprintf("%s:%d", address, port)

	log.Info(fmt.Sprintf("Service started, address: %s", addr))

	log.Fatal(http.ListenAndServe(addr, nil))
}
