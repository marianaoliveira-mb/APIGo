package models

type Vendedor struct {
	VendedorID   int    `json:"vendedor_id"`
	NomeVendedor string `json:"nome_vendedor"`
}

var Vendedores []Vendedor
