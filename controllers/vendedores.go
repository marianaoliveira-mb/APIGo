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

func GetVendedores(w http.ResponseWriter, r *http.Request) {
	var v []models.Vendedor
	if err := database.DB.Find(&v).Error; err != nil {
		http.Error(w, "Erro ao buscar vendedores: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(v); err != nil {
		http.Error(w, "Erro ao codificar a resposta: "+err.Error(), http.StatusInternalServerError)
		return
	}
}

func GetVendedor(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var vendedor models.Vendedor
	if err := database.DB.First(&vendedor, id).Error; err != nil {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "Vendedor não encontrado")
		return
	}

	err := json.NewEncoder(w).Encode(vendedor)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Erro ao codificar vendedor em JSON: %v", err)
		return
	}
}

func CreateVendedor(w http.ResponseWriter, r *http.Request) {
	var novoVendedor models.Vendedor

	if err := json.NewDecoder(r.Body).Decode(&novoVendedor); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Erro ao ler o corpo da requisição: %v", err)
		return
	}
	if novoVendedor.NomeVendedor == "" {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "O nome não deve ser vazio")
		return
	}

	if err := database.DB.Create(&novoVendedor).Error; err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Erro ao criar o novo cliente: %v", err)
		return
	}

	if err := json.NewEncoder(w).Encode(novoVendedor); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Erro ao codificar o produto em JSON: %v", err)
		return
	}
}

func DeleteVendedor(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var vendedor models.Vendedor
	result := database.DB.Delete(&vendedor, id)
	if result.Error != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Erro ao excluir o vendedor: %v", result.Error)
		return
	}

	if result.RowsAffected == 0 {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "Vendedor não encontrado com o ID: %s", id)
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "vendedor excluído com sucesso")
}

func UpdateVendedor(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var vendedor models.Vendedor
	err := json.NewDecoder(r.Body).Decode(&vendedor)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Erro ao decodificar o corpo da requisição: %v", err)
		return
	}

	if vendedor.VendedorID == strconv.Itoa(0) {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "ID do vendedor não fornecido ou inválido")
		return
	}

	result := database.DB.Model(&models.Vendedor{}).Where("vendedor_id = ?", id).Updates(&vendedor)
	if result.Error != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Erro ao atualizar o vendedor no banco de dados: %v", result.Error)
		return
	}

	if result.RowsAffected == 0 {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprint(w, "vendedor não encontrado")
		return
	}

	err = json.NewEncoder(w).Encode(vendedor)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Erro ao codificar o vendedor em JSON: %v", err)
		return
	}
}
