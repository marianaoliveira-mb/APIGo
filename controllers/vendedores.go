package controllers

import (
	"encoding/json"
	"net/http"
	"errors"

	"github.com/darahayes/go-boom"
	"github.com/Matari73/APIGo/adapters/vendedores"
	"github.com/Matari73/APIGo/validators"
	"github.com/gorilla/mux"
)

func GetVendedores(w http.ResponseWriter, r *http.Request) {
	vendedores , err := adapters.BuscarVendedores()
	if err != nil {
		boom.BadImplementation(w, err)
		return
	}

	if err := adapters.CodificarRespostaVendedor(w, vendedores) ; err != nil {
		boom.BadImplementation(w, err)
		return
	}
}

func GetVendedor(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	v, err := adapters.BuscarVendedorById(id)
	if err != nil {
		boom.NotFound(w, err)
		return
	}

	if err := codificarEmJson(w, v); err != nil {
		erro:= errors.New("Erro ao codificar o vendedor em JSON")
		boom.BadRequest(w, erro)
		return
	}
}

func CreateVendedor(w http.ResponseWriter, r *http.Request) {
	novoVendedor, err := adapters.LerCorpoRequisicao(r)
	if err != nil {
		boom.BadRequest(w, err)
		return
	}
	
	if err:= validators.ValidateNome(novoVendedor); err != nil {
		boom.BadRequest(w, err)
		return
	}

	novoVendedor, err = adapters.CriarVendedor(novoVendedor)
	if err != nil {
		boom.BadImplementation(w, err)
		return
	}

	if err := codificarEmJson(w, novoVendedor); err != nil {
		erro:= errors.New("Erro ao codificar o vendedor em JSON")
		boom.BadImplementation(w, erro)
		return
	}
}

func DeleteVendedor(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	_ , err:= adapters.BuscarVendedorById(id)
	if  err != nil {
		boom.BadRequest(w, err)
		return
	}

	result := adapters.DeletarVendedor(id)
	if result != nil {
		boom.BadImplementation(w, result)
		return
	}

	w.WriteHeader(http.StatusOK)
	sucesso:= CreateResposta("Vendedor excluído com sucesso!")
	json.NewEncoder(w).Encode(sucesso)
}

func UpdateVendedor(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	vendedor, err := adapters.LerCorpoRequisicao(r)
	if err != nil {
		boom.BadRequest(w, err)
		return
	}

	if err:= validators.ValidateNome(vendedor); err != nil {
		boom.BadRequest(w, err)
		return
	}

	_ , erro:= adapters.BuscarVendedorById(id)
	if  erro != nil {
		boom.BadRequest(w, erro)
		return
	}

	result := adapters.AtualizarVendedor(vendedor,id)
	if result != nil {
		boom.BadImplementation(w, erro)
		return
	}

	if err := codificarEmJson(w, vendedor); err != nil {
		erro:= errors.New("Erro ao codificar o vendedor em JSON")
		boom.BadRequest(w, erro)
		return
	}
}
