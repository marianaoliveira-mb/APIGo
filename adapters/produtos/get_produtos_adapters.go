package adapters 

import(
	"encoding/json"
	"errors"
	"net/http"

	"github.com/Matari73/APIGo/database"
	"github.com/Matari73/APIGo/models"
)

func BuscarProdutos() ([]models.Produto, error) {
	var p []models.Produto
	if err := database.DB.Find(&p).Error; err != nil {
		erro:= errors.New("Erro ao buscar Produtos")
		return p, erro
	}

	return p, nil
}

func CodificarRespostaProduto(w http.ResponseWriter, produtos []models.Produto)  error  {
	w.Header().Set("Content-Type", "application/json")
	if err:= json.NewEncoder(w).Encode(produtos); err != nil{
		erro := errors.New("Erro ao codificar a resposta")
		return erro
	}

	return nil
}

func BuscarProdutoById(id string) (models.Produto, error) {
	var produto models.Produto
	if err := database.DB.First(&produto, id).Error; err != nil {
		erro:= errors.New("Produto não encontrado")
		return produto, erro
	}

	return produto, nil
}