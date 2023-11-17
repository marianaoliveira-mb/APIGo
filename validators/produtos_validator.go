package validators

import(
	"errors"

	"github.com/Matari73/APIGo/models"
)

func ValidateProduto(produto models.Produto) error  {
	if produto.NomeProduto == "" {
		return errors.New("O nome não pode ser vazio")
	}

	if produto.Estoque <= 0 {
		return errors.New("A quantidade do estoque deve ser maior que 0")
	}

	if produto.ValorProduto <= 0 {
		return errors.New("O valor do produto deve ser maior que 0")
	}

	return nil
}