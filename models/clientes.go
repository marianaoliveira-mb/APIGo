package models

type Cliente struct {
	ClienteID       string  `json:"cliente_id" gorm:"default:uuid_generate_v3()"`
	NomeCliente     string  `json:"nome_cliente"`
	TelefoneCliente string  `json:"telefone_cliente"`
	Saldo           float64 `json:"saldo"`
}

var Clientes []Cliente
