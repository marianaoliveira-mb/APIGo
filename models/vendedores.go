package models

type Vendedor struct {
	VendedorID   string `json:"vendedor_id" gorm:"default:uuid_generate_v3()"`
	NomeVendedor string `json:"nome_vendedor"`
}

var Vendedores []Vendedor
