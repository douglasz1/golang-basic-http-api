package main

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// User is a struct that represents a user in our application
type User struct {
	FullName string `json:"fullname"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

// Post is a struct that represents a single post
type Post struct {
	Title  string `json:"title"`
	Body   string `json:"body"`
	Author User   `json:"author"`
}

var posts []Post = []Post{}

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/posts", addItem).Methods("POST")

	router.HandleFunc("/posts", getAllPosts).Methods("GET")

	router.HandleFunc("/posts/{id}", getPost).Methods("GET")

	http.ListenAndServe(":5000", router)
}

func getPost(w http.ResponseWriter, r *http.Request) {
	// get the ID of the post from the route parameter
	var idParam string = mux.Vars(r)["id"]
	id, err := strconv.Atoi(idParam)
	if err != nil {
		// there was an error
		w.WriteHeader(400)
		w.Write([]byte("ID could not be converted to integer"))
		return
	}

	// error checking
	if id >= len(posts) {
	}
}

func getAllPosts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(posts)
}

func addItem(w http.ResponseWriter, r *http.Request) {
	// get Item value from the JSON body
	var newPost Post
	json.NewDecoder(r.Body).Decode(&newPost)
	posts = append(posts, newPost)

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(posts)
}
