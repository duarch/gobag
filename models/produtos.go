package models

import "github.com/duarch/gobag/db"

type Produto struct {
	Nome, Descricao string
	Preco           float64
	id, Quantidade  int
}

// "github.com/duarch/gobag/db"

func BuscaTodosProdutos() []Produto {
	database := db.ConnectDB()
	selectProdutos, err := database.Query("SELECT * FROM produtos")
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

	defer database.Close()
	return produtos
}
