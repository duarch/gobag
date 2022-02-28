package controllers

import (
	m "duarch/gobag/models"
	"net/http"
	"text/template"
)

var temp = template.Must(template.ParseGlob("templates/*.html"))

func Index(w http.ResponseWriter, r *http.Request) {
	TodosProdutos := m.BuscaProdutos()
	temp.ExecuteTemplate(w, "Index", TodosProdutos)

}
