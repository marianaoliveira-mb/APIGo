package adapters

import(
	"errors"

	"github.com/Matari73/APIGo/database"
	"github.com/Matari73/APIGo/models"
)

func DeletarCliente(id string) error{
	var cliente models.Cliente
	result:= database.DB.Delete(&cliente, id)

	if result.Error != nil{
		erro:= errors.New("Erro ao excluir o cliente")
		return erro
	}

	return nil
}
