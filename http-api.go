package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/test", test)
}

func test(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("This is a test")
}
