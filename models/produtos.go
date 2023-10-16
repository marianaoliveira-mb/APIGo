package models

type Produto struct {
	ProdutoID    string   `json:"produto_id" gorm:"default:uuid_generate_v3()"`
	NomeProduto  string   `json:"nome_produto"`
	ValorProduto float64  `json:"valor_produto"`
	Estoque      int      `json:"estoque"`
	Pedidos      []Pedido `gorm:"many2many:produto_pedidos;foreignKey:ProdutoID;joinForeignKey:ProdutoID;joinReferences:PedidoID"`
}

var Produtos []Produto
