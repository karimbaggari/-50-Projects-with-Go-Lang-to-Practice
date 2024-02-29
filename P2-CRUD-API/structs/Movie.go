package structs

type Movie struct {
	ID       string    `json:"id"`
	Isbn     string    `json:"isbn"`
	Title    string    `json:"title"`
	Director *Director `json:"director"`
}

type Director struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Age       int    `json:"age"`
}

var MovieList = []Movie{
	{ID: "1", Isbn: "428423", Title: "Movie One", Director: &Director{FirstName: "Movie One", LastName: "Movie"}},
	{ID: "2", Isbn: "428424", Title: "Movie Two", Director: &Director{FirstName: "Movie Two", LastName: "Movie"}},
	{ID: "3", Isbn: "428425", Title: "Movie Three", Director: &Director{FirstName: "Movie Three", LastName: "Movie"}},
}

func MovieData() []Movie {

	return MovieList
}