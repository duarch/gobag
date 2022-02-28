package models

import "duarch/gobag/db"

type Produto struct {
	Nome, Descricao string
	Preco           float64
	Id, Quantidade  int
}

func BuscaProdutos() []Produto {
	db := db.ConnectDB()
	selectProdutos, err := db.Query("SELECT * FROM produtos ORDER BY id ASC")
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
		p.Id = id
		p.Nome = nome
		p.Descricao = descricao
		p.Preco = preco
		p.Quantidade = quantidade
		produtos = append(produtos, p)

	}
	defer db.Close()
	return produtos
}

func InsereProduto(nome, descricao string, preco float64, quantidade int) {
	db := db.ConnectDB()
	insertProduto, err := db.Prepare("INSERT INTO produtos(nome, descricao, preco, quantidade) VALUES($1, $2, $3, $4)")
	if err != nil {
		panic(err)
	}
	_, err = insertProduto.Exec(nome, descricao, preco, quantidade)
	if err != nil {
		panic(err)
	}
	defer db.Close()
}

func DeletaProduto(id string) {
	db := db.ConnectDB()
	deleteProduto, err := db.Prepare("DELETE FROM produtos WHERE id=$1")
	if err != nil {
		panic(err)
	}
	_, err = deleteProduto.Exec(id)
	if err != nil {
		panic(err)
	}
	defer db.Close()
}

func EditaProduto(id string) Produto {
	db := db.ConnectDB()
	ProdDB, err := db.Prepare("SELECT * FROM produtos WHERE id=$1")
	if err != nil {
		panic(err)
	}
	p := Produto{}
	ProdDB.QueryRow(id).Scan(&p.Id, &p.Nome, &p.Descricao, &p.Preco, &p.Quantidade)
	defer db.Close()
	return p

}

func AtualizaProduto(id int, nome, descricao string, preco float64, quantidade int) {
	db := db.ConnectDB()
	updateProduto, err := db.Prepare("UPDATE produtos SET nome=$1, descricao=$2, preco=$3, quantidade=$4 WHERE id=$5")
	if err != nil {
		panic(err)
	}
	_, err = updateProduto.Exec(nome, descricao, preco, quantidade, id)
	if err != nil {
		panic(err)
	}
	defer db.Close()
}
