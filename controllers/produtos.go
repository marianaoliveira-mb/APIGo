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

//produtos
func GetProdutos(w http.ResponseWriter, r *http.Request) {
	var p []models.Produto
	if err := database.DB.Find(&p).Error; err != nil {
		erro:= errors.New("Erro ao buscar produtos")
		boom.BadImplementation(w, erro)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(p); err != nil {
		erro:= errors.New("Erro ao codificar a resposta")
		boom.BadImplementation(w, erro)
		return
	}

	w.WriteHeader(http.StatusOK)
	sucesso:= CreateResposta("Produtos listados com sucesso!")
	json.NewEncoder(w).Encode(sucesso)
}

func GetProduto(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	var produto models.Produto

	if err := database.DB.First(&produto, id).Error; err != nil {
		erro:= errors.New("Produto não encontrado")
		boom.NotFound(w, erro)
		return
	}

	if err := codificarEmJson(w, produto); err != nil {
		erro:= errors.New("Erro ao codificar o produto em JSON")
		boom.BadRequest(w, erro)
		return
	}
	w.WriteHeader(http.StatusOK)
	sucesso:= CreateResposta("Produto listado com sucesso!")
	json.NewEncoder(w).Encode(sucesso)
}

func CreateProduto(w http.ResponseWriter, r *http.Request) {
	var novoProduto models.Produto
	if err := json.NewDecoder(r.Body).Decode(&novoProduto); err != nil {
		erro:= errors.New("Erro ao ler o corpo da requisição")
		boom.BadRequest(w, erro)
		return
	}

	if novoProduto.NomeProduto == "" {
		erro:= errors.New("O nome não pode ser vazio")
		boom.BadRequest(w, erro)
		return
	}

	if novoProduto.Estoque <= 0 {
		erro:= errors.New("A quantidade do estoque deve ser maior que 0")
		boom.BadRequest(w, erro)
		return
	}

	if novoProduto.ValorProduto <= 0 {
		erro:= errors.New("O valor do produto deve ser maior que 0")
		boom.BadRequest(w, erro)
		return
	}

	// Verificar se o nome do produto já existe
	var produtoExistente models.Produto
	if err := database.DB.Where("nome_produto = ?", novoProduto.NomeProduto).First(&produtoExistente).Error; err == nil {
		erro:= errors.New("Já existe um produto com este nome")
		boom.BadRequest(w, erro)
		return
	}

	if err := database.DB.Create(&novoProduto).Error; err != nil {
		erro:= errors.New("Erro ao criar o novo produto")
		boom.BadImplementation(w, erro)
		return
	}

	if err := codificarEmJson(w, novoProduto); err != nil {
		erro:= errors.New("Erro ao codificar o cliente em JSON")
		boom.BadRequest(w, erro)
		return
	}

	w.WriteHeader(http.StatusCreated)
	sucesso:= CreateResposta("Produto criado com sucesso!")
	json.NewEncoder(w).Encode(sucesso)
}

func DeleteProduto(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var produto models.Produto
	result := database.DB.Delete(&produto, id)

	if result.Error != nil {
		erro:= errors.New("Erro ao excluir o produto")
		boom.BadImplementation(w, erro)
		return
	}

	if result.RowsAffected == 0 {
		erro:= errors.New("Produto não encontrado com este ID")
		boom.NotFound(w, erro)
		return
	}

	w.WriteHeader(http.StatusOK)
	sucesso:= CreateResposta("Produto excluído com sucesso!")
	json.NewEncoder(w).Encode(sucesso)
}

func UpdateProduto(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var produto models.Produto
	err := json.NewDecoder(r.Body).Decode(&produto)
	if err != nil {
		erro:= errors.New("Erro ao decodificar o corpo da requisição")
		boom.BadRequest(w, erro)
		return
	}

	if produto.NomeProduto == "" {
		erro:= errors.New("O nome não deve ser vazio")
		boom.BadRequest(w, erro)
		return
	}

	if produto.Estoque < 0 {
		erro:= errors.New("A quantidade do estoque deve ser maior ou igual a 0")
		boom.BadRequest(w, erro)
		return
	}

	if produto.ValorProduto <= 0 {
		erro:= errors.New("O valor do produto deve ser maior que 0")
		boom.BadRequest(w, erro)
		return
	}

	result := database.DB.Model(&models.Produto{}).Where("produto_id = ?", id).Updates(&produto)
	if result.Error != nil {
		erro:= errors.New("Erro ao atualizar o produto no banco de dados")
		boom.BadImplementation(w, erro)
		return
	}

	if result.RowsAffected == 0 {
		erro:= errors.New("Produto não encontrado")
		boom.NotFound(w, erro)
		return
	}

	if err := codificarEmJson(w, produto); err != nil {
		erro:= errors.New("Erro ao codificar o produto em JSON")
		boom.BadRequest(w, erro)
		return
	}

	w.WriteHeader(http.StatusOK)
	sucesso:= CreateResposta("Produto atualizado com sucesso!")
	json.NewEncoder(w).Encode(sucesso)
}
