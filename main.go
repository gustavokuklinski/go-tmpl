package main

import (
	"html/template"
	"log"
	"net/http"
)

// Name struct store form request
type Name struct {
	FName string
}

// Global variable to parse all templates
var goTmpl = template.Must(template.ParseGlob("templates/*"))

// Index calls the main page
func Index(w http.ResponseWriter, r *http.Request) {

	// Render the "Index.gtpl"
	goTmpl.ExecuteTemplate(w, "Index", nil)

	//name := r.FormValue("name")

	// Logging the Action
	log.Println("Render: Index.gtpl")

}

// Show name request from Index
func Show(w http.ResponseWriter, r *http.Request) {

	// Request form value into 'name' variable
	name := r.FormValue("name")

	// Save the request inside struct
	sName := Name{FName: name}

	// Display template with the struct variable
	goTmpl.ExecuteTemplate(w, "Show", sName)

	// Logging the Action
	log.Println("Render: Show.gtpl")
	log.Println("Requested value: " + name)
}

func main() {
	// Route Index page
	http.HandleFunc("/", Index)

	// Route Show page
	http.HandleFunc("/show", Show)

	// Start HTTP server
	http.ListenAndServe(":9000", nil)

}
