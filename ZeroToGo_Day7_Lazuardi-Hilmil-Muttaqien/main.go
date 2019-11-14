package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type article struct {
	ID      int
	Title   string
	Content string
}

var data = []article{
	article{1, "one piece", "anime"},
	article{2, "donald duck", "fantasy"},
}

type AnimeChar struct {
	Name      string   `json:"name"`
	Height    string   `json:"height"`
	Mass      string   `json:"mass"`
	HairColor string   `json:"hair_color"`
	SkinColor string   `json:"skin_color"`
	EyeColor  string   `json:"eye_color"`
	BirthYear string   `json:"birth_year"`
	Gender    string   `json:"gender"`
	Film      []string `json:"films"`
}

func articles(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.Method == "GET" {
		var result, err = json.Marshal(data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Write(result)
		return
	}
	http.Error(w, "", http.StatusBadRequest)
}

func xerror(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		http.Error(w, "response code: 99", http.StatusBadRequest)
		return
	}
	http.Error(w, "response code: 89", http.StatusBadRequest)
	return

}

func main() {
	//no 2
	// http.HandleFunc("/articles", articles)
	// http.HandleFunc("/404", xerror)
	// //fmt.Println("starting web server at http://localhost:8080/")
	// fmt.Println("server sampun mekan at http://localhost:8080/")
	// http.ListenAndServe(":8080", nil)

	//no 3
	response, _ := http.Get("https://swapi.co/api/people/1")
	responseData, _ := ioutil.ReadAll(response.Body)
	defer response.Body.Close()

	var Luffy AnimeChar
	json.Unmarshal(responseData, &Luffy)

	fmt.Println("--- Print Field ---")
	fmt.Println(Luffy.Name)
	fmt.Println(Luffy.Height)
	fmt.Println(Luffy.Mass)
	fmt.Println(Luffy.HairColor)
	fmt.Println(Luffy.Film[0])
	fmt.Println(Luffy.BirthYear)
}
