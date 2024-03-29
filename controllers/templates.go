package controllers

import (
	m "duarch/gobag/models"
	"log"
	"net/http"
	"strconv"
	"text/template"
)

var temp = template.Must(template.ParseGlob("templates/*.html"))

func Index(w http.ResponseWriter, r *http.Request) {
	TodosProdutos := m.BuscaProdutos()
	temp.ExecuteTemplate(w, "Index", TodosProdutos)

}
func New(w http.ResponseWriter, r *http.Request) {
	temp.ExecuteTemplate(w, "New", nil)

}
func Insert(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		nome := r.FormValue("nome")
		descricao := r.FormValue("descricao")
		preco := r.FormValue("preco")
		quantidade := r.FormValue("quantidade")

		precoConv, err := strconv.ParseFloat(preco, 64)
		if err != nil {
			log.Println("Erro na conversao do preco", err)
		}
		quantidadeConv, err := strconv.Atoi(quantidade)
		if err != nil {
			log.Println("Erro na conversao da quantidade", err)
		}

		m.InsereProduto(nome, descricao, precoConv, quantidadeConv)
	}

	http.Redirect(w, r, "/", 301)

}

func Delete(w http.ResponseWriter, r *http.Request) {
	idProd := r.URL.Query().Get("id")
	m.DeletaProduto(idProd)
	http.Redirect(w, r, "/", 301)

}

func Edit(w http.ResponseWriter, r *http.Request) {
	idProd := r.URL.Query().Get("id")
	produto := m.EditaProduto(idProd)
	temp.ExecuteTemplate(w, "Edit", produto)

}

func Update(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		idProd := r.FormValue("id")
		nome := r.FormValue("nome")
		descricao := r.FormValue("descricao")
		preco := r.FormValue("preco")
		quantidade := r.FormValue("quantidade")

		precoConv, err := strconv.ParseFloat(preco, 64)
		if err != nil {
			log.Println("Erro na conversao do preco", err)
		}
		idProdConv, err := strconv.Atoi(idProd)
		if err != nil {
			log.Println("Erro na conversao do id", err)
		}
		quantidadeConv, err := strconv.Atoi(quantidade)
		if err != nil {
			log.Println("Erro na conversao da quantidade", err)
		}

		m.AtualizaProduto(idProdConv, nome, descricao, precoConv, quantidadeConv)
	}
	http.Redirect(w, r, "/", 301)

}
