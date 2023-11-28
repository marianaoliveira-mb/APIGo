package adapters 

import(
	"errors"

	"github.com/Matari73/APIGo/database"
	"github.com/Matari73/APIGo/models"
)

func CriarVendedor(novoVendedor models.Vendedor) (models.Vendedor, error) {
	if  err := database.DB.Create(&novoVendedor).Error; err != nil {
		erro:= errors.New("Erro ao criar o novo Vendedor")
		return novoVendedor, erro
	}

	return novoVendedor, nil
}