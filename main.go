package main

import (
	"flag"
	"fmt"
	"go-postbar/logger"
	"go-postbar/router"
	"os"
)

var log = logger.Log

var (
	address = flag.String("address", "0.0.0.0", "Bind Address")
	port    = flag.Int("port", 8080, "HTTP port")
	help    = flag.Bool("help", false, "Help")
)

func main() {
	flag.Parse()

	if *help {
		fmt.Fprintf(os.Stderr, "Usage of %s:\n", os.Args[0])
		flag.PrintDefaults()
		return
	}

	router.StartServer(*address, *port)
}
