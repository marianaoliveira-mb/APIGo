package main

import (
	"fmt"

	"github.com/Matari73/APIGo/database"
	"github.com/Matari73/APIGo/routes"
)

func main() {
	database.ConectaComBancoDeDados()
	fmt.Println("Iniciando o servidor Rest com Go")
	routes.HandleRequest()
}
