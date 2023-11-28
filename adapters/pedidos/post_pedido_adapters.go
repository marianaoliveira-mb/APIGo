package adapters 

import(
	"errors"

	"github.com/Matari73/APIGo/database"
	"github.com/Matari73/APIGo/models"
)

func CriarPedido(novoPedido models.Pedido) (models.Pedido, error) {
	if err := database.DB.Create(&novoPedido).Error; err != nil {
		erro:= errors.New("Erro ao criar o novo pedido")
		return novoPedido, erro
	}

	return novoPedido, nil
}