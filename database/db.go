package database

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var (
	DB  *gorm.DB
	err error
)

func ConectaComBancoDeDados() {
	stringDeConexao := "host=localhost user=postgres password=postgres dbname=postgres port=5432 sslmode=disable"
	DB, err = gorm.Open(postgres.Open(stringDeConexao), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true, // Evitar pluralização das tabelas
		},
	})
	if err != nil {
		log.Panic("Erro ao conectar com banco de dados")
	}
}
