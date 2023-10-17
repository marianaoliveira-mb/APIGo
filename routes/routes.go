package routes

import (
	"log"
	"net/http"

	"github.com/Matari73/APIGo/controllers"
	"github.com/gorilla/mux"
)

func HandleResquest() {
	r := mux.NewRouter()
	http.HandleFunc("/", controllers.Home)

	//rotas GetAll
	r.HandleFunc("/api/produtos", controllers.GetProdutos).Methods("Get")
	r.HandleFunc("/api/clientes", controllers.GetClientes).Methods("Get")
	r.HandleFunc("/api/vendedores", controllers.GetVendedores).Methods("Get")
	r.HandleFunc("/api/pedidos", controllers.GetPedidos).Methods("Get")

	//rotas GetById
	r.HandleFunc("/api/produtos/{id}", controllers.GetProduto).Methods("Get")
	r.HandleFunc("/api/clientes/{id}", controllers.GetCliente).Methods("Get")
	r.HandleFunc("/api/vendedores/{id}", controllers.GetVendedor).Methods("Get")
	r.HandleFunc("/api/pedidos/{id}", controllers.GetPedido).Methods("Get")

	//rotas Create
	r.HandleFunc("/api/produtos", controllers.CreateProduto).Methods("Post")
	r.HandleFunc("/api/clientes", controllers.CreateCliente).Methods("Post")
	r.HandleFunc("/api/vendedores", controllers.CreateVendedor).Methods("Post")
	r.HandleFunc("/api/pedidos", controllers.CreatePedido).Methods("Post")

	//rotas Delete
	r.HandleFunc("/api/produtos/{id}", controllers.DeleteProduto).Methods("Delete")
	r.HandleFunc("/api/clientes/{id}", controllers.DeleteCliente).Methods("Delete")
	r.HandleFunc("/api/vendedores/{id}", controllers.DeleteVendedor).Methods("Delete")
	r.HandleFunc("/api/pedidos/{id}", controllers.DeletePedido).Methods("Delete")

	//rotas Update
	r.HandleFunc("/api/produtos/{id}", controllers.UpdateProduto).Methods("Put")
	r.HandleFunc("/api/clientes/{id}", controllers.UpdateCliente).Methods("Put")
	r.HandleFunc("/api/vendedores/{id}", controllers.UpdateVendedor).Methods("Put")
	r.HandleFunc("/api/pedidos/{id}", controllers.UpdatePedido).Methods("Put")

	//Rota carrinho de produto
	r.HandleFunc("/api/adicionar-produto-pedido", controllers.AdicionarProdutoAoPedidoHandler).Methods("POST")
	r.HandleFunc("/api/clientes/{cliente_id}/historico_compras", controllers.HistoricoCompras).Methods("GET")

	// r.HandleFunc("/api/produto-pedido", controllers.CreateProdutoPedidoHandler).Methods("Post")
	// r.HandleFunc("/api/produto-pedido", controllers.HandleAssociacoesProdutoPedido).Methods("Get")
	log.Fatal(http.ListenAndServe(":8000", r))

}
