package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Movie struct {
	ID       string    `json::"id`
	Isbn     string    `json:"isbn"`
	Title    string    `json:"title`
	Director *Director `json:"director`
}

type Director struct {
	firstName string `json:"firstname"`
	lastName  string `json:"lastname"`
}

var movies []Movie

func getMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(movies)

}

func main() {
	r := mux.NewRouter()

	movies = append(movies, Movie{ID: "1", Isbn: " 123", Title: "Movie one", Director: &Director{firstName: "Excel", lastName: "Baraka"}})
	movies = append(movies, Movie{ID: "2", Isbn: " 124", Title: "Movie Two", Director: &Director{firstName: "Fav", lastName: "Neema"}})
	movies = append(movies, Movie{ID: "3", Isbn: " 125", Title: "Movie Three", Director: &Director{firstName: "Faith", lastName: "Natasha"}})
	movies = append(movies, Movie{ID: "4", Isbn: " 126", Title: "Movie Four", Director: &Director{firstName: "Ian", lastName: "Dior"}})
	movies = append(movies, Movie{ID: "5", Isbn: " 127", Title: "Movie Five", Director: &Director{firstName: "Peter", lastName: "Thiel"}})

	r.HandleFunc("/movies", getMovies).Methods("GET")
	r.HandleFunc("/movies/{id}", getMovie).Methods("GET")
	r.HandleFunc("/movies", createMovie).Methods("POST")
	r.HandleFunc("/movies/{id}", deleteMovie).Methods("DELETE")
	r.HandleFunc("/movies/{id}", updateMovie).Methods("PUT")

	fmt.Printf("Start the server at port 8000\n")
	log.Fatal(http.ListenAndServe(":8000", r))

}
