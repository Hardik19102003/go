package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

// Serves the index page
func indexHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("static/index.html")
	if err != nil {
		http.Error(w, "Error loading index page", http.StatusInternalServerError)
		return
	}
	tmpl.Execute(w, nil)
}

// Serves the form page
func formPageHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("static/form.html")
	if err != nil {
		http.Error(w, "Error loading form", http.StatusInternalServerError)
		return
	}
	tmpl.Execute(w, nil)
}

// Handles form submission
func formHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	// Parse form data
	if err := r.ParseForm(); err != nil {
		http.Error(w, fmt.Sprintf("Error parsing form: %v", err), http.StatusBadRequest)
		return
	}

	// Retrieve form values
	name := r.FormValue("name")
	address := r.FormValue("address")

	// Return the response
	w.Header().Set("Content-Type", "text/plain")
	fmt.Fprintf(w, "Form submitted successfully!\n")
	fmt.Fprintf(w, "Name: %s\nAddress: %s\n", name, address)

	log.Printf("Received form submission - Name: %s, Address: %s\n", name, address)
}

// Simple Hello API
func helloHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed.", http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-Type", "text/plain")
	fmt.Fprint(w, "Hello World!")
}

func main() {
	// Serve static files (CSS, JS, etc.)
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	// Route Handlers
	http.HandleFunc("/", indexHandler)        // Serves index.html
	http.HandleFunc("/form", formPageHandler) // Serves form.html
	http.HandleFunc("/submit", formHandler)   // Handles form submission
	http.HandleFunc("/hello", helloHandler)   // Simple Hello API

	fmt.Println("Server started on http://localhost:8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("Server failed:", err)
	}
}
