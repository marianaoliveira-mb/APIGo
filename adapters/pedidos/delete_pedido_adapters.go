package adapters

import(
	"errors"

	"github.com/Matari73/APIGo/database"
	"github.com/Matari73/APIGo/models"
)


func DeletarPedido(id string) error{
	var pedido models.Pedido
	result:= database.DB.Delete(&pedido, id)

	if result.Error != nil{
		erro:= errors.New("Erro ao excluir o pedido")
		return erro
	}

	return nil
}