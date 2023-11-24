package adapters 

import(
	"encoding/json"
	"errors"
	"net/http"

	"github.com/Matari73/APIGo/database"
	"github.com/Matari73/APIGo/models"
)

func LerCorpoRequisicao(r *http.Request) (models.Pedido, error) {
	var novoPedido models.Pedido
	if err:= json.NewDecoder(r.Body).Decode(&novoPedido); err != nil{
		erro:= errors.New("Erro ao ler o corpo da requisição")
		return novoPedido, erro
	}
	
	return novoPedido, nil
}

func CriarPedido(novoPedido models.Pedido) (models.Pedido, error) {
	if err := database.DB.Create(&novoPedido).Error; err != nil {
		erro:= errors.New("Erro ao criar o novo pedido")
		return novoPedido, erro
	}

	return novoPedido, nil
}