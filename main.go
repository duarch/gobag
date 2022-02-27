package main

import (
	"duarch/gobag/models"
	"net/http"
	"text/template"
)

var temp = template.Must(template.ParseGlob("templates/*.html"))

func main() {
	http.HandleFunc("/", index)
	http.ListenAndServe(":8000", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	produtos := models.BuscaTodosProdutos
	temp.ExecuteTemplate(w, "Index", produtos)
}
