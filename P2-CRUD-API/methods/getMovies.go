package methods

import (
	"net/http"
    "encoding/json"
	"crud-api/structs"
	
)


func GetMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	movieList := structs.MovieData()
	json.NewEncoder(w).Encode(movieList)
}