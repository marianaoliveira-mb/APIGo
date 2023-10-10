package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Matari73/APIGo/database"
	"github.com/Matari73/APIGo/models"
	"github.com/gorilla/mux"
)

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Home Page")
}

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
	json.NewDecoder(r.Body).Decode(&produto)
	database.DB.Model(&models.Produto{}).Where("produto_id = ?", id).Updates(&produto)
	json.NewEncoder(w).Encode(produto)
}

//clientes
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
	database.DB.First(&cliente, id)
	json.NewDecoder(r.Body).Decode(&cliente)
	database.DB.Model(&models.Cliente{}).Where("cliente_id = ?", id).Updates(&cliente)
	json.NewEncoder(w).Encode(cliente)
}

//vendedores
func GetVendedores(w http.ResponseWriter, r *http.Request) {
	var v []models.Vendedor
	if err := database.DB.Find(&v).Error; err != nil {
		http.Error(w, "Erro ao buscar clientes: "+err.Error(), http.StatusInternalServerError)
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
	database.DB.First(&vendedor, id)
	json.NewDecoder(r.Body).Decode(&vendedor)
	database.DB.Model(&models.Vendedor{}).Where("vendedor_id = ?", id).Updates(&vendedor)
	json.NewEncoder(w).Encode(vendedor)
}

//pedidos
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

	if err := database.DB.Create(&novoPedido).Error; err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Erro ao criar o novo cliente: %v", err)
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
	database.DB.First(&pedido, id)
	json.NewDecoder(r.Body).Decode(&pedido)
	database.DB.Model(&models.Pedido{}).Where("pedido_id = ?", id).Updates(&pedido)
	json.NewEncoder(w).Encode(pedido)
}

//produtosPedidos
func GetProdutosPedidos(w http.ResponseWriter, r *http.Request) {
	var p []models.Pedido
	database.DB.Find(&p)
	json.NewEncoder(w).Encode(p)
}

/*func GetProdutoPedido(w http.ResponseWriter, r *http.Request) { //ver erro
	vars := mux.Vars(r)
	id := vars["id"]

	//log.Printf("ID do Produto/Pedido  recebido: %s", id)
	for _, ProdutoPedido := range models.ProdutosPedidos {
		//log.Printf("Produto e Pedido encontrado: %+v", ProdPed)
		if strconv.Itoa(ProdutoPedido.ClienteID) == id {
			json.NewEncoder(w).Encode(ProdutoPedido)
		}
	}
}*/
