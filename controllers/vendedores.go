package controllers

import (
	"encoding/json"
	"net/http"
	"errors"

	"github.com/darahayes/go-boom"
	"github.com/Matari73/APIGo/adapters/vendedores"
	"github.com/Matari73/APIGo/database"
	"github.com/Matari73/APIGo/models"
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

	var vendedor models.Vendedor
	if err := database.DB.First(&vendedor, id).Error; err != nil {
		erro:= errors.New("Vendedor não encontrado")
		boom.NotFound(w, erro)
		return
	}

	if err := codificarEmJson(w, vendedor); err != nil {
		erro:= errors.New("Erro ao codificar o vendedor em JSON")
		boom.BadRequest(w, erro)
		return
	}
}

func CreateVendedor(w http.ResponseWriter, r *http.Request) {
	var novoVendedor models.Vendedor

	if err := json.NewDecoder(r.Body).Decode(&novoVendedor); err != nil {
		erro:= errors.New("Erro ao ler o corpo da requisição")
		boom.BadRequest(w, erro)
		return
	}
	
	if err:= validators.ValidateNome(novoVendedor); err != nil {
		boom.BadRequest(w, err)
		return
	}

	if err := database.DB.Create(&novoVendedor).Error; err != nil {
		erro:= errors.New("Erro ao criar o novo vendedor")
		boom.BadImplementation(w, erro)
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

	var vendedor models.Vendedor
	result := database.DB.Delete(&vendedor, id)
	if result.Error != nil {
		erro:= errors.New("Erro ao excluir o vendedor")
		boom.BadImplementation(w, erro)
		return
	}

	if result.RowsAffected == 0 {
		erro:= errors.New("Vendedor não encontrado com este ID")
		boom.NotFound(w, erro)
		return
	}

	w.WriteHeader(http.StatusOK)
	sucesso:= CreateResposta("Vendedor excluído com sucesso!")
	json.NewEncoder(w).Encode(sucesso)
}

func UpdateVendedor(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var vendedor models.Vendedor
	err := json.NewDecoder(r.Body).Decode(&vendedor)
	if err != nil {
		erro:= errors.New("Erro ao decodificar o corpo da requisição")
		boom.BadRequest(w, erro)
		return
	}

	if err:= validators.ValidateNome(vendedor); err != nil {
		boom.BadRequest(w, err)
		return
	}

	result := database.DB.Model(&models.Vendedor{}).Where("vendedor_id = ?", id).Updates(&vendedor)
	if result.Error != nil {
		erro:= errors.New("Erro ao atualizar o vendedor no banco de dados")
		boom.BadImplementation(w, erro)
		return
	}

	if result.RowsAffected == 0 {
		erro:= errors.New("Vendedor não encontrado")
		boom.NotFound(w, erro)
		return
	}

	if err := codificarEmJson(w, vendedor); err != nil {
		erro:= errors.New("Erro ao codificar o vendedor em JSON")
		boom.BadRequest(w, erro)
		return
	}
}
