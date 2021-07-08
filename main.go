package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
)

func main() {
	fmt.Println("Server Starting on Port 8080...")

	// handling "/" by passing our ./assets fileserver
	fs := http.FileServer(http.Dir("./assets"))
	http.Handle("/", fs)

	http.HandleFunc("/api", HunterAPI) // capital H because we are exporting data
	http.HandleFunc("/form", formData)
	http.HandleFunc("/behaviorForm", behaviorData)
	http.HandleFunc("/trainingForm", commandData)
	http.HandleFunc("/tricksapi", TricksAPI)

	log.Fatal(http.ListenAndServe(":8080", nil))
}

// formData is a handler recieving a post request from html, parsing
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
		formData.insert()
		http.Redirect(w, r, "/", http.StatusFound)
	}
}

// behaviorData recieves the post request from .behavior.html
// and inserts that data into mongo db
func behaviorData(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	if r.Method == "POST" {
		behaviorData := Behavior{
			Date:  r.FormValue("date"),
			Crate: intConv(r.FormValue("crate")),
			Notes: r.FormValue("notes"),
		}
		behaviorData.insert()
		http.Redirect(w, r, "/", http.StatusFound)
	}
}

// commandData is parsing our command section of the website
// and if we have a new trick it inserts name, detail, level
// if we have a existing trick it's only submitting the name
func commandData(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	if r.Method == "POST" {
		switch commandStatus := r.FormValue("command"); commandStatus {
		case "existing":
			existingCommand := Trick{
				Name: r.FormValue("name"),
			}
			existingCommand.trainingSessionInsert()

		default:
			newCommand := Trick{
				Name:   strings.ToLower(r.FormValue("name")),
				Detail: r.FormValue("description"),
				Level:  r.FormValue("level"),
			}
			newCommand.trickInsert()
		}
	}
	http.Redirect(w, r, "/", http.StatusFound)
}

// HunterAPI is a handler serving our cloudsql data from the hunter table
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

// TricksAPI is serving all the documents in our tricks collection
// from mongo to the web api "tricksapi"
func TricksAPI(w http.ResponseWriter, r *http.Request) {
	tricks := readTricks("hunter", "tricks")
	trickData, err := json.Marshal(tricks)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println("error: marshalling bson to json from mongo", err)
		return
	}
	w.Header().Set("content-type", "application/json; charset=utf-8")
	if _, err := w.Write(trickData); err != nil {
		log.Println("error writing trick content to response", err)
	}
}
