package main

//https://www.youtube.com/watch?v=SonwZ6MF5BE&t=6s
import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Article struct {
	ID      string `json:"id"`
	Title   string `json:"title"`
	Desc    string `json:"desc"`
	Content string `json:"content"`
}

var Articles []Article

func getArticles(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(Articles)
}
func getArticle(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	i := mux.Vars(r)

	// for undefine, item to the range of Article
	for _, item := range Articles {
		if item.ID == i["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Article{})
}

func createNewArticle(w http.ResponseWriter, r *http.Request) {

}

func updateArticle(w http.ResponseWriter, r *http.Request) {
}

func deleteArticle(w http.ResponseWriter, r *http.Request) {

}

func handlers() {
	myRouter := mux.NewRouter().StrictSlash(true)

	myRouter.HandleFunc("/articles", getArticles).Methods("GET")
	myRouter.HandleFunc("/article/{id}", getArticle).Methods("GET")
	myRouter.HandleFunc("/article", createNewArticle).Methods("POST")
	myRouter.HandleFunc("/articles/{id}", updateArticle).Methods("Put")
	myRouter.HandleFunc("/articles/{id}", deleteArticle).Methods("DELETE")
	myRouter.PathPrefix("/").Handler(http.StripPrefix("/", http.FileServer(http.Dir("static/"))))

	fmt.Printf("Starting server at port 8080\n")
	log.Fatal(http.ListenAndServe(":8080", myRouter))

}

func main() {
	Articles = []Article{
		Article{ID: "1", Title: "Hello", Desc: "Article Description", Content: "Article"},
		Article{ID: "2", Title: "Hello2", Desc: "Article Description", Content: "Article"},
	}
	handlers()

}
