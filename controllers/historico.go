package controllers

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
	"time"

	"github.com/darahayes/go-boom"
	"github.com/Matari73/APIGo/database"
	"github.com/Matari73/APIGo/models"
	"github.com/Matari73/APIGo/adapters/historico"
	"github.com/gorilla/mux"
)

func AdicionarProdutoAoPedidoHandler(w http.ResponseWriter, r *http.Request) {
	var dadosRequisicao struct {
		PedidoID   uint `json:"pedido_id"`
		ProdutoID  uint `json:"produto_id"`
		Quantidade int  `json:"quantidade"`
	}

	err := json.NewDecoder(r.Body).Decode(&dadosRequisicao)
	if err != nil {
		erro:= errors.New("Erro ao decodificar o corpo da requisição")
		boom.BadRequest(w, erro)
	}

	err = adapters.AdicionarProdutoAoPedido(dadosRequisicao.PedidoID, dadosRequisicao.ProdutoID, dadosRequisicao.Quantidade)
	if err != nil {
		boom.BadImplementation(w, err)
	}

	w.WriteHeader(http.StatusCreated)
	sucesso:= CreateResposta("Produto adicionado ao pedido!")
	json.NewEncoder(w).Encode(sucesso)
}


type HistoricoCompraStru struct {
	PedidoID     uint              `json:"pedido_id"`
	DataPedido   time.Time         `json:"data_pedido"`
	StatusPedido string            `json:"status_pedido"`
	ValorPedido  float64           `json:"valor_pedido"`
	Produtos     []*models.Produto `gorm:"many2many:produto_pedido;" json:"produtos"`
}

func HistoricoCompras(w http.ResponseWriter, r *http.Request) {
	vars:= mux.Vars(r)
	clienteIDStr , ok := vars["cliente_id"]
	if !ok {
		http.Error(w, "ID do cliente não encontrado na URL", http.StatusBadRequest)
		return
	}

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
		erro:= errors.New("Erro ao obter os pedidos do cliente")
		boom.BadImplementation(w, erro)
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
			erro:= errors.New("Erro ao obter os produtos associados ao pedido")
			boom.BadImplementation(w, erro)
			return
		}
	}

	w.Header().Set("Content-Type", "application/json")
	encoder := json.NewEncoder(w)
	if err := encoder.Encode(historicoCompras); err != nil {
		erro:= errors.New("Erro ao codificar para JSON")
		boom.BadImplementation(w, erro)
		return
	}
}

type HistoricoVendaStru struct {
	PedidoID     uint              `json:"pedido_id"`
	DataPedido   time.Time         `json:"data_pedido"`
	StatusPedido string            `json:"status_pedido"`
	ValorPedido  float64           `json:"valor_pedido"`
	Produtos     []*models.Produto `gorm:"many2many:produto_pedido;" json:"produtos"`
}

func HistoricoVendasVendedor(w http.ResponseWriter, r *http.Request) {
	vars:= mux.Vars(r)
	vendedorIDStr, ok := vars["vendedor_id"]

	if !ok {
		http.Error(w, "ID do vendedor não encontrado na URL", http.StatusBadRequest)
		return
	}

	vendedorID, err := strconv.Atoi(vendedorIDStr)
	if err != nil {
		http.Error(w, "ID do vendedor inválido", http.StatusBadRequest)
		return
	}

	//adapter
	var historicoVendas []HistoricoVendaStru

	err = database.DB.Table("pedido").
    Select("pedido.pedido_id, pedido.data_pedido, pedido.status_pedido, pedido.valor_pedido, produto.produto_id, produto.nome_produto, produto.valor_produto, produto.estoque, produto_pedido.quantidade").
    Joins("JOIN produto_pedido ON produto_pedido.pedido_id = pedido.pedido_id").
    Joins("JOIN produto ON produto_pedido.produto_id = produto.produto_id").
    Where("pedido.vendedor_id = ?", vendedorID).
    Scan(&historicoVendas).Error

	if err != nil {
		erro:= errors.New("Erro ao obter as vendas do vendedor")
		boom.BadImplementation(w, erro)
		return
	}

	for i := range historicoVendas {
		pedidoID := historicoVendas[i].PedidoID

		err = database.DB.Table("produto_pedido").
			Select("produto.produto_id, produto.nome_produto, produto.valor_produto, produto.estoque, produto_pedido.quantidade").
			Joins("JOIN produto ON produto_pedido.produto_id = produto.produto_id").
			Where("produto_pedido.pedido_id = ?", pedidoID).
			Scan(&historicoVendas[i].Produtos).Error

		if err != nil {
			erro:= errors.New("Erro ao obter os produtos associados ao pedido")
			boom.BadImplementation(w, erro)
			return
		}
	}

	w.Header().Set("Content-Type", "application/json")
	encoder := json.NewEncoder(w)
	if err := encoder.Encode(historicoVendas); err != nil {
		erro:= errors.New("Erro ao codificar para JSON")
		boom.BadImplementation(w, erro)
		return
	}
}

func HistoricoGeral(w http.ResponseWriter, r *http.Request)  {
	var historicoVendas []HistoricoVendaStru

	err := database.DB.Table("pedido").
    Select("pedido.pedido_id, pedido.data_pedido, pedido.status_pedido, pedido.valor_pedido, produto.produto_id, produto.nome_produto, produto.valor_produto, produto.estoque, produto_pedido.quantidade").
    Joins("JOIN produto_pedido ON produto_pedido.pedido_id = pedido.pedido_id").
    Joins("JOIN produto ON produto_pedido.produto_id = produto.produto_id").
    Scan(&historicoVendas).Error

	if err != nil {
		erro:= errors.New("Erro ao obter o histórico de pedidos")
		boom.BadImplementation(w, erro)
	}

	for i := range historicoVendas {
		pedidoID := historicoVendas[i].PedidoID

		err = database.DB.Table("produto_pedido").
			Select("produto.produto_id, produto.nome_produto, produto.valor_produto, produto.estoque, produto_pedido.quantidade").
			Joins("JOIN produto ON produto_pedido.produto_id = produto.produto_id").
			Where("produto_pedido.pedido_id = ?", pedidoID).
			Scan(&historicoVendas[i].Produtos).Error

		if err != nil {
			erro:= errors.New("Erro ao obter os produtos associados ao pedido")
			boom.BadImplementation(w, erro)
			return
		}
	}

	w.Header().Set("Content-Type", "application/json")
	encoder := json.NewEncoder(w)
	if err := encoder.Encode(historicoVendas); err != nil {
		erro:= errors.New("Erro ao codificar para JSON")
		boom.BadImplementation(w, erro)
		return
	}
}

