package adapters

import(
	"errors"

	"github.com/Matari73/APIGo/models"
	"github.com/Matari73/APIGo/database"
)

func AdicionarProdutoAoPedido(pedidoID uint, produtoID uint, quantidade int) error {
	pedido := models.Pedido{}
	if err := database.DB.First(&pedido, pedidoID).Error; err != nil {
		return errors.New("Erro ao carregar o pedido")
	}

	produto := models.Produto{}
	if err := database.DB.First(&produto, produtoID).Error; err != nil {
		return errors.New("Erro ao carregar o produto")

	}

	if quantidade > produto.Estoque {
		return errors.New("Quantidade maior do que o estoque disponível")
	}

	novoEstoque := produto.Estoque - quantidade
	if err := database.DB.Model(&models.Produto{}).Where("produto_id = ?", produtoID).
		Update("estoque", novoEstoque).Error; err != nil {
			return errors.New("Erro ao atualizar o estoque do produto")
	}

	produtoPedido := models.ProdutoPedido{}
	resultProdutoPedido := database.DB.Where("produto_id = ? AND pedido_id = ?", produtoID, pedidoID).First(&produtoPedido)
	if resultProdutoPedido.Error == nil {
		errors.New("A associação entre produto e pedido já existe.")
		return nil
	}

	produtoPedido = models.ProdutoPedido{
		PedidoID:   int(pedidoID),
		ProdutoID:  int(produtoID),
		Quantidade: quantidade,
	}

	resultCriacao := database.DB.Create(&produtoPedido)
	if resultCriacao.Error != nil {
		return errors.New("Erro ao criar a associação entre produto e pedido")
	}

	return nil
}