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

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Home Page")
}

//produtos
func GetProdutos(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(models.Produtos)
}

func GetProduto(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	//log.Printf("ID do Produto recebido: %s", id)
	for _, Produto := range models.Produtos {
		//log.Printf("Produto encontrado: %+v", Produto)
		if strconv.Itoa(Produto.ProdutoID) == id {
			json.NewEncoder(w).Encode(Produto)
		}
	}
}

func CreateProduto(w http.ResponseWriter, r *http.Request) {
	var novoProduto models.Produto
	json.NewDecoder(r.Body).Decode(&novoProduto)
	fmt.Println(novoProduto)
	database.DB.Create(&novoProduto)
	json.NewEncoder(w).Encode(novoProduto)
}

//clientes
func GetClientes(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(models.Clientes)
}

func GetCliente(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	//log.Printf("ID do Cliente recebido: %s", id)
	for _, Cliente := range models.Clientes {
		//log.Printf("Cliente encontrado: %+v", Cliente)
		if strconv.Itoa(Cliente.ClienteID) == id {
			json.NewEncoder(w).Encode(Cliente)
		}
	}
}

func CreateCliente(w http.ResponseWriter, r *http.Request) {
	var novoCliente models.Cliente
	json.NewDecoder(r.Body).Decode(&novoCliente)
	database.DB.Create(&novoCliente)
	json.NewEncoder(w).Encode(novoCliente)
}

//vendedores
func GetVendedores(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(models.Vendedores)
}

func GetVendedor(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	//log.Printf("ID do Vendedor recebido: %s", id)
	for _, Vendedor := range models.Vendedores {
		//log.Printf("Vendedor encontrado: %+v", Vendedor)
		if strconv.Itoa(Vendedor.VendedorID) == id {
			json.NewEncoder(w).Encode(Vendedor)
		}
	}
}

//Dando erro
func CreateVendedor(w http.ResponseWriter, r *http.Request) {
	var novoVendedor models.Vendedor
	json.NewDecoder(r.Body).Decode(&novoVendedor)
	database.DB.Create(&novoVendedor)
	json.NewEncoder(w).Encode(novoVendedor)
}

//pedidos
func GetPedidos(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(models.Pedidos)
}
func GetPedido(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	//log.Printf("ID do Cliente recebido: %s", id)
	for _, Pedido := range models.Pedidos {
		//log.Printf("Cliente encontrado: %+v", Pedido)
		if strconv.Itoa(Pedido.PedidoID) == id {
			json.NewEncoder(w).Encode(Pedido)
		}
	}
}

func CreatePedido(w http.ResponseWriter, r *http.Request) {
	var novoPedido models.Pedido
	json.NewDecoder(r.Body).Decode(&novoPedido)
	database.DB.Create(&novoPedido)
	json.NewEncoder(w).Encode(novoPedido)
}

//produtosPedidos
func GetProdutosPedidos(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(models.ProdutosPedidos)
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
