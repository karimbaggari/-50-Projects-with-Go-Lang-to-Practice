package methods

import (
	"net/http"
    "encoding/json"
	"crud-api/structs"

	"math/rand"
	"strconv"
)

func CreateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type","application/json")
	var movie structs.Movie
	movieList := structs.MovieData()
	_ = json.NewDecoder(r.Body).Decode(&movie)
	movie.ID = strconv.Itoa(rand.Intn(10000))
	movieList = append(movieList, movie)
	json.NewEncoder(w).Encode(movieList)
}