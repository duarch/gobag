package models

import "duarch/gobag/db"

type Produto struct {
	Nome, Descricao string
	Preco           float64
	id, Quantidade  int
}

func BuscaProdutos() []Produto {
	db := db.ConnectDB()
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
	defer db.Close()
	return produtos
}
