package main

import(
"fmt"
"log"
"encoding/json"
"math/rand"
"net/http"
"strconv"
"github.com/gorilla/mux"
)

type Movie struct{
	ID string  `json:"id"` 
	Isbn string `json:"isbn"`
	Title string `json:"title"`
	Director string  `json:"director"`
}


type Director struct{
	FirstName string  `json:"firstname"`
	LastName string  `json:"lastname"` 
	Country string `json:"country"`
}

var movies []Movie

func main(){
	r:=mux.NewRouter()
	movie=append(movies, Movie{ID: "1",Isbn:"12323", Title: "Armour of God 1",Director: &Director{FirstName:"Jackie",LastName:"Chan",Country:"China"}})	
	movie=append(movies, Movie{ID: "2",Isbn:"12324", Title: "Armour of God 2",Director: &Director{FirstName:"Jackie",LastName:"Chan",Country:"China"}})	
	r.HandleFunc("/movies",GetMovies).Methods(GET)
	r.HandfleFunc("/movies/{id}",GetMovie).Methods(GET)
	r.HandleFunc("/movies",CreateMovies).Methods(CREATE)
	r.HandleFunc("/movies/{id}",UpdateMovie).Methods(UPDATE)
	r.HandleFunc("/movies/{id}",DeleteMovies).Methods(DELETE)

fmt.Printf("starting server at 8000\n")
log.Fatal(http.ListenAndServeTLS(":8000",r))
}

func GetMovies(w http.ResponseWriter,r *http.Request){
	w.Header().Set("Content-Type","applcation/json")
	json.NewEncoder(w).Encode(movies)
}


func DeleteMovies(w http.ResponseWriter,r *http.Request){
	w.Header().Set("Content-Type","application/json")
	params:= mux.Vars(r)
	for index, item:= range movies{
		if item.ID == params["id"]{
			movies=append(movies[:index], movies[index+1:]...)
			break
		} 
	
	}
}


func GetMovies( w http.ResponseWriter,r * http.Request){
	w.Header().Set("Content-Type","application/json")
	params:= mux.Vars(r)
	for _,item:= range movies{
	
		if item.ID == params["id"]{
			json.NewEncoder(w).Encode(item)
			w.WriteHeader(http.StatusOK)
			return Movie
		}
	}
	
}

func CreateMovies(w http.ResponseWriter, r * http.Request){
	w.Header().Set("Content-Type","application/json")
	w.WriteHeader(http.StatusOK)
	var Movie movies
	_ = json.NewDecoder(r.Body).Decode(&Movie)
		Movie.ID = strconv.Itoa(rand.Intn(100000000))
		movies = append(movies,Movie)
		json.NewEncoder(w).Encode(Movie)
}


func UpdateMovies(w http.ResponseWriter, r * http.Request){
	w.Header().Set("Content-Type", "application/json")
	params:= mux.Vars(r)
	for index,item:= range movies{
	
		if items.ID == params["id"]{
			movies = append(movies[:index],movies[index + 1:]...)
			var Movie movies
			_=json.NewDecoder(r.body).Decode(&Movie)
			  Movie.ID = params["id"]
			  movies = append(movies,Movie)
			  json.NewEncoder(w).Encode(Movie)		
		}
	}
}	
