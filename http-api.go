package main

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

var data []string = []string{}

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/test", test)

	router.HandleFunc("/add/{data}", addItem)

	http.ListenAndServe(":5000", router)
}

func test(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(struct {
		ID string
	}{ID: "555"})
}

fun addItem(w http.ResponseWriter, r *http.Request) {
	routeVariable := mux.Vars(r)["data"]
	data = append(data, routeVariable)
}
