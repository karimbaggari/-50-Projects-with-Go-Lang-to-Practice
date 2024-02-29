package main

import (
	"fmt"
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"crud-api/methods"
)



func main() {
	r:= mux.NewRouter()
	r.HandleFunc("/movies",methods.GetMovies).Methods("GET")
	r.HandleFunc("/movies/{id}",methods.GetMovie).Methods("GET")
	r.HandleFunc("/movies", methods.CreateMovie).Methods("POST")
	r.HandleFunc("/movies/{id}",methods.UpdateMovie).Methods("PUT")
	r.HandleFunc("/movies/{id}",methods.DeleteMovie).Methods("DELETE")

	fmt.Println("starting server on port:8080")
	log.Fatal(http.ListenAndServe(":8080",r))
}




