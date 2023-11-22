package adapters 

import(
	"encoding/json"
	"errors"
	"net/http"

	"github.com/Matari73/APIGo/database"
	"github.com/Matari73/APIGo/models"
)

func LerCorpoRequisicao(r *http.Request) (models.Vendedor, error) {
	var novoVendedor models.Vendedor
	if err:= json.NewDecoder(r.Body).Decode(&novoVendedor); err != nil{
		erro:= errors.New("Erro ao ler o corpo da requisição")
		return novoVendedor, erro
	}
	
	return novoVendedor, nil
}

func CriarVendedor(novoVendedor models.Vendedor) (models.Vendedor, error) {
	if  err := database.DB.Create(&novoVendedor).Error; err != nil {
		erro:= errors.New("Erro ao criar o novo Vendedor")
		return novoVendedor, erro
	}

	return novoVendedor, nil
}