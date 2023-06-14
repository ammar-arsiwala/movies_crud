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
fatal.log(http.ListenAndServeTLS(":8000,r"))
}

func GetMovies(w http.ResponseWriter,r *http.Request){
	w.Header().Set("Content-Type","applcation/json")
	json.NewEncoder(w).Encode(movies)
}


func DeleteMovies(w http.ResponseWriter,r *http.Request){
	w.Header().Set("Content-Type","application/json")
	params:= Mux.vars(r)
	for index, item:= range movies{
		if items.ID == params["id"]{
			movies=append(movies[:index], movies[index+1:]...)
			break
		} 
	
	}
}


func GetMovies( w http.ResponseWriter,r * http.Request){
	w.Header().Set("Content-Type","application/json")
	params:= Mux.vars(r)
	for _,item:= range movies
	{
		if items.ID == params["id"]{
			json.NewEncoder(w).Encode(items)
			writer(w).WriteHeader(http.StatusOk)
			return movie
		}
	}
	
}

func CreateMovies(w http.ResponseWriter, r * http.Request){
	w.Header().Set("Content-Type","application/json")
	writer.WriteHeader(http.StatusOk)
	var movie Movie
	_ = json.NewDecoder(r.Body).Decode(&movie)
		movie.ID = stringconv.Itoa(rand.Intn(100000000))
		movies = append(movies,movie)
		json.NewEncoder(w).Encode(movie)
}


func UpdateMovies(w http.ResponseWriter, r * http.Request){
	w.Heaser().Set("Content-Type", "application/json")
	params:= Mux.vars(r)
	for index,item:= range movies
	{
		if items.ID == params["id"]{
			movies = append(movies[:index],movies[index + 1:]...)
			var Movie movies
			_=json.NewDecoder(r.body)Decode(&movie)
			  movie.ID = params["id"]
			  movies = append(movies,movie)
			  json.NewEncoder(w).Encode(movie)		
		}
	}
