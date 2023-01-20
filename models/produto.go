package models

import (
	"GOWEB/db"
)

type Produto struct {
	Id         int
	Nome       string
	Descricao  string
	Preco      float64
	Quantidade int
}

func BuscaTodosProdutos() []Produto {
	db := db.ConectaComBanco()
	todosProdutos, err := db.Query("select * from produtos")
	produto := Produto{}
	produtos := []Produto{}

	if err != nil {
		panic(err.Error())
	}

	for todosProdutos.Next() {

		var id, quantidade int
		var descricao, nome string
		var preco float64

		err := todosProdutos.Scan(&id, &nome, &descricao, &preco, &quantidade)
		if err != nil {
			panic(err.Error())
		}

		produto.Id = id
		produto.Nome = nome
		produto.Descricao = descricao
		produto.Preco = preco
		produto.Quantidade = quantidade

		produtos = append(produtos, produto)
	}
	defer db.Close()
	return produtos
}

func InserirProduto(nome, descricao string, preco float64, quantidade int) {
	db := db.ConectaComBanco()
	insereDados, err := db.Prepare("insert into produtos (nome, descricao, preco, quantidade) values ($1, $2, $3, $4)")
	if err != nil {
		panic(err.Error())
	}
	insereDados.Exec(nome, descricao, preco, quantidade)
	defer db.Close()

}
