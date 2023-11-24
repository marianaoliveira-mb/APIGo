package controllers

import (
	"encoding/json"
	"net/http"
	"errors"

	"github.com/darahayes/go-boom"
	// "github.com/Matari73/APIGo/database"
	// "github.com/Matari73/APIGo/models"
	"github.com/Matari73/APIGo/adapters/pedidos"
	"github.com/Matari73/APIGo/validators"
	"github.com/gorilla/mux"
)

func GetPedidos(w http.ResponseWriter, r *http.Request) {
	p, err := adapters.BuscarPedidos()
	if err != nil {
		boom.BadImplementation(w, err)
		return
	}

	if err := adapters.CodificarRespostaPedido(w, p); err != nil {
		boom.BadImplementation(w, err)
		return
	}
}
func GetPedido(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	pedido, err := adapters.BuscarPedidoById(id) 
	if err != nil {
		boom.NotFound(w, err)
		return
	}

	if err := codificarEmJson(w, pedido); err != nil {
		erro:= errors.New("Erro ao codificar o pedido em JSON")
		boom.BadRequest(w, erro)
		return
	}
}

func CreatePedido(w http.ResponseWriter, r *http.Request) {
	novoPedido, err:= adapters.LerCorpoRequisicao(r)
	if err != nil {
		boom.BadRequest(w, err)
		return
	}

	if err:= validators.ValidatePedido(novoPedido); err != nil {
		boom.BadRequest(w, err)
		return
	}

	novoPedido, err = adapters.CriarPedido(novoPedido)
	if err != nil {
		boom.BadImplementation(w, err)
		return
	}

	if err := validators.AtualizarSaldoCliente(uint(novoPedido.ClienteID), novoPedido.ValorPedido); err != nil {
		boom.BadImplementation(w, err)
		return
	}

	if err := codificarEmJson(w, novoPedido); err != nil {
		erro:= errors.New("Erro ao codificar o pedido em JSON")
		boom.BadRequest(w, erro)
		return
	}
}

func DeletePedido(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	_ , err:= adapters.BuscarPedidoById(id)
	if  err != nil {
		boom.BadRequest(w, err)
		return
	}

	result := adapters.DeletarPedido(id)
	if result != nil {
		boom.BadImplementation(w, result)
		return
	}

	w.WriteHeader(http.StatusOK)
	sucesso:= CreateResposta("Pedido excluído com sucesso!")
	json.NewEncoder(w).Encode(sucesso)
}

func UpdatePedido(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	pedido, err := adapters.LerCorpoRequisicao(r)
	if err != nil {
		boom.BadRequest(w, err)
		return
	}

	if err:= validators.ValidatePedido(pedido); err != nil {
		boom.BadRequest(w, err)
		return
	}

	_ , erro := adapters.BuscarPedidoById(id)
	if  erro != nil {
		boom.BadRequest(w, erro)
		return
	}

	result:= adapters.AtualizarPedido(pedido, id)
	if result != nil {
		boom.BadImplementation(w, erro)
		return
	}

	if err := codificarEmJson(w, pedido); err != nil {
		erro:= errors.New("Erro ao codificar o pedido em JSON")
		boom.BadImplementation(w, erro)
		return
	}
}
