package adapters 

import(
	"encoding/json"
	"errors"
	"net/http"

	"github.com/Matari73/APIGo/database"
	"github.com/Matari73/APIGo/models"
)

func LerCorpoRequisicao(r *http.Request) (models.Cliente, error) {
	var novoCliente models.Cliente
	if err:= json.NewDecoder(r.Body).Decode(&novoCliente); err != nil{
		erro:= errors.New("Erro ao ler o corpo da requisição")
		return novoCliente, erro
	}
	
	return novoCliente, nil
}

func CriarCliente(novoCliente models.Cliente) (models.Cliente, error) {
	if  err := database.DB.Create(&novoCliente).Error; err != nil {
		erro:= errors.New("Erro ao criar o novo cliente")
		return novoCliente, erro
	}

	return novoCliente, nil
}