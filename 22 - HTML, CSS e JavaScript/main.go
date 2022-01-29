package main

import (
	"fmt"
	"html/template" // pacote que permite criar templates baseados em html
	"log"
	"net/http"
)

var templates *template.Template

type usuario struct {
	Nome  string
	Email string
}

func main() {

	// busca todos os templates
	templates = template.Must(template.ParseGlob("*.html"))

	http.HandleFunc("/home", func(w http.ResponseWriter, r *http.Request) {
		u := usuario{"Maria", "maria.joao@gmail.com"} // dado a ser apresentado
		templates.ExecuteTemplate(w, "home.html", u)
	})

	fmt.Println("Escutando na porta 5000")
	log.Fatal(http.ListenAndServe(":5000", nil))
}
