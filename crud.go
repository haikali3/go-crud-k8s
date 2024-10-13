package main

import (
	"encoding/json"
	"log"
	"net/http"
	"github.com/gorilla/mux"
)

type Item struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

var items []Item

func createItem(w http.ResponseWriter, r *http.Request) {
	var item Item
	_ = json.NewDecoder(r.Body).Decode(&item)
	items = append(items, item)
	json.NewEncoder(w).Encode(item)
}

func getItems(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(items)
}

func updateItem(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for index, item := range items {
		if item.ID == params["id"] {
			items = append(items[:index], items[index+1:]...)
			var updatedItem Item
			_ = json.NewDecoder(r.Body).Decode(&updatedItem)
			updatedItem.ID = params["id"]
			items = append(items, updatedItem)
			json.NewEncoder(w).Encode(updatedItem)
			return
		}
	}
}

func deleteItem(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for index, item := range items {
		if item.ID == params["id"] {
			items = append(items[:index], items[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(items)
}

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/items", createItem).Methods("POST")
	router.HandleFunc("/items", getItems).Methods("GET")
	router.HandleFunc("/items/{id}", updateItem).Methods("PUT")
	router.HandleFunc("/items/{id}", deleteItem).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000", router))
}

