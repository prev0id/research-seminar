package main

import (
	_ "embed"
	"html/template"
	"log"
	"net/http"
)

//go:embed index.go.html
var htmlTemplate []byte

func main() {
	tmpl := template.Must(template.New("roadmap").Parse(string(htmlTemplate)))

	http.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		err := tmpl.Execute(w, roadmap)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	log.Println("Server started at: http://localhost:8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err.Error())
	}
}
