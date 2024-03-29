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
	todosProdutos, err := db.Query("select * from produtos order by id asc")
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

func AtualizaProduto(id int, nome, descricao string, preco float64, quantidade int) {
	db := db.ConectaComBanco()
	atulizaDados, err := db.Prepare("update produtos set nome=$1, descricao=$2, preco=$3, quantidade=$4 where id=$5")
	if err != nil {
		panic(err.Error())
	}
	atulizaDados.Exec(nome, descricao, preco, quantidade, id)
	defer db.Close()

}

func DeletarProduto(id string) {
	db := db.ConectaComBanco()
	deletarProduto, err := db.Prepare("delete from produtos where id = $1")
	if err != nil {
		panic(err.Error())
	}

	deletarProduto.Exec(id)
	defer db.Close()
}

func EditaProduto(id string) Produto {
	db := db.ConectaComBanco()
	produtoDoBanco, err := db.Query("select * from produtos where id=$1", id)
	if err != nil {
		panic(err.Error())
	}
	produtoParaAtualizar := Produto{}

	for produtoDoBanco.Next() {
		var id, quantidade int
		var nome, descricao string
		var preco float64

		err = produtoDoBanco.Scan(&id, &nome, &descricao, &preco, &quantidade)
		if err != nil {
			panic(err.Error())
		}

		produtoParaAtualizar.Id = id
		produtoParaAtualizar.Nome = nome
		produtoParaAtualizar.Descricao = descricao
		produtoParaAtualizar.Preco = preco
		produtoParaAtualizar.Quantidade = quantidade
	}
	defer db.Close()

	return produtoParaAtualizar
}
