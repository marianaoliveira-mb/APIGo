package controllers

import (
	"encoding/json"
	"net/http"
	"errors"

	"github.com/darahayes/go-boom"
	"github.com/Matari73/APIGo/database"
	"github.com/Matari73/APIGo/models"
	"github.com/Matari73/APIGo/validators"
	"github.com/Matari73/APIGo/adapters/clientes"
	"github.com/gorilla/mux"
)

func GetClientes(w http.ResponseWriter, r *http.Request) {
	clientes , err := adapters.BuscarClientes()
	if err != nil {
		boom.BadImplementation(w, err)
		return
	}

	if err := adapters.CodificarResposta(w, clientes) ; err != nil {
		boom.BadImplementation(w, err)
		return
	}
}

func GetCliente(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	c, err := adapters.BuscarClienteById(id)
	if err != nil {
		boom.NotFound(w, err)
		return
	}

	if err := codificarEmJson(w, c); err != nil {
		erro:= errors.New("Erro ao codificar o cliente em JSON")
		boom.BadRequest(w, erro)
		return
	} 
}

func CreateCliente(w http.ResponseWriter, r *http.Request) {
	novoCliente, err := adapters.LerCorpoRequisicao(r)
	if err != nil {
		boom.BadRequest(w, err)
		return
	}

	if err:= validators.ValidateTelNome(novoCliente); err != nil {
		boom.BadRequest(w, err)
		return
	}

	novoCliente, err = adapters.CriarCliente(novoCliente)
	if err != nil {
		boom.BadImplementation(w, err)
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

	// result := adapters.DeletarCliente(id)
	// fmt.Println(result)
	// if result.Error != nil {
	// 	boom.BadImplementation(w, result.Error)
	// 	return
	// }

	// if result.RowsAffected == 0 {
	// 	erro:= errors.New("Cliente não encontrado com este ID")
	// 	boom.NotFound(w, erro)
	// 	return
	// }
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
	
	if err:= validators.ValidateTelNome(cliente); err != nil {
		boom.BadRequest(w, err)
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
