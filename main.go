package main

import (
	"backend/src/apis"
	"fmt"
	"log"
	"net/http"
	"time"
)

var (
	enablePrintRecovery bool
	address             string
	timeout             int
)

func init() {
	enablePrintRecovery = true
	address = ":3001"
	timeout = 15
}

func main() {
	controllers := apis.NewControllers()
	handler := apis.NewRouter(controllers)

	server := &http.Server{
		Handler:      handler,
		Addr:         address,
		WriteTimeout: time.Duration(timeout) * time.Second,
		ReadTimeout:  time.Duration(timeout) * time.Second,
	}

	fmt.Println("Server start at: http://localhost" + address)
	log.Fatal(server.ListenAndServe())
}
