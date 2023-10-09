package models

import "time"

type Pedido struct {
	PedidoID     int       `json:"pedido_id"`
	DataPedido   time.Time `json:"data_pedido"`
	StatusPedido string    `json:"status_pedido"`
	ValorPedido  float64   `json:"valor_pedido"`
	Quantidade   int       `json:"quantidade"`
	ClienteID    int       `json:"cliente_id"`
	VendedorID   int       `json:"vendedor_id"`
}

var Pedidos []Pedido
