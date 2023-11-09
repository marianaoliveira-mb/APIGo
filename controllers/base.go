package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Home Page")
}

func codificarEmJson(w http.ResponseWriter, model interface{}) error {
	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(model)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return err
	}
	return nil
}

type Resposta struct{
	Mensagem string `json:"mensagem"`
}

func CreateResposta(mensagem string) Resposta {
	return Resposta{
		Mensagem: mensagem,
	}
}