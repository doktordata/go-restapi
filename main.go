package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// Vehichle Struct
type Vehichle struct {
	ID      string `json:"id"`
	Make    *Make  `json:"make"`
	Model   string `json:"model"`
	Vintage string `json:"vintage"`
}

// Make Struct
type Make struct {
	Name string `json:"name"`
}

var vehichles []Vehichle

func getVehichles(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(vehichles)
}

func getVehichle(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	for _, vehichle := range vehichles {
		if vehichle.ID == params["id"] {
			json.NewEncoder(w).Encode(vehichle)
			return
		}
	}
	w.WriteHeader(404)
}

func main() {
	r := mux.NewRouter()

	// Mock data
	vehichles = append(vehichles, Vehichle{ID: "1", Model: "911", Vintage: "2000", Make: &Make{Name: "Porsche"}})
	vehichles = append(vehichles, Vehichle{ID: "2", Model: "E320", Vintage: "1997", Make: &Make{Name: "Mercedes-Benz"}})
	vehichles = append(vehichles, Vehichle{ID: "3", Model: "XM", Vintage: "1997", Make: &Make{Name: "Citroën"}})
	vehichles = append(vehichles, Vehichle{ID: "4", Model: "924", Vintage: "1980", Make: &Make{Name: "Porsche"}})
	vehichles = append(vehichles, Vehichle{ID: "5", Model: "E200", Vintage: "1987", Make: &Make{Name: "Mercedes-Benz"}})

	r.HandleFunc("/api/vehichles", getVehichles).Methods("GET")
	r.HandleFunc("/api/vehichles/{id}", getVehichle).Methods("GET")

	log.Fatal(http.ListenAndServe(":8000", r))
}
