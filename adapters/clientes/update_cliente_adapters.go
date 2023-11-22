package adapters 

import(
	"errors"

	"github.com/Matari73/APIGo/database"
	"github.com/Matari73/APIGo/models"
)

func AtualizarCliente(cliente models.Cliente, id string) error{
	result := database.DB.Model(&models.Cliente{}).Where("cliente_id = ?", id).Updates(&cliente)

	if result.Error != nil{
		erro:= errors.New("Erro ao atualizar o cliente no banco de dados")		
		return erro
	}
	return nil
}
