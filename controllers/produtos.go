package controllers

import (
	"encoding/json"
	"net/http"
	"errors"

	"github.com/darahayes/go-boom"
	"github.com/Matari73/APIGo/validators"
	"github.com/Matari73/APIGo/adapters/produtos"
	"github.com/gorilla/mux"
)

//produtos
func GetProdutos(w http.ResponseWriter, r *http.Request) {
	produtos , err := adapters.BuscarProdutos()
	if err != nil {
		boom.BadImplementation(w, err)
		return
	}

	if err := adapters.CodificarRespostaProduto(w, produtos) ; err != nil {
		boom.BadImplementation(w, err)
		return
	}
}

func GetProduto(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	
	p, err := adapters.BuscarProdutoById(id)
	if err != nil {
		boom.NotFound(w, err)
		return
	}

	if err := codificarEmJson(w, p); err != nil {
		erro:= errors.New("Erro ao codificar o produto em JSON")
		boom.BadRequest(w, erro)
		return
	}
}

func CreateProduto(w http.ResponseWriter, r *http.Request) {
	novoProduto, err := adapters.LerCorpoRequisicao(r)
	if  err != nil {
		boom.BadRequest(w, err)
		return
	}

	if err:= validators.ValidateProduto(novoProduto); err != nil {
		boom.BadRequest(w, err)
		return
	}

	// Verificar se o nome do produto já existe
	erro := adapters.VerificarSeExiste(novoProduto)
	if  erro != nil {
		boom.BadRequest(w, erro)
		return
	}

	novoProduto, err = adapters.CriarProduto(novoProduto)
	if err == nil {
		boom.BadImplementation(w, err)
		return
	}

	if err := codificarEmJson(w, novoProduto); err != nil {
		erro:= errors.New("Erro ao codificar o cliente em JSON")
		boom.BadImplementation(w, erro)
		return
	}
}

func DeleteProduto(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	_ , err:= adapters.BuscarProdutoById(id)
	if  err != nil {
		boom.BadRequest(w, err)
		return
	}

	result := adapters.DeletarProduto(id)
	if result != nil {
		boom.BadImplementation(w, result)
		return
	}

	w.WriteHeader(http.StatusCreated)
	sucesso:= CreateResposta("Produto excluido com sucesso!")
	json.NewEncoder(w).Encode(sucesso)
}

func UpdateProduto(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	produto, err := adapters.LerCorpoRequisicao(r)
	if err != nil {
		boom.BadRequest(w, err)
		return
	}

	if err:= validators.ValidateProduto(produto); err != nil {
		boom.BadRequest(w, err)
		return
	}

	_ , erro:= adapters.BuscarProdutoById(id)
	if  erro != nil {
		boom.BadRequest(w, erro)
		return
	}

	result := adapters.AtualizarProduto(produto,id)
	if result != nil {
		boom.BadImplementation(w, erro)
		return
	}

	if err := codificarEmJson(w, produto); err != nil {
		erro:= errors.New("Erro ao codificar o produto em JSON")
		boom.BadImplementation(w, erro)
		return
	}
}
