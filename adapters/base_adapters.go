package adapters 

// import(
// 	"encoding/json"
// 	"errors"
// 	"net/http"

// 	"github.com/Matari73/APIGo/database"
// 	"github.com/Matari73/APIGo/models"
// )

// func BuscarClienteById(id string) (models.Cliente, error) {
// 	var cliente models.Cliente
// 	if err := database.DB.First(&cliente, id).Error; err != nil{
// 		erro:= errors.New("Cliente não encontrado")
// 		return cliente, erro
// 	}

// 	return cliente, nil
// }