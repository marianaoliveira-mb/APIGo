package models

type Produto struct {
	//ProdutoID    int     `json:"produto_id"`
	NomeProduto  string  `json:"nome_produto"`
	ValorProduto float64 `json:"valor_produto"`
	Estoque      int     `json:"estoque"`
}

var Produtos []Produto
