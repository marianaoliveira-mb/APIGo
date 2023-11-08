package routes

import (
	"net/http"

	"github.com/Matari73/APIGo/controllers"
	"github.com/gorilla/mux"
)

func HandleResquest() *mux.Router {
	r := mux.NewRouter()
	http.HandleFunc("/", controllers.Home)

	//rotas produtos
	r.HandleFunc("/api/produtos", controllers.GetProdutos).Methods("Get")
	r.HandleFunc("/api/produtos/{id}", controllers.GetProduto).Methods("Get")
	r.HandleFunc("/api/produtos", controllers.CreateProduto).Methods("Post")
	r.HandleFunc("/api/produtos/{id}", controllers.DeleteProduto).Methods("Delete")
	r.HandleFunc("/api/produtos/{id}", controllers.UpdateProduto).Methods("Put")

	//rotas clientes
	r.HandleFunc("/api/clientes", controllers.GetClientes).Methods("Get")
	r.HandleFunc("/api/clientes/{id}", controllers.GetCliente).Methods("Get")
	r.HandleFunc("/api/clientes", controllers.CreateCliente).Methods("Post")
	r.HandleFunc("/api/clientes/{id}", controllers.DeleteCliente).Methods("Delete")
	r.HandleFunc("/api/clientes/{id}", controllers.UpdateCliente).Methods("Put")

	//rotas vendedores
	r.HandleFunc("/api/vendedores", controllers.GetVendedores).Methods("Get")
	r.HandleFunc("/api/vendedores/{id}", controllers.GetVendedor).Methods("Get")
	r.HandleFunc("/api/vendedores", controllers.CreateVendedor).Methods("Post")
	r.HandleFunc("/api/vendedores/{id}", controllers.DeleteVendedor).Methods("Delete")
	r.HandleFunc("/api/vendedores/{id}", controllers.UpdateVendedor).Methods("Put")

	//rotas pedidos
	r.HandleFunc("/api/pedidos", controllers.GetPedidos).Methods("Get")
	r.HandleFunc("/api/pedidos/{id}", controllers.GetPedido).Methods("Get")
	r.HandleFunc("/api/pedidos", controllers.CreatePedido).Methods("Post")
	r.HandleFunc("/api/pedidos/{id}", controllers.DeletePedido).Methods("Delete")
	r.HandleFunc("/api/pedidos/{id}", controllers.UpdatePedido).Methods("Put")

	//Rota histórico
	r.HandleFunc("/api/adicionar-produto-pedido", controllers.AdicionarProdutoAoPedidoHandler).Methods("POST")
	r.HandleFunc("/api/clientes/{cliente_id}/historico_compras", controllers.HistoricoCompras).Methods("GET")
	r.HandleFunc("/api/vendedores/{vendedor_id}/historico_vendas", controllers.HistoricoVendasVendedor).Methods("GET")
	r.HandleFunc("/api/historico-geral", controllers.HistoricoGeral).Methods("GET")

	return r
}
