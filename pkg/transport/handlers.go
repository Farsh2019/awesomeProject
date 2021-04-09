package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func Router() *mux.Router {
	r := mux.NewRouter()
	s := r.PathPrefix("/api/v1").Subrouter()
	s.HandleFunc("/hello-world", helloWorld).Methods(http.MethodGet)
	return r
}

func helloWorld(w http.ResponseWriter, _*http.Request){
	fmt.Fprint(w,"Hello world!")
}
