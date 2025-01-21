package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Go ArtlistPirate")

	// Handler for the root path
	h1 := func(w http.ResponseWriter, r *http.Request) {
		templ := template.Must(template.ParseFiles("index.html"))
		templ.Execute(w, nil)
	}

	// Serve static files from the "css" directory
	http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("css"))))

	// Serve static files from the "assets" directory
	fs := http.FileServer(http.Dir("assets"))
	http.Handle("/assets/", http.StripPrefix("/assets/", fs))

	musicFS := http.FileServer(http.Dir("music"))
	http.Handle("/music/", http.StripPrefix("/music/", musicFS))

	// Handle the root path
	http.HandleFunc("/", h1)

	// Start the server
	log.Fatal(http.ListenAndServe(":8000", nil))
}
