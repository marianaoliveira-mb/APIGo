package models

type ProdutoPedido struct {
	ProdutoID  int `json:"produto_id"`
	PedidoID   int `json:"pedido_id"`
	Quantidade int `json:"quantidade"`
}

var ProdutosPedidos []ProdutoPedido
