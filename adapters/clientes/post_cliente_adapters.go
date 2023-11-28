package adapters 

import(
	"errors"

	"github.com/Matari73/APIGo/database"
	"github.com/Matari73/APIGo/models"
)

func CriarCliente(novoCliente models.Cliente) (models.Cliente, error) {
	if  err := database.DB.Create(&novoCliente).Error; err != nil {
		erro:= errors.New("Erro ao criar o novo cliente")
		return novoCliente, erro
	}

	return novoCliente, nil
}