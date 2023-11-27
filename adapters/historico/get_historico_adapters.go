package adapters

import(
	"errors"
	"net/http"
	"time"
	"encoding/json"

	"github.com/Matari73/APIGo/database"
	"github.com/Matari73/APIGo/models"
)

type HistoricoVendaStru struct {
	PedidoID     uint              `json:"pedido_id"`
	DataPedido   time.Time         `json:"data_pedido"`
	StatusPedido string            `json:"status_pedido"`
	ValorPedido  float64           `json:"valor_pedido"`
	Produtos     []*models.Produto `gorm:"many2many:produto_pedido;" json:"produtos"`
}

func ObterVendas(w http.ResponseWriter, vendedorID int)  error{
	var historicoVendas []HistoricoVendaStru

	err := database.DB.Table("pedido").
    Select("pedido.pedido_id, pedido.data_pedido, pedido.status_pedido, pedido.valor_pedido, produto.produto_id, produto.nome_produto, produto.valor_produto, produto.estoque, produto_pedido.quantidade").
    Joins("JOIN produto_pedido ON produto_pedido.pedido_id = pedido.pedido_id").
    Joins("JOIN produto ON produto_pedido.produto_id = produto.produto_id").
    Where("pedido.vendedor_id = ?", vendedorID).
    Scan(&historicoVendas).Error

	if err != nil {
		return errors.New("Erro ao obter as vendas do vendedor")
	}

	for i := range historicoVendas {
		pedidoID := historicoVendas[i].PedidoID

		err := database.DB.Table("produto_pedido").
			Select("produto.produto_id, produto.nome_produto, produto.valor_produto, produto.estoque, produto_pedido.quantidade").
			Joins("JOIN produto ON produto_pedido.produto_id = produto.produto_id").
			Where("produto_pedido.pedido_id = ?", pedidoID).
			Scan(&historicoVendas[i].Produtos).Error

		if err != nil {
			return errors.New("Erro ao obter os produtos associados ao pedido")
		}
	}

	w.Header().Set("Content-Type", "application/json")
	encoder := json.NewEncoder(w)
	if err := encoder.Encode(historicoVendas); err != nil {
		return errors.New("Erro ao codificar para JSON")
	}

	return nil
}


type HistoricoCompraStru struct {
	PedidoID     uint              `json:"pedido_id"`
	DataPedido   time.Time         `json:"data_pedido"`
	StatusPedido string            `json:"status_pedido"`
	ValorPedido  float64           `json:"valor_pedido"`
	Produtos     []*models.Produto `gorm:"many2many:produto_pedido;" json:"produtos"`
}

func ObterCompras(w http.ResponseWriter, clienteID int) error {
	var historicoCompras []HistoricoCompraStru

	err := database.DB.Table("pedido").
    Select("pedido.pedido_id, pedido.data_pedido, pedido.status_pedido, pedido.valor_pedido, produto.produto_id, produto.nome_produto, produto.valor_produto, produto.estoque, produto_pedido.quantidade").
    Joins("JOIN produto_pedido ON produto_pedido.pedido_id = pedido.pedido_id").
    Joins("JOIN produto ON produto_pedido.produto_id = produto.produto_id").
    Where("pedido.cliente_id = ?", clienteID).
    Scan(&historicoCompras).Error

	if err != nil {
		return errors.New("Erro ao obter os pedidos do cliente")
	}

	for i := range historicoCompras {
		pedidoID := historicoCompras[i].PedidoID

		err := database.DB.Table("produto_pedido").
			Select("produto.produto_id, produto.nome_produto, produto.valor_produto, produto.estoque, produto_pedido.quantidade").
			Joins("JOIN produto ON produto_pedido.produto_id = produto.produto_id").
			Where("produto_pedido.pedido_id = ?", pedidoID).
			Scan(&historicoCompras[i].Produtos).Error

		if err != nil {
			return errors.New("Erro ao obter os produtos associados ao pedido")
		}
	}

	w.Header().Set("Content-Type", "application/json")
	encoder := json.NewEncoder(w)
	if err := encoder.Encode(historicoCompras); err != nil {
		return errors.New("Erro ao codificar para JSON")
	}

	return nil
}

func ObterHistoricoGeral(w http.ResponseWriter)  error{
	var historicoVendas []HistoricoVendaStru

	err := database.DB.Table("pedido").
    Select("pedido.pedido_id, pedido.data_pedido, pedido.status_pedido, pedido.valor_pedido, produto.produto_id, produto.nome_produto, produto.valor_produto, produto.estoque, produto_pedido.quantidade").
    Joins("JOIN produto_pedido ON produto_pedido.pedido_id = pedido.pedido_id").
    Joins("JOIN produto ON produto_pedido.produto_id = produto.produto_id").
    Scan(&historicoVendas).Error

	if err != nil {
		return errors.New("Erro ao obter o histórico de pedidos")
	}

	for i := range historicoVendas {
		pedidoID := historicoVendas[i].PedidoID

		err = database.DB.Table("produto_pedido").
			Select("produto.produto_id, produto.nome_produto, produto.valor_produto, produto.estoque, produto_pedido.quantidade").
			Joins("JOIN produto ON produto_pedido.produto_id = produto.produto_id").
			Where("produto_pedido.pedido_id = ?", pedidoID).
			Scan(&historicoVendas[i].Produtos).Error

		if err != nil {
			return errors.New("Erro ao obter os produtos associados ao pedido")
		}
	}

	w.Header().Set("Content-Type", "application/json")
	encoder := json.NewEncoder(w)
	if err := encoder.Encode(historicoVendas); err != nil {
		return errors.New("Erro ao codificar para JSON")
	}

	return nil
}