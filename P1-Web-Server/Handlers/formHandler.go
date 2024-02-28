package Handlers

import (
	"fmt"
	"net/http"
)

func FormHandler(w http.ResponseWriter, r *http.Request) {

	if r.URL.Path != "/form" {
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}

	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm error: %v", err)
	}

	name := r.FormValue("name")
	address := r.FormValue("address")
	fmt.Fprintf(w, "Name = %s\n", name)
	fmt.Fprintf(w, "Address = %s\n", address)

	fmt.Fprintf(w, "Hello Form")
}
