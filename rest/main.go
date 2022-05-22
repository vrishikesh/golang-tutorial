package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Article struct {
	Title   string `json:"Title"`
	Desc    string `json:"desc"`
	Content string `json:"content"`
}

func fetchArticles(w http.ResponseWriter, r *http.Request) {
	articles := []Article{
		{
			Title:   "Test Title",
			Desc:    "Test Description",
			Content: "Hello World",
		},
	}

	fmt.Println("Endpoint Hit: All Articles Endpoint")
	json.NewEncoder(w).Encode(articles)
}

func createArticle(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Test POST endpoint worked")
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "HomePage Endpoint Hit")
}

func handleRequests() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", homePage)
	router.HandleFunc("/articles", fetchArticles).Methods("GET")
	router.HandleFunc("/articles", createArticle).Methods("POST")
	log.Fatal(http.ListenAndServe(":8081", router))
}

func main() {
	handleRequests()
}
