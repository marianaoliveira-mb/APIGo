package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Matari73/APIGo/database"
	"github.com/Matari73/APIGo/routes"
)

func main() {
	database.ConectaComBancoDeDados()
	fmt.Println("Iniciando o servidor com Go")
	r := routes.HandleResquest()
	log.Fatal(http.ListenAndServe(":8000", r))
}
