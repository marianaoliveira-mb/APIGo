package adapters 

import(
	"errors"

	"github.com/Matari73/APIGo/database"
	"github.com/Matari73/APIGo/models"
)

func AtualizarProduto(produto models.Produto, id string) error {
	result := database.DB.Model(&models.Produto{}).Where("produto_id = ?", id).Updates(&produto)

	if result.Error != nil {
		erro:= errors.New("Erro ao atualizar o produto no banco de dados")
		return erro
	}

	return nil
}