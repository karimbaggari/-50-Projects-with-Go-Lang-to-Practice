package main

import (
	"fmt"
	"log"
	"net/http"
	"Go-P1/Handlers"
)

func main() {
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/",fileServer)
	http.HandleFunc("/form",Handlers.FormHandler)
	http.HandleFunc("/hello",Handlers.HelloHandler)
	fmt.Println("starting server at port : 8080")
	if err := http.ListenAndServe(":8080", nil);err != nil {
		log.Fatal(err)
	}
}