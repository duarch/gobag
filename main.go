package main

import (
	"net/http"
	"text/template"

	m "duarch/gobag/models"

	_ "github.com/lib/pq"
)

var temp = template.Must(template.ParseGlob("templates/*.html"))

func main() {
	http.HandleFunc("/", index)
	http.ListenAndServe(":8000", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	TodosProdutos := m.BuscaProdutos()
	temp.ExecuteTemplate(w, "Index", TodosProdutos)

}
