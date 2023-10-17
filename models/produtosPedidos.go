package models

type ProdutoPedido struct {
	PedidoID   int `json:"pedido_id" gorm:"primaryKey;foreignKey:PedidoID"`
	ProdutoID  int `json:"produto_id" gorm:"primaryKey;foreignKey:ProdutoID"`
	Quantidade int `json:"quantidade"`
}

var ProdutosPedidos []ProdutoPedido
