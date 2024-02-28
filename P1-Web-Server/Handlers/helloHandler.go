package Handlers

import (
	"fmt"
	"net/http"
)


func HelloHandler(w http.ResponseWriter,r *http.Request)  {
	if r.URL.Path != "/hello" {
		http.Error(w,"404 not found", http.StatusNotFound)
		return
	}
	if r.Method != "GET" {
		http.Error(w,"GET not supported", http.StatusNotFound)
		return
	}
	fmt.Fprintf(w,"Hello World")
}