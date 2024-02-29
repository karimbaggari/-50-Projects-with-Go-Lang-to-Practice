package methods


import (
    "net/http"
    "encoding/json"
    "crud-api/structs"
	"github.com/gorilla/mux"
)

func UpdateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	movieList := structs.MovieData()
	for index, item := range movieList {
		if item.ID == params["id"] {
			movieList = append(movieList[:index], movieList[index+1:]...)
			var movie structs.Movie
			_ = json.NewDecoder(r.Body).Decode(&movie)
			movie.ID = params["id"]
			movieList = append(movieList, movie)
			json.NewEncoder(w).Encode(movieList)
			return
		}
	}
}
