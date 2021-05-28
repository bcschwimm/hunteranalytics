// HunterAnalytics API will collect/serve 5 data points. Time Spent Playing, Training, Exercising, Woofing, Date.

package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// Metrics are datapoints tracked in miniutes for hunters behavoir
type Metrics struct {
	Playing    int    `json:"playing"`
	Training   int    `json:"training"`
	Exercising int    `json:"exercising"`
	Woofing    int    `json:"woofing"`
	Date       string `json:"date"`
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/api", HunterAPI)

	log.Fatal(http.ListenAndServe(":8080", nil))
}

// index is our main page of HunterAnalytics
func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hunter Analytics is under construction")
}

// HunterAPI is using our Metrics struct to store and serve data
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
