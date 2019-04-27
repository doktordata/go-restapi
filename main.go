package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

// Vehichle Struct
type Vehichle struct {
	ID      string `json:"id"`
	Make    string `json:"make"`
	Model   string `json:"model"`
	Vintage string `json:"vintage"`
}

// Make Struct
// type Make struct {
// 	Name string `json:"name"`
// }

var db *sql.DB

func connectDb() {
	var err error

	db, err = sql.Open("postgres", "user=erik dbname=bumper sslmode=disable")
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}
}

func getVehichles(w http.ResponseWriter, r *http.Request) {
	var vehichles []Vehichle

	rows, err := db.Query(`SELECT id, make, model, vintage FROM vehichles`)
	if err != nil {
		panic(err)
	}

	for rows.Next() {
		vehichle := Vehichle{}

		err = rows.Scan(&vehichle.ID, &vehichle.Make, &vehichle.Model, &vehichle.Vintage)
		if err != nil {
			panic(err)
		}

		vehichles = append(vehichles, vehichle)
	}

	err = rows.Err()
	if err != nil {
		panic(err)
	}

	w.Header().Set("Content-Type", "application/json")
	if len(vehichles) != 0 {
		json.NewEncoder(w).Encode(vehichles)
	} else {
		w.WriteHeader(404)
	}
}

func getVehichle(w http.ResponseWriter, r *http.Request) {
	vehichle := Vehichle{}

	id := mux.Vars(r)["id"]
	row := db.QueryRow(`SELECT id, make, model, vintage FROM vehichles WHERE id=$1`, id)
	err := row.Scan(&vehichle.ID, &vehichle.Make, &vehichle.Model, &vehichle.Vintage)

	if err != nil {
		if err == sql.ErrNoRows {
			w.WriteHeader(404)
			return
		} else {
			w.WriteHeader(500)
			return
		}
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(vehichle)
}

func main() {
	connectDb()
	defer db.Close()

	r := mux.NewRouter()

	r.HandleFunc("/api/vehichles", getVehichles).Methods("GET")
	r.HandleFunc("/api/vehichles/{id}", getVehichle).Methods("GET")

	log.Fatal(http.ListenAndServe(":8000", r))
}
