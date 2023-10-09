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
	r.HandleFunc("/api/prodped", controllers.GetProdutosPedidos).Methods("Get")

	//rotas GetById
	r.HandleFunc("/api/produtos/{id}", controllers.GetProduto).Methods("Get")
	r.HandleFunc("/api/clientes/{id}", controllers.GetCliente).Methods("Get")
	r.HandleFunc("/api/vendedores/{id}", controllers.GetVendedor).Methods("Get")
	r.HandleFunc("/api/pedidos/{id}", controllers.GetPedido).Methods("Get")
	//r.HandleFunc("/api/prodped/{id}", controllers.GetProdutoPedido)

	//rotas Create
	r.HandleFunc("/api/produtos", controllers.CreateProduto).Methods("Post")
	r.HandleFunc("/api/clientes", controllers.CreateCliente).Methods("Post")
	r.HandleFunc("/api/vendedores", controllers.CreateVendedor).Methods("Post")
	r.HandleFunc("/api/pedidos", controllers.CreatePedido).Methods("Post")
	//r.HandleFunc("/api/prodped", controllers.GetProdutosPedidos).Methods("Get")

	log.Fatal(http.ListenAndServe(":8000", r))

}
