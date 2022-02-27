package main

import (
	"net/http"
	"text/template"
)

type Produto struct {
	Nome, Descricao string
	Preco           float64
	Quantidade      int
}

var temp = template.Must(template.ParseGlob("templates/*.html"))

func main() {
	http.HandleFunc("/", index)
	http.ListenAndServe(":8000", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	produtos := []Produto{
		{"Camiseta", "Azul", 39, 1},
		{"Calça", "Vermelha", 79, 3},
		{"Camiseta", "Verde", 29, 5},
		{"Camiseta", "Amarela", 19, 2},
	}
	temp.ExecuteTemplate(w, "Index", produtos)
}
