package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"
	"text/template"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func getEnvVars() {
	err := godotenv.Load("credentials.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func conectDB() *sql.DB {
	getEnvVars()
	dbName := os.Getenv("DB_NAME")
	dbSecret := os.Getenv("DB_KEY")
	openKey := "user=postgres password=" + dbSecret + " dbname=" + dbName + " host=localhost sslmode=disable"
	db, err := sql.Open("postgres", openKey)
	if err != nil {
		panic(err)
	}
	return db
}

type Produto struct {
	Nome, Descricao string
	Preco           float64
	id, Quantidade  int
}

var temp = template.Must(template.ParseGlob("templates/*.html"))

func main() {
	http.HandleFunc("/", index)
	http.ListenAndServe(":8000", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	db := conectDB()
	selectProdutos, err := db.Query("SELECT * FROM produtos")
	if err != nil {
		panic(err)
	}
	p := Produto{}
	produtos := []Produto{}
	for selectProdutos.Next() {
		var id, quantidade int
		var nome, descricao string
		var preco float64

		err = selectProdutos.Scan(&id, &nome, &descricao, &preco, &quantidade)
		if err != nil {
			panic(err)
		}
		// p.id = id
		p.Nome = nome
		p.Descricao = descricao
		p.Preco = preco
		p.Quantidade = quantidade
		produtos = append(produtos, p)

	}
	temp.ExecuteTemplate(w, "Index", produtos)
	defer db.Close()
}
