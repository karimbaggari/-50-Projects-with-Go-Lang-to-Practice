package methods

import (
	"crud-api/structs"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)

func DeleteMovie(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    params := mux.Vars(r)
    movieList := structs.MovieList

    for index, item := range movieList {
        if item.ID == params["id"] {
            // Remove the movie at index from movieList
            movieList = append(movieList[:index], movieList[index+1:]...)
            break
        }
    }

    // Update the movieList in the structs package
    structs.MovieList = movieList

    // Encode and return the updated movieList
    json.NewEncoder(w).Encode(movieList)
}

