package adapters

import(
	"errors"

	"github.com/Matari73/APIGo/database"
	"github.com/Matari73/APIGo/models"
)

func DeletarCliente(id string) error{
	var cliente models.Cliente
	if err:= database.DB.Delete(&cliente, id).Error; err != nil{
		erro:= errors.New("Erro ao excluir o cliente")
		return erro
	}

	return nil
}