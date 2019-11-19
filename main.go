package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

// Article structure is the structure of an article
type Article struct {
	Title   string `json:"title"`
	Desc    string `json:"description"`
	Content string `json:"content"`
}

// Articles array
type Articles []Article

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Print(w, "Hit enpoint now")
}

func allArticles(w http.ResponseWriter, r *http.Request) {
	articles := Articles{
		Article{
			Title:   "Google's go",
			Desc:    "The new lang",
			Content: "Lorem ipsum",
		},
	}
	json.NewEncoder(w).Encode(articles)
}

func handleRequests() {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/", homePage).Methods("GET")
	router.HandleFunc("/articles", allArticles).Methods("GET")

	log.Fatal(http.ListenAndServe(":8000", router))
}

func main() {
	handleRequests()
}
