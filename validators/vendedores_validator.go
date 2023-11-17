package validators

import(
	"errors"

	"github.com/Matari73/APIGo/models"
)

func ValidateNome(vendedor models.Vendedor) error  {
	if vendedor.NomeVendedor == "" {
		return errors.New("O nome não deve ser vazio")
	}

	return nil
}