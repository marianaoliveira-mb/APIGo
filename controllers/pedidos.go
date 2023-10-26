package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/Matari73/APIGo/database"
	"github.com/Matari73/APIGo/models"
	"github.com/gorilla/mux"
)

func GetPedidos(w http.ResponseWriter, r *http.Request) {
	var p []models.Pedido
	if err := database.DB.Find(&p).Error; err != nil {
		http.Error(w, "Erro ao buscar clientes: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(p); err != nil {
		http.Error(w, "Erro ao codificar a resposta: "+err.Error(), http.StatusInternalServerError)
		return
	}
}
func GetPedido(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var pedido models.Pedido
	if err := database.DB.First(&pedido, id).Error; err != nil {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "Pedido não encontrado")
		return
	}

	err := json.NewEncoder(w).Encode(pedido)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Erro ao codificar produto em JSON: %v", err)
		return
	}
}

func CreatePedido(w http.ResponseWriter, r *http.Request) {
	var novoPedido models.Pedido

	if err := json.NewDecoder(r.Body).Decode(&novoPedido); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Erro ao ler o corpo da requisição: %v", err)
		return
	}

	novoPedido.DataPedido = time.Now()

	saldoCliente, err := obterSaldoCliente(uint(novoPedido.ClienteID))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Erro ao obter o saldo do cliente: %v", err)
		return
	}

	clienteID := uint(novoPedido.ClienteID)
	existeCliente, err := verificaClienteExistente(clienteID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Erro ao verificar o cliente: %v", err)
		return
	}

	if !existeCliente {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "ID do cliente inválido")
		return
	}

	vendedorID := uint(novoPedido.VendedorID)
	existeVendedor, err := verificaVendedorExistente(vendedorID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Erro ao verificar o Vendedor: %v", err)
		return
	}

	if !existeVendedor {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "ID do vendedor inválido")
		return
	}

	if strings.ToUpper(novoPedido.StatusPedido) != "EM ANDAMENTO" &&
		strings.ToUpper(novoPedido.StatusPedido) != "ENVIADO" &&
		strings.ToUpper(novoPedido.StatusPedido) != "CONCLUIDO" {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Status inválido")
		return
	}

	if saldoCliente < novoPedido.ValorPedido {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "Saldo do cliente insuficiente para o pedido")
		return
	}

	if err := database.DB.Create(&novoPedido).Error; err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Erro ao criar o novo cliente: %v", err)
		return
	}

	if err := atualizarSaldoCliente(uint(novoPedido.ClienteID), novoPedido.ValorPedido); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Erro ao atualizar o saldo do cliente: %v", err)
		return
	}

	if err := json.NewEncoder(w).Encode(novoPedido); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Erro ao codificar o produto em JSON: %v", err)
		return
	}
}

func DeletePedido(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var pedido models.Pedido
	result := database.DB.Delete(&pedido, id)
	if result.Error != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Erro ao excluir o pedido: %v", result.Error)
		return
	}

	if result.RowsAffected == 0 {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "Pedido não encontrado com o ID: %s", id)
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Pedido excluído com sucesso")
}

func UpdatePedido(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var pedido models.Pedido
	err := json.NewDecoder(r.Body).Decode(&pedido)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Erro ao decodificar o corpo da requisição: %v", err)
		return
	}

	if pedido.PedidoID == strconv.Itoa(0) {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "ID do pedido não fornecido ou inválido")
		return
	}

	result := database.DB.Model(&models.Pedido{}).Where("pedido_id = ?", id).Updates(&pedido)
	if result.Error != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Erro ao atualizar o pedido no banco de dados: %v", result.Error)
		return
	}

	if result.RowsAffected == 0 {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprint(w, "pedido não encontrado")
		return
	}

	err = json.NewEncoder(w).Encode(pedido)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Erro ao codificar o pedido em JSON: %v", err)
		return
	}
}
