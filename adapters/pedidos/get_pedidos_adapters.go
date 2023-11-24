package adapters 

import(
	"encoding/json"
	"errors"
	"net/http"

	"github.com/Matari73/APIGo/database"
	"github.com/Matari73/APIGo/models"
)

func BuscarPedidos() ([]models.Pedido, error) {
	var p []models.Pedido
	if err := database.DB.Find(&p).Error; err != nil {
		erro:= errors.New("Erro ao buscar Pedidos")
		return p, erro
	}

	return p, nil
}

func CodificarRespostaPedido(w http.ResponseWriter, p []models.Pedido) error {
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(p); err != nil {
		erro:= errors.New("Erro ao codificar a resposta")
		return erro
	}

	return nil
}

func BuscarPedidoById(id string) (models.Pedido, error) {
	var pedido models.Pedido
	if err := database.DB.First(&pedido, id).Error; err != nil {
		erro:= errors.New("Pedido não encontrado")
		return pedido, erro
	}

	return pedido, nil
}