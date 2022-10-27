package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Movie struct {
	ID string `json:"id"`
	ISBN string `json:"isbn"`
	Title string `json:"title"`
	Director *Director `json:"director"`
}

type Director struct {
	FirstName string  `json:"firstname"`
	LastName string  `json:"lastname"`
}

var movies []Movie

func getAllMovies(wr http.ResponseWriter, req *http.Request) {
	wr.Header().Set("Content-Type", "application/json")
	json.NewEncoder(wr).Encode(movies)
}

func getMovie(wr http.ResponseWriter, req *http.Request) {
	wr.Header().Set("Content-Type", "application/json")
	params := mux.Vars(req)
	for _, item := range movies {
		if item.ID == params["id"] {
			json.NewEncoder(wr).Encode(item)
			return
		}
	}
}

func createMovie(wr http.ResponseWriter, req *http.Request) {
	wr.Header().Set("Content-Type", "application/json")
	var movie Movie
	_ = json.NewDecoder(req.Body).Decode(&movie)
	movie.ID = strconv.Itoa(rand.Intn(100000000))
	movies = append(movies, movie)
	json.NewEncoder(wr).Encode(movie)
}

func updateMovie(wr http.ResponseWriter, req *http.Request) {
	wr.Header().Set("Content-Type", "application/json")
	params := mux.Vars(req)
	for index, item := range movies {
		if item.ID == params["id"] {
			movies = append(movies[:index], movies[index+1:]...)
			var movie Movie
			_ = json.NewDecoder(req.Body).Decode(&movie)
			movie.ID = params["id"]
			movies = append(movies, movie)
			json.NewEncoder(wr).Encode(movie)
		}
	}
}

func deleteMovie(wr http.ResponseWriter, req *http.Request) {
	wr.Header().Set("Content-Type", "application/json")
	params := mux.Vars(req)
	for index, item := range movies {
		if item.ID == params["id"] {
			movies = append(movies[:index], movies[index+1:]...)
			break
		}
	}
}

func main() {

	router := mux.NewRouter()

	director1 := Director{FirstName:"Sanghun", LastName:"Song"}
	director2 := Director{FirstName:"Taehee", LastName:"Jeong"}

	movies = append(movies, Movie{ID:"MOV_NUM_1", ISBN:"123122868462", Title:"YEE", Director: &director1})
	movies = append(movies, Movie{ID:"MOV_NUM_2", ISBN:"145142462454", Title:"WOWOW", Director: &director1})
	movies = append(movies, Movie{ID:"MOV_NUM_3", ISBN:"537527845214", Title:"JACKEE", Director: &director2})

	router.HandleFunc("/movies", getAllMovies).Methods("GET")
	router.HandleFunc("/movies/{id}", getMovie).Methods("GET")
	router.HandleFunc("/movies", createMovie).Methods("POST")
	router.HandleFunc("/movies{id}", updateMovie).Methods("PUT")
	router.HandleFunc("/movies{id}", deleteMovie).Methods("DELETE")

	fmt.Print("Starting Server at prot 8080\n")
	log.Fatal(http.ListenAndServe(":8080", router))
}