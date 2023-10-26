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

//produtos
func GetProdutos(w http.ResponseWriter, r *http.Request) {
	var p []models.Produto
	if err := database.DB.Find(&p).Error; err != nil {
		http.Error(w, "Erro ao buscar produtos: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(p); err != nil {
		http.Error(w, "Erro ao codificar a resposta: "+err.Error(), http.StatusInternalServerError)
		return
	}
}

func GetProduto(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	var produto models.Produto

	if err := database.DB.First(&produto, id).Error; err != nil {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "Produto não encontrado")
		return
	}

	err := json.NewEncoder(w).Encode(produto)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Erro ao codificar produto em JSON: %v", err)
		return
	}
}

func CreateProduto(w http.ResponseWriter, r *http.Request) {
	var novoProduto models.Produto
	if err := json.NewDecoder(r.Body).Decode(&novoProduto); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Erro ao ler o corpo da requisição: %v", err)
		return
	}

	if novoProduto.NomeProduto == "" {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "O nome não pode ser vazio")
		return
	}

	if novoProduto.Estoque <= 0 {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "A quantidade do estoque deve ser maior que 0")
		return
	}

	if novoProduto.ValorProduto <= 0 {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "O valor do produto deve ser maior que 0")
		return
	}

	// Verificar se o nome do produto já existe
	var produtoExistente models.Produto
	if err := database.DB.Where("nome_produto = ?", novoProduto.NomeProduto).First(&produtoExistente).Error; err == nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Já existe um produto com o nome '%s'", novoProduto.NomeProduto)
		return
	}

	if err := database.DB.Create(&novoProduto).Error; err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Erro ao criar o produto: %v", err)
		return
	}

	if err := json.NewEncoder(w).Encode(novoProduto); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Erro ao codificar o produto em JSON: %v", err)
		return
	}
}

func DeleteProduto(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var produto models.Produto
	result := database.DB.Delete(&produto, id)

	if result.Error != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Erro ao excluir o produto: %v", result.Error)
		return
	}

	if result.RowsAffected == 0 {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "Produto não encontrado com o ID: %s", id)
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Produto excluído com sucesso")
}

func UpdateProduto(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var produto models.Produto
	err := json.NewDecoder(r.Body).Decode(&produto)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Erro ao decodificar o corpo da requisição: %v", err)
		return
	}

	if produto.NomeProduto == "" {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "O nome não pode ser vazio")
		return
	}

	if produto.Estoque < 0 {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "A quantidade do estoque deve ser maior ou igual a 0")
		return
	}

	if produto.ValorProduto <= 0 {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "O valor do produto deve ser maior que 0")
		return
	}

	if produto.ProdutoID == strconv.Itoa(0) {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "ID do produto não fornecido ou inválido")
		return
	}

	result := database.DB.Model(&models.Produto{}).Where("produto_id = ?", id).Updates(&produto)
	if result.Error != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Erro ao atualizar o produto no banco de dados: %v", result.Error)
		return
	}

	if result.RowsAffected == 0 {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprint(w, "Produto não encontrado")
		return
	}

	err = json.NewEncoder(w).Encode(produto)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Erro ao codificar produto em JSON: %v", err)
		return
	}
}
