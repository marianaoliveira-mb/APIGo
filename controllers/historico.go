package controllers

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/Matari73/APIGo/database"
	"github.com/Matari73/APIGo/models"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func verificaClienteExistente(clienteID uint) (bool, error) {
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

func verificaVendedorExistente(vendedorID uint) (bool, error) {
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

func obterSaldoCliente(clienteID uint) (float64, error) {
	var cliente models.Cliente
	result := database.DB.First(&cliente, clienteID)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return 0.00000000, nil
		}
	}
	fmt.Println(result)
	return cliente.Saldo, nil
}

func atualizarSaldoCliente(clienteID uint, valorPedido float64) error {
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

func AdicionarProdutoAoPedidoHandler(w http.ResponseWriter, r *http.Request) {
	var dadosRequisicao struct {
		PedidoID   uint `json:"pedido_id"`
		ProdutoID  uint `json:"produto_id"`
		Quantidade int  `json:"quantidade"`
	}

	err := json.NewDecoder(r.Body).Decode(&dadosRequisicao)
	if err != nil {
		http.Error(w, "Erro ao decodificar o corpo da requisição", http.StatusBadRequest)
		return
	}

	fmt.Println("Chamando AdicionarProdutoAoPedido")
	err = AdicionarProdutoAoPedido(dadosRequisicao.PedidoID, dadosRequisicao.ProdutoID, dadosRequisicao.Quantidade)
	if err != nil {
		fmt.Printf("Erro ao adicionar produto ao pedido: %v\n", err)
		http.Error(w, fmt.Sprintf("Erro ao adicionar produto ao pedido: %v", err), http.StatusInternalServerError)
		return
	}

	fmt.Println("Produto adicionado ao pedido com sucesso")
	w.WriteHeader(http.StatusNoContent)
}

func AdicionarProdutoAoPedido(pedidoID uint, produtoID uint, quantidade int) error {
	fmt.Printf("Adicionando produto ao pedido. PedidoID: %d, ProdutoID: %d, Quantidade: %d\n", pedidoID, produtoID, quantidade)

	pedido := models.Pedido{}
	if err := database.DB.First(&pedido, pedidoID).Error; err != nil {
		return fmt.Errorf("Erro ao carregar o pedido: %v", err)
	}

	produto := models.Produto{}
	if err := database.DB.First(&produto, produtoID).Error; err != nil {
		return fmt.Errorf("Erro ao carregar o produto: %v", err)
	}

	if quantidade > produto.Estoque {
		return errors.New("Quantidade maior do que o estoque disponível")
	}

	novoEstoque := produto.Estoque - quantidade
	if err := database.DB.Model(&models.Produto{}).Where("produto_id = ?", produtoID).
		Update("estoque", novoEstoque).Error; err != nil {
		return fmt.Errorf("Erro ao atualizar o estoque do produto: %v", err)
	}

	produtoPedido := models.ProdutoPedido{}
	resultProdutoPedido := database.DB.Where("produto_id = ? AND pedido_id = ?", produtoID, pedidoID).First(&produtoPedido)
	if resultProdutoPedido.Error == nil {
		fmt.Println("A associação entre produto e pedido já existe.")
		return nil
	}

	produtoPedido = models.ProdutoPedido{
		PedidoID:   int(pedidoID),
		ProdutoID:  int(produtoID),
		Quantidade: quantidade,
	}

	resultCriacao := database.DB.Create(&produtoPedido)
	if resultCriacao.Error != nil {
		return fmt.Errorf("Erro ao criar a associação entre produto e pedido: %v", resultCriacao.Error)
	}

	fmt.Println("Associação entre produto e pedido criada com sucesso.")

	return nil
}

type HistoricoCompraStru struct {
	PedidoID     uint              `json:"pedido_id"`
	DataPedido   time.Time         `json:"data_pedido"`
	StatusPedido string            `json:"status_pedido"`
	ValorPedido  float64           `json:"valor_pedido"`
	Produtos     []*models.Produto `gorm:"many2many:produto_pedido;" json:"produtos"`
}

func HistoricoCompras(w http.ResponseWriter, r *http.Request) {
	clienteIDStr := mux.Vars(r)["cliente_id"]
	clienteID, err := strconv.Atoi(clienteIDStr)
	if err != nil {
		http.Error(w, "ID do cliente inválido", http.StatusBadRequest)
		return
	}

	var historicoCompras []HistoricoCompraStru

	err = database.DB.Table("pedido").
		Select("pedido.pedido_id, pedido.data_pedido, pedido.status_pedido, pedido.valor_pedido, produto.produto_id, produto.nome_produto, produto.valor_produto, produto.estoque, produto_pedido.quantidade").
		Joins("JOIN produto_pedido ON produto_pedido.pedido_id = pedido.pedido_id").
		Joins("JOIN produto ON produto_pedido.produto_id = produto.produto_id").
		Where("pedido.cliente_id = ?", clienteID).
		Scan(&historicoCompras).Error

	if err != nil {
		http.Error(w, "Erro ao obter os pedidos do cliente", http.StatusInternalServerError)
		return
	}

	for i := range historicoCompras {
		pedidoID := historicoCompras[i].PedidoID

		err = database.DB.Table("produto_pedido").
			Select("produto.produto_id, produto.nome_produto, produto.valor_produto, produto.estoque, produto_pedido.quantidade").
			Joins("JOIN produto ON produto_pedido.produto_id = produto.produto_id").
			Where("produto_pedido.pedido_id = ?", pedidoID).
			Scan(&historicoCompras[i].Produtos).Error

		if err != nil {
			http.Error(w, "Erro ao obter os produtos associados ao pedido", http.StatusInternalServerError)
			return
		}
	}

	w.Header().Set("Content-Type", "application/json")
	encoder := json.NewEncoder(w)

	if err := encoder.Encode(historicoCompras); err != nil {
		http.Error(w, fmt.Sprintf("Erro ao codificar para JSON: %v", err), http.StatusInternalServerError)
		return
	}
}
