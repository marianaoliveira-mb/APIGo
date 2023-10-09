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
	database.DB.Find(&p)
	json.NewEncoder(w).Encode(p)
}

func GetProduto(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	var produto models.Produto
	database.DB.First(&produto, id)
	json.NewEncoder(w).Encode(produto)
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
	var c []models.Cliente
	database.DB.Find(&c)
	json.NewEncoder(w).Encode(c)
}

func GetCliente(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	var cliente models.Cliente
	database.DB.First(&cliente, id)
	json.NewEncoder(w).Encode(cliente)
}

func CreateCliente(w http.ResponseWriter, r *http.Request) {
	var novoCliente models.Cliente
	json.NewDecoder(r.Body).Decode(&novoCliente)
	database.DB.Create(&novoCliente)
	json.NewEncoder(w).Encode(novoCliente)
}

//vendedores
func GetVendedores(w http.ResponseWriter, r *http.Request) {
	var v []models.Vendedor
	database.DB.Find(&v)
	json.NewEncoder(w).Encode(v)
}

func GetVendedor(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	var vendedor models.Vendedor
	database.DB.First(&vendedor, id)
	json.NewEncoder(w).Encode(vendedor)
}

func CreateVendedor(w http.ResponseWriter, r *http.Request) {
	var novoVendedor models.Vendedor
	json.NewDecoder(r.Body).Decode(&novoVendedor)
	database.DB.Create(&novoVendedor)
	json.NewEncoder(w).Encode(novoVendedor)
}

//pedidos
func GetPedidos(w http.ResponseWriter, r *http.Request) {
	var p []models.Pedido
	database.DB.Find(&p)
	json.NewEncoder(w).Encode(p)
}
func GetPedido(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	//log.Printf("ID do Produto recebido: %s", id)
	var pedido models.Pedido
	database.DB.First(&pedido, id)
	json.NewEncoder(w).Encode(pedido)
}

func CreatePedido(w http.ResponseWriter, r *http.Request) {
	var novoPedido models.Pedido
	json.NewDecoder(r.Body).Decode(&novoPedido)
	database.DB.Create(&novoPedido)
	json.NewEncoder(w).Encode(novoPedido)
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
