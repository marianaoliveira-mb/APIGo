package adapters 

import(
	"encoding/json"
	"errors"
	"net/http"

	"github.com/Matari73/APIGo/database"
	"github.com/Matari73/APIGo/models"
)


func BuscarClientes() ([]models.Cliente, error) {
	var c []models.Cliente
	if err := database.DB.Find(&c).Error; err != nil {
		erro:= errors.New("Erro ao buscar clientes")
		return c, erro
	}

	return c, nil
}

func CodificarResposta(w http.ResponseWriter, clientes []models.Cliente)  error{
	w.Header().Set("Content-Type", "application/json")
	if err:= json.NewEncoder(w).Encode(clientes); err != nil{
		erro := errors.New("Erro ao codificar a resposta")
		return erro
	}

	return nil
}

func BuscarClienteById(id string) ([]models.Cliente, error) {
	var cliente []models.Cliente
	if err := database.DB.First(&cliente, id).Error; err != nil{
		erro:= errors.New("Cliente não encontrado")
		return cliente, erro
	}

	return cliente, nil
}