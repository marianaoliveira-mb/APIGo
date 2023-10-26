package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/Matari73/APIGo/database"
	"github.com/Matari73/APIGo/models"
	"github.com/gorilla/mux"
)

func GetClientes(w http.ResponseWriter, r *http.Request) {
	var c []models.Cliente
	if err := database.DB.Find(&c).Error; err != nil {
		http.Error(w, "Erro ao buscar clientes: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(c); err != nil {
		http.Error(w, "Erro ao codificar a resposta: "+err.Error(), http.StatusInternalServerError)
		return
	}
}

func GetCliente(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	var cliente models.Cliente

	if err := database.DB.First(&cliente, id).Error; err != nil {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "Cliente não encontrado")
		return
	}

	err := json.NewEncoder(w).Encode(cliente)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Erro ao codificar cliente em JSON: %v", err)
		return
	}
}

func CreateCliente(w http.ResponseWriter, r *http.Request) {
	var novoCliente models.Cliente

	if err := json.NewDecoder(r.Body).Decode(&novoCliente); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Erro ao ler o corpo da requisição: %v", err)
		return
	}

	if novoCliente.NomeCliente == "" {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "O nome não deve ser vazio")
		return
	}

	if len(novoCliente.TelefoneCliente) < 10 || len(novoCliente.TelefoneCliente) > 12 {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "O telefone deve ter entre 10 e 12 caracteres")
		return
	}

	if err := database.DB.Create(&novoCliente).Error; err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Erro ao criar o novo cliente: %v", err)
		return
	}

	if err := json.NewEncoder(w).Encode(novoCliente); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Erro ao codificar o produto em JSON: %v", err)
		return
	}
}

func DeleteCliente(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var cliente models.Cliente
	result := database.DB.Delete(&cliente, id)
	if result.Error != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Erro ao excluir o cliente: %v", result.Error)
		return
	}

	if result.RowsAffected == 0 {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "Cliente não encontrado com o ID: %s", id)
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Cliente excluído com sucesso")
}

func UpdateCliente(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var cliente models.Cliente
	err := json.NewDecoder(r.Body).Decode(&cliente)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Erro ao decodificar o corpo da requisição: %v", err)
		return
	}

	if cliente.ClienteID == strconv.Itoa(0) {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "ID do cliente não fornecido ou inválido")
		return
	}

	result := database.DB.Model(&models.Cliente{}).Where("cliente_id = ?", id).Updates(&cliente)
	if result.Error != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Erro ao atualizar o cliente no banco de dados: %v", result.Error)
		return
	}

	if result.RowsAffected == 0 {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprint(w, "Cliente não encontrado")
		return
	}

	err = json.NewEncoder(w).Encode(cliente)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Erro ao codificar cliente em JSON: %v", err)
		return
	}
}
