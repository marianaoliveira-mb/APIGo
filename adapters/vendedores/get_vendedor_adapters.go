package adapters 

import(
	"encoding/json"
	"errors"
	"net/http"

	"github.com/Matari73/APIGo/database"
	"github.com/Matari73/APIGo/models"
)
func BuscarVendedores()([]models.Vendedor, error){
	var v []models.Vendedor
	if err := database.DB.Find(&v).Error; err != nil {
		erro:= errors.New("Erro ao buscar Vendedores")
		return v, erro
	}

	return v, nil
}

func CodificarRespostaVendedor(w http.ResponseWriter, vendedores []models.Vendedor)  error{
	w.Header().Set("Content-Type", "application/json")
	if err:= json.NewEncoder(w).Encode(vendedores); err != nil{
		erro := errors.New("Erro ao codificar a resposta")
		return erro
	}

	return nil
}
