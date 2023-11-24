package adapters 

import(
	"errors"

	"github.com/Matari73/APIGo/database"
	"github.com/Matari73/APIGo/models"
)

func AtualizarPedido(pedido models.Pedido, id string) error{
	result := database.DB.Model(&models.Pedido{}).Where("pedido_id = ?", id).Updates(&pedido)

	if result.Error != nil{
		erro:= errors.New("Erro ao atualizar o cliente no banco de dados")		
		return erro
	}
	return nil
}