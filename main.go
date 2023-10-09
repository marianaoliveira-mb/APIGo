package main

import (
	"fmt"
	"time"

	"github.com/Matari73/APIGo/database"
	"github.com/Matari73/APIGo/models"
	"github.com/Matari73/APIGo/routes"
)

func parseDataHora(layout, value string) (time.Time, error) {
	return time.Parse(layout, value)
}

func main() {
	models.Produtos = []models.Produto{
		{ProdutoID: 1, NomeProduto: "cinto", ValorProduto: 15, Estoque: 10},
		{ProdutoID: 2, NomeProduto: "Bolsa", ValorProduto: 30, Estoque: 20},
	}

	models.Clientes = []models.Cliente{
		{ClienteID: 1, NomeCliente: "Mariana", TelefoneCliente: "1", Saldo: 1000},
		{ClienteID: 2, NomeCliente: "Sergio", TelefoneCliente: "2", Saldo: 100},
	}

	models.Vendedores = []models.Vendedor{
		{VendedorID: 1, NomeVendedor: "Rita"},
		{VendedorID: 2, NomeVendedor: "Tais"},
	}

	dataPedido1, _ := parseDataHora("02/Jan/2003 - 15:04:05", "16/Jul/2023 - 08:30:45")
	dataPedido2, _ := parseDataHora("02/Jan/2003 - 15:04:05", "31/Jul/2023 - 22:30:45")
	models.Pedidos = []models.Pedido{
		{PedidoID: 1, DataPedido: dataPedido1, StatusPedido: "concluido", ValorPedido: 30, Quantidade: 2, ClienteID: 2, VendedorID: 1},
		{PedidoID: 2, DataPedido: dataPedido2, StatusPedido: "Em andamento", ValorPedido: 90, Quantidade: 3, ClienteID: 1, VendedorID: 2},
	}

	models.ProdutosPedidos = []models.ProdutoPedido{
		{ProdutoID: 1, PedidoID: 2},
		{ProdutoID: 2, PedidoID: 1},
	}

	database.ConectaComBancoDeDados()
	fmt.Println("Iniciando o servidor com Go")
	routes.HandleResquest()
}
