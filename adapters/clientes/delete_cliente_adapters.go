package adapters

import(
	"errors"

	"github.com/Matari73/APIGo/database"
	"github.com/Matari73/APIGo/models"
)

func DeletarCliente(id string) error{
	var cliente models.Cliente
	result:= database.DB.Delete(&cliente, id)

	if result.Error != nil{
		erro:= errors.New("Erro ao excluir o cliente")
		return erro
	}

	return nil
}




// rows, result:= database.DB.Delete(&cliente, id).Rows()
// 	fmt.Println("rows")
// 	fmt.Println(rows)
// 	i := 0
// 	for rows {
// 		i++
// 		fmt.Println(i)
// 	}
// 	fmt.Println(i)
// 	if i == 0 {
// 		err:= errors.New("Cliente não encontrado com este ID")
// 		return err
// 	}