package controllers

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/darahayes/go-boom"
	"github.com/Matari73/APIGo/adapters/historico"
	"github.com/gorilla/mux"
)

func AdicionarProdutoAoPedidoHandler(w http.ResponseWriter, r *http.Request) {
	var dadosRequisicao struct {
		PedidoID   uint `json:"pedido_id"`
		ProdutoID  uint `json:"produto_id"`
		Quantidade int  `json:"quantidade"`
	}

	err := json.NewDecoder(r.Body).Decode(&dadosRequisicao)
	if err != nil {
		erro:= errors.New("Erro ao decodificar o corpo da requisição")
		boom.BadRequest(w, erro)
	}

	err = adapters.AdicionarProdutoAoPedido(dadosRequisicao.PedidoID, dadosRequisicao.ProdutoID, dadosRequisicao.Quantidade)
	if err != nil {
		boom.BadImplementation(w, err)
	}

	w.WriteHeader(http.StatusCreated)
	sucesso:= CreateResposta("Produto adicionado ao pedido!")
	json.NewEncoder(w).Encode(sucesso)
}

func HistoricoCompras(w http.ResponseWriter, r *http.Request) {
	vars:= mux.Vars(r)
	clienteIDStr , ok := vars["cliente_id"]
	if !ok {
		http.Error(w, "ID do cliente não encontrado na URL", http.StatusBadRequest)
		return
	}

	clienteID, err := strconv.Atoi(clienteIDStr)
	if err != nil {
		http.Error(w, "ID do cliente inválido", http.StatusBadRequest)
		return
	}

	err = adapters.ObterCompras(w, clienteID)
	if err != nil {
		boom.BadImplementation(w, err)
		return
	}
}

func HistoricoVendasVendedor(w http.ResponseWriter, r *http.Request) {
	vars:= mux.Vars(r)
	vendedorIDStr, ok := vars["vendedor_id"]

	if !ok {
		http.Error(w, "ID do vendedor não encontrado na URL", http.StatusBadRequest)
		return
	}

	vendedorID, err := strconv.Atoi(vendedorIDStr)
	if err != nil {
		http.Error(w, "ID do vendedor inválido", http.StatusBadRequest)
		return
	}

	err = adapters.ObterVendas(w, vendedorID)
	if err != nil {
		boom.BadImplementation(w, err)
		return
	}
}

func HistoricoGeral(w http.ResponseWriter, r *http.Request)  {
	err:= adapters.ObterHistoricoGeral(w)
	if err != nil {
		boom.BadImplementation(w, err)
	}
}

