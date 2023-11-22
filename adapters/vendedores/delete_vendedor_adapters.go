package adapters

import(
	"errors"

	"github.com/Matari73/APIGo/database"
	"github.com/Matari73/APIGo/models"
)

func DeletarVendedor(id string) error{
	var vendedor models.Vendedor
	result:= database.DB.Delete(&vendedor, id)

	if result.Error != nil{
		erro:= errors.New("Erro ao excluir o vendedor")
		return erro
	}

	return nil
}
