package main

//Random joke : https://official-joke-api.appspot.com/random_joke
//Random joke : https://official-joke-api.appspot.com/jokes/random

//Ten random jokes:https://official-joke-api.appspot.com/random_ten

//random ChuckJokes:https://api.chucknorris.io/jokes/random
//Chuck Category: https://api.chucknorris.io/jokes/random?category={category}
//ChuckCategories:https://api.chucknorris.io/jokes/categories
//Jokes Query:https://api.chucknorris.io/jokes/search?query={query}

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"
	"html/template"
)

type RandomJoke struct {
	ID        int    `json:"id"`
	Type      string `json:"type"`
	Setup     string `json:"setup"`
	Punchline string `json:"punchline"`
}

type ProgrammingJoke struct {
	Category string `json:"category"`
	Type     string `json:"type"`
	Setup    string `json:"setup"`
	Delivery string `json:"delivery"`
	Joke     string `json:"joke"`
	ID       int    `json:"id"`
}

type ChuckJoke struct {
	Categories []interface{} `json:"categories"`
	CreatedAt  string        `json:"created_at"`
	IconURL    string        `json:"icon_url"`
	ID         string        `json:"id"`
	UpdatedAt  string        `json:"updated_at"`
	URL        string        `json:"url"`
	Value      string        `json:"value"`
}

func main(){
	routes := mux.NewRouter()

	routes.HandleFunc("/", Home).Methods("GET")
	routes.HandleFunc("/programming",Programming).Methods("GET")
	routes.HandleFunc("/chuck",ChuckJokes).Methods("GET")
	log.Println(http.ListenAndServe(":8081", routes))
}

func Home(w http.ResponseWriter, r *http.Request) {
	resp,err:=http.Get("https://official-joke-api.appspot.com/random_joke")

	if(err!=nil){
		log.Fatalln(err)
	}
	defer resp.Body.Close()

	body,err:=ioutil.ReadAll(resp.Body)
	if err!=nil {
		log.Fatalln(err)
	}
	log.Println(string(body))

	var randomjoke RandomJoke

	json.Unmarshal(body,&randomjoke)

	var templates = template.Must(template.ParseFiles("views/home.gohtml", "views/base.gohtml"))
	varmap := map[string]interface{}{
		"RandomJoke": randomjoke,
	}

	err = templates.ExecuteTemplate(w, "base", varmap)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

}

func Programming(w http.ResponseWriter,r *http.Request){
	resp2,err:=http.Get("https://sv443.net/jokeapi/category/programming/")

	if (err!=nil){
		log.Fatalln(err)
	}
	defer resp2.Body.Close()

	body,err:=ioutil.ReadAll(resp2.Body)
	if err!=nil {
		log.Fatalln(err)
	}
	log.Println(string(body))

	var codejoke ProgrammingJoke

	json.Unmarshal(body,&codejoke)

	var templates = template.Must(template.ParseFiles("views/codejokes.gohtml", "views/base.gohtml"))
	varmap := map[string]interface{}{
		"CodeJoke": codejoke,
	}

	err = templates.ExecuteTemplate(w, "base", varmap)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}


}

func ChuckJokes(w http.ResponseWriter , r *http.Request){
	resp3,err:=http.Get("https://api.chucknorris.io/jokes/random")

	if (err!=nil){
		log.Fatalln(err)
	}
	defer resp3.Body.Close()

	body,err:=ioutil.ReadAll(resp3.Body)
	if err!=nil {
		log.Fatalln(err)
	}
	log.Println(string(body))

	var chuckjoke ChuckJoke

	json.Unmarshal(body,&chuckjoke)

	var templates = template.Must(template.ParseFiles("views/chuckjokes.gohtml", "views/base.gohtml"))
	varmap := map[string]interface{}{
		"ChuckJoke": chuckjoke,
	}

	err = templates.ExecuteTemplate(w, "base", varmap)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}



