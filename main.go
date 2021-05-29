// HunterAnalytics API will collect/serve 5 data points. Time Spent Playing, Training, Exercising, Woofing, Date.
// use regex must compile to validate form

package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func main() {
	fmt.Println("Server Starting on Port 8080...")

	http.HandleFunc("/", index)
	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("assets"))))
	http.HandleFunc("/api", HunterAPI)
	http.HandleFunc("/form", formData)

	log.Fatal(http.ListenAndServe(":8080", nil))
}

// index serves our index.html file to the homepage
func index(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "assets/index.html")
}

// formData is recieving a post request from html, parsing
// that data and redirecting the user to the home page
func formData(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	if r.Method == "POST" {
		formData := Metrics{
			Playing:    intConv(r.FormValue("playing")),
			Training:   intConv(r.FormValue("training")),
			Exercising: intConv(r.FormValue("exercising")),
			Woofing:    intConv(r.FormValue("woofing")),
			Date:       r.FormValue("date"),
		}
		// insert into db here
		fmt.Println(formData)
		http.Redirect(w, r, "/", http.StatusFound)
	}
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

func intConv(formSubmission string) int {
	i, err := strconv.Atoi(formSubmission)
	if err != nil {
		fmt.Println("error converting string from html form", err)
	}
	return i
}
