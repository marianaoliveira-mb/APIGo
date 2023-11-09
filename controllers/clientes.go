package controllers

import (
	"encoding/json"
	"net/http"
	"errors"

	"github.com/darahayes/go-boom"
	"github.com/Matari73/APIGo/database"
	"github.com/Matari73/APIGo/models"
	"github.com/gorilla/mux"
)

func GetClientes(w http.ResponseWriter, r *http.Request) {
	var c []models.Cliente
	if err := database.DB.Find(&c).Error; err != nil {
		erro:= errors.New("Erro ao buscar clientes")
		boom.BadImplementation(w, erro)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(c); err != nil {
		erro:= errors.New("Erro ao codificar a resposta")
		boom.BadImplementation(w, erro)
		return
	}
}

func GetCliente(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	var cliente models.Cliente

	if err := database.DB.First(&cliente, id).Error; err != nil {
		erro:= errors.New("Cliente não encontrado")
		boom.NotFound(w, erro)
		return
	}

	if err := codificarEmJson(w, cliente); err != nil {
		erro:= errors.New("Erro ao codificar o cliente em JSON")
		boom.BadRequest(w, erro)
		return
	} 
}

func CreateCliente(w http.ResponseWriter, r *http.Request) {
	var novoCliente models.Cliente

	if err := json.NewDecoder(r.Body).Decode(&novoCliente); err != nil {
		erro:= errors.New("Erro ao ler o corpo da requisição")
		boom.BadRequest(w, erro)
		return
	}

	if novoCliente.NomeCliente == "" {
		erro:= errors.New("O nome não deve ser vazio")
		boom.BadRequest(w, erro)
		return
	}

	if len(novoCliente.TelefoneCliente) < 10 || len(novoCliente.TelefoneCliente) > 12 {
		erro:= errors.New("O telefone deve ter entre 10 e 12 caracteres")
		boom.BadRequest(w, erro)
		return
	}

	if err := database.DB.Create(&novoCliente).Error; err != nil {
		erro:= errors.New("Erro ao criar o novo cliente")
		boom.BadImplementation(w, erro)
		return
	}

	if err := codificarEmJson(w, novoCliente); err != nil {
		erro:= errors.New("Erro ao codificar o cliente em JSON")
		boom.BadImplementation(w, erro)
		return
	}
}

func DeleteCliente(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var cliente models.Cliente
	result := database.DB.Delete(&cliente, id)
	if result.Error != nil {
		erro:= errors.New("Erro ao excluir o cliente")
		boom.BadImplementation(w, erro)
		return
	}

	if result.RowsAffected == 0 {
		erro:= errors.New("Cliente não encontrado com este ID")
		boom.NotFound(w, erro)
		return
	}

	w.WriteHeader(http.StatusOK)
	sucesso:= CreateResposta("Cliente excluído com sucesso!")
	json.NewEncoder(w).Encode(sucesso)
}

func UpdateCliente(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var cliente models.Cliente
	err := json.NewDecoder(r.Body).Decode(&cliente)
	if err != nil {
		erro:= errors.New("Erro ao decodificar o corpo da requisição")
		boom.BadRequest(w, erro)
		return
	}
	
	if cliente.NomeCliente == "" {
		erro:= errors.New("O nome não deve ser vazio")
		boom.BadRequest(w, erro)
		return
	}

	if len(cliente.TelefoneCliente) < 10 || len(cliente.TelefoneCliente) > 12 {
		erro:= errors.New("O telefone deve ter entre 10 e 12 caracteres")
		boom.BadRequest(w, erro)
		return
	}

	result := database.DB.Model(&models.Cliente{}).Where("cliente_id = ?", id).Updates(&cliente)
	if result.Error != nil {
		erro:= errors.New("Erro ao atualizar o cliente no banco de dados")
		boom.BadImplementation(w, erro)
		return
	}

	if result.RowsAffected == 0 {
		erro:= errors.New("Cliente não encontrado")
		boom.NotFound(w, erro)
		return
	}

	if err := codificarEmJson(w, cliente); err != nil {
		erro:= errors.New("Erro ao codificar o cliente em JSON")
		boom.BadRequest(w, erro)
		return
	}
}
