package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"time"
)

var tlsCertFile = "certs/server.crt"
var tlsKeyFile = "certs/server.key"
var Router = mux.NewRouter().StrictSlash(true)
var Server = &http.Server{}

func main() {
	Server = &http.Server{
		Handler:      Router,
		Addr:         "127.0.0.1:8000",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	setAuthorizationHandlers()

	fmt.Println(Server.ListenAndServeTLS(tlsCertFile, tlsKeyFile))
}
