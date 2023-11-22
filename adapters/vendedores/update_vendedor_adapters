package adapters 

import(
	"errors"

	"github.com/Matari73/APIGo/database"
	"github.com/Matari73/APIGo/models"
)

func AtualizarVendedor(vendedor models.Vendedor, id string) error{
	result := database.DB.Model(&models.Vendedor{}).Where("vendedor_id = ?", id).Updates(&vendedor)

	if result.Error != nil{
		erro:= errors.New("Erro ao atualizar o vendedor no banco de dados")		
		return erro
	}
	return nil
}