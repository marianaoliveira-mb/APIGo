package adapters

import(
	"errors"

	"github.com/Matari73/APIGo/database"
	"github.com/Matari73/APIGo/models"
)


func DeletarProduto(id string) error {
	var produto models.Produto

	result := database.DB.Delete(&produto, id)
	if result.Error != nil{
		erro:= errors.New("Erro ao excluir o produto")
		return erro
	}

	return nil
}