package main

import (
	"html/template"
	"log"
	"net/http"
)

func main () {
	http.Handle("/dist/", http.StripPrefix("/dist/", http.FileServer(http.Dir("dist"))))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmpl := template.Must(template.ParseFiles("./templates/index.html"))
		tmpl.Execute(w, nil)
	})

	log.Println("App running on port 8000...")
	log.Fatal(http.ListenAndServe(":8000", nil))
	
}