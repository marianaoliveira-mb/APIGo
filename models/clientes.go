package models

type Cliente struct {
	ClienteID       int     `json:"cliente_id"`
	NomeCliente     string  `json:"nome_cliente"`
	TelefoneCliente string  `json:"telefone_cliente"`
	Saldo           float64 `json:"saldo"`
}

var Clientes []Cliente
