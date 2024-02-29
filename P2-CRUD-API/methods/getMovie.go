package methods

import (
	"net/http"
    "encoding/json"
	"github.com/gorilla/mux"
	"crud-api/structs"
)


func GetMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	movieList := structs.MovieData()
	for _, item := range movieList {
		if item.ID == params["id"] {
            json.NewEncoder(w).Encode(item)
            return
        }
	}
	http.NotFound(w, r)
}