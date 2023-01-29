package main

import (
	"github.com/gorilla/mux"
	"net/http"
)

var Router = mux.NewRouter().StrictSlash(true)
var Server = &http.Server{}

func main() {

}