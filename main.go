// HunterAnalytics API will collect/serve 5 data points. Time Spent Playing, Training, Exercising, Woofing, Date.

package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Server Starting on Port 8080...")

	http.HandleFunc("/", index)
	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("assets"))))
	http.HandleFunc("/api", HunterAPI)

	log.Fatal(http.ListenAndServe(":8080", nil))
}

// index serves our index.html file to the homepage
func index(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "assets/index.html")
}

// HunterAPI is serving our cloudsql data from the hunter table
func HunterAPI(w http.ResponseWriter, r *http.Request) {
	db, err := Open()
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	list, err := List(db)
	if err != nil {
		panic(err.Error())
	}

	data, err := json.Marshal(list)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println("error marshalling", err)
		return
	}
	w.Header().Set("content-type", "applicaiton/json; charset=utf-8")
	if _, err := w.Write(data); err != nil {
		log.Println("error writing content to response", err)
	}
}
