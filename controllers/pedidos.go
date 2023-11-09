package controllers

import (
	"encoding/json"
	"net/http"
	"strings"
	"time"
	"errors"

	"github.com/darahayes/go-boom"
	"github.com/Matari73/APIGo/database"
	"github.com/Matari73/APIGo/models"
	"github.com/gorilla/mux"
)

func GetPedidos(w http.ResponseWriter, r *http.Request) {
	var p []models.Pedido
	if err := database.DB.Find(&p).Error; err != nil {
		erro:= errors.New("Erro ao buscar Pedidos")
		boom.BadImplementation(w, erro)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(p); err != nil {
		erro:= errors.New("Erro ao codificar a resposta")
		boom.BadImplementation(w, erro)
		return
	}
}
func GetPedido(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var pedido models.Pedido
	if err := database.DB.First(&pedido, id).Error; err != nil {
		erro:= errors.New("Pedido não encontrado")
		boom.NotFound(w, erro)
		return
	}

	if err := codificarEmJson(w, pedido); err != nil {
		erro:= errors.New("Erro ao codificar o pedido em JSON")
		boom.BadRequest(w, erro)
		return
	}
}

func CreatePedido(w http.ResponseWriter, r *http.Request) {
	var novoPedido models.Pedido

	if err := json.NewDecoder(r.Body).Decode(&novoPedido); err != nil {
		erro:= errors.New("Erro ao ler o corpo da requisição")
		boom.BadRequest(w, erro)
		return
	}

	novoPedido.DataPedido = time.Now()
	saldoCliente, err := obterSaldoCliente(uint(novoPedido.ClienteID))
	if err != nil {
		erro:= errors.New("Erro ao obter o saldo do cliente")
		boom.BadImplementation(w, erro)
		return
	}

	clienteID := uint(novoPedido.ClienteID)
	existeCliente, err := verificaClienteExistente(clienteID)
	if err != nil {
		erro:= errors.New("Erro ao verificar o cliente")
		boom.BadImplementation(w, erro)
		return
	}

	if !existeCliente {
		erro:= errors.New("ID do cliente inválido")
		boom.BadRequest(w, erro)
		return
	}

	vendedorID := uint(novoPedido.VendedorID)
	existeVendedor, err := verificaVendedorExistente(vendedorID)
	if err != nil {
		erro:= errors.New("Erro ao verificar o Vendedor")
		boom.BadImplementation(w, erro)
		return
	}

	if !existeVendedor {
		erro:= errors.New("ID do vendedor inválido")
		boom.BadImplementation(w, erro)
		return
	}

	if strings.ToUpper(novoPedido.StatusPedido) != "EM ANDAMENTO" &&
		strings.ToUpper(novoPedido.StatusPedido) != "ENVIADO" &&
		strings.ToUpper(novoPedido.StatusPedido) != "CONCLUIDO" {
		erro:= errors.New("Status inválido")
		boom.BadRequest(w, erro)
		return
	}

	if saldoCliente < novoPedido.ValorPedido {
		erro:= errors.New("Saldo do cliente insuficiente para o pedido")
		boom.BadRequest(w, erro)
		return
	}

	if err := database.DB.Create(&novoPedido).Error; err != nil {
		erro:= errors.New("Erro ao criar o novo pedido")
		boom.BadImplementation(w, erro)
		return
	}

	if err := atualizarSaldoCliente(uint(novoPedido.ClienteID), novoPedido.ValorPedido); err != nil {
		erro:= errors.New("Erro ao atualizar o saldo do cliente")
		boom.BadImplementation(w, erro)
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

	var pedido models.Pedido
	result := database.DB.Delete(&pedido, id)
	if result.Error != nil {
		erro:= errors.New("Erro ao excluir o pedido")
		boom.BadImplementation(w, erro)
		return
	}

	if result.RowsAffected == 0 {
		erro:= errors.New("Pedido não encontrado com este ID")
		boom.NotFound(w, erro)
		return
	}

	w.WriteHeader(http.StatusOK)
	sucesso:= CreateResposta("Pedido excluído com sucesso!")
	json.NewEncoder(w).Encode(sucesso)
}

func UpdatePedido(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var pedido models.Pedido
	err := json.NewDecoder(r.Body).Decode(&pedido)
	if err != nil {
		erro:= errors.New("Erro ao decodificar o corpo da requisição")
		boom.BadRequest(w, erro)
		return
	}

	result := database.DB.Model(&models.Pedido{}).Where("pedido_id = ?", id).Updates(&pedido)
	if result.Error != nil {
		erro:= errors.New("Erro ao atualizar o pedido no banco de dados")
		boom.BadImplementation(w, erro)
		return
	}

	if result.RowsAffected == 0 {
		erro:= errors.New("Pedido não encontrado")
		boom.NotFound(w, erro)
		return
	}

	if err := codificarEmJson(w, pedido); err != nil {
		erro:= errors.New("Erro ao codificar o pedido em JSON")
		boom.BadImplementation(w, erro)
		return
	}
}
