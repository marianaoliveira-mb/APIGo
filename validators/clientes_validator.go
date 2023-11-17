package validators

import(
	"errors"

	"github.com/Matari73/APIGo/models"
)

func ValidateTelNome(cliente models.Cliente) error {

	if cliente.NomeCliente == "" {
		return errors.New("O nome não deve ser vazio")
	}

	if len(cliente.TelefoneCliente) < 10 || len(cliente.TelefoneCliente) > 12 {
		return errors.New("O telefone deve ter entre 10 e 12 caracteres")
	}

	return nil
}

