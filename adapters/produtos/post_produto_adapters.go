package adapters 

import(
	"errors"

	"github.com/Matari73/APIGo/database"
	"github.com/Matari73/APIGo/models"
)

func VerificarSeExiste(novoProduto models.Produto)  error{
	var produtoExistente models.Produto
	if err := database.DB.Where("nome_produto = ?", novoProduto.NomeProduto).First(&produtoExistente).Error; err == nil{
		erro:= errors.New("Já existe um produto com este nome")
		return erro
	}

	return nil
}

func CriarProduto(novoProduto models.Produto) (models.Produto, error)  {
	if err := database.DB.Create(&novoProduto).Error; err != nil {
		erro:= errors.New("Erro ao criar o novo produto")
		return novoProduto, erro
	}

	return novoProduto, nil
}