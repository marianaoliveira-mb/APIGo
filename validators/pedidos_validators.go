package validators

import(
	"strings"
	"time"
	"errors"
	"fmt"
	"gorm.io/gorm"

	"github.com/Matari73/APIGo/database"
	"github.com/Matari73/APIGo/models"

)
func ValidatePedido(novoPedido models.Pedido) error {
	novoPedido.DataPedido = time.Now()
	saldoCliente, err := ObterSaldoCliente(uint(novoPedido.ClienteID))
	if err != nil {
		return errors.New("Erro ao obter o saldo do cliente")
	}

	clienteID := uint(novoPedido.ClienteID)
	existeCliente, err := VerificaClienteExistente(clienteID)
	if err != nil {
		return errors.New("Erro ao verificar o cliente")
	}

	if !existeCliente {
		return errors.New("ID do cliente inválido")

	}

	vendedorID := uint(novoPedido.VendedorID)
	existeVendedor, err := VerificaVendedorExistente(vendedorID)
	if err != nil {
		return errors.New("Erro ao verificar o Vendedor")

	}

	if !existeVendedor {
		return errors.New("ID do vendedor inválido")
	}

	if strings.ToUpper(novoPedido.StatusPedido) != "EM ANDAMENTO" &&
		strings.ToUpper(novoPedido.StatusPedido) != "ENVIADO" &&
		strings.ToUpper(novoPedido.StatusPedido) != "CONCLUIDO" {
		return errors.New("Status inválido")
	}

	if saldoCliente < novoPedido.ValorPedido {
		return errors.New("Saldo do cliente insuficiente para o pedido")
	}

	return nil
}


func VerificaClienteExistente(clienteID uint) (bool, error) {
	var cliente models.Cliente
	result := database.DB.First(&cliente, clienteID)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return false, nil // Cliente não encontrado
		}
		return false, result.Error //erro ao consultar o banco
	}
	return true, nil // Se o cliente for encontrado
}

func VerificaVendedorExistente(vendedorID uint) (bool, error) {
	var vendedor models.Vendedor
	result := database.DB.First(&vendedor, vendedorID)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return false, nil
		}
		return false, result.Error
	}
	return true, nil
}

func ObterSaldoCliente(clienteID uint) (float64, error) {
	var cliente models.Cliente
	result := database.DB.First(&cliente, clienteID)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return 0.00000000, nil
		}
	}
	return cliente.Saldo, nil
}

func AtualizarSaldoCliente(clienteID uint, valorPedido float64) error {
	var cliente models.Cliente
	result := database.DB.Where("cliente_id = ?", clienteID).First(&cliente)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return fmt.Errorf("cliente com ID %d não encontrado", clienteID)
		}
		return result.Error
	}

	cliente.Saldo -= valorPedido
	if err := database.DB.Where("cliente_id = ?", clienteID).Save(&cliente).Error; err != nil {
		return fmt.Errorf("erro ao atualizar o saldo do cliente: %v", err)
	}

	return nil
}
