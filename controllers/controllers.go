package controllers

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/Matari73/APIGo/database"
	"github.com/Matari73/APIGo/models"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Home Page")
}

//produtos
func GetProdutos(w http.ResponseWriter, r *http.Request) {
	var p []models.Produto
	if err := database.DB.Find(&p).Error; err != nil {
		http.Error(w, "Erro ao buscar produtos: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(p); err != nil {
		http.Error(w, "Erro ao codificar a resposta: "+err.Error(), http.StatusInternalServerError)
		return
	}
}

func GetProduto(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	var produto models.Produto

	if err := database.DB.First(&produto, id).Error; err != nil {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "Produto não encontrado")
		return
	}

	err := json.NewEncoder(w).Encode(produto)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Erro ao codificar produto em JSON: %v", err)
		return
	}
}

func CreateProduto(w http.ResponseWriter, r *http.Request) {
	var novoProduto models.Produto
	if err := json.NewDecoder(r.Body).Decode(&novoProduto); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Erro ao ler o corpo da requisição: %v", err)
		return
	}

	if novoProduto.NomeProduto == "" {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "O nome não pode ser vazio")
		return
	}

	if novoProduto.Estoque <= 0 {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "A quantidade do estoque deve ser maior que 0")
		return
	}

	if novoProduto.ValorProduto <= 0 {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "O valor do produto deve ser maior que 0")
		return
	}

	if err := database.DB.Create(&novoProduto).Error; err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Erro ao criar o produto: %v", err)
		return
	}

	if err := json.NewEncoder(w).Encode(novoProduto); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Erro ao codificar o produto em JSON: %v", err)
		return
	}
}

func DeleteProduto(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var produto models.Produto
	result := database.DB.Delete(&produto, id)
	if result.Error != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Erro ao excluir o produto: %v", result.Error)
		return
	}

	if result.RowsAffected == 0 {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "Produto não encontrado com o ID: %s", id)
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Produto excluído com sucesso")
}

func UpdateProduto(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var produto models.Produto
	err := json.NewDecoder(r.Body).Decode(&produto)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Erro ao decodificar o corpo da requisição: %v", err)
		return
	}

	if produto.NomeProduto == "" {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "O nome não pode ser vazio")
		return
	}

	if produto.Estoque < 0 {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "A quantidade do estoque deve ser maior ou igual a 0")
		return
	}

	if produto.ValorProduto <= 0 {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "O valor do produto deve ser maior que 0")
		return
	}

	if produto.ProdutoID == strconv.Itoa(0) {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "ID do produto não fornecido ou inválido")
		return
	}

	result := database.DB.Model(&models.Produto{}).Where("produto_id = ?", id).Updates(&produto)
	if result.Error != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Erro ao atualizar o produto no banco de dados: %v", result.Error)
		return
	}

	if result.RowsAffected == 0 {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprint(w, "Produto não encontrado")
		return
	}

	err = json.NewEncoder(w).Encode(produto)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Erro ao codificar produto em JSON: %v", err)
		return
	}
}

//clientes
func GetClientes(w http.ResponseWriter, r *http.Request) {
	var c []models.Cliente
	if err := database.DB.Find(&c).Error; err != nil {
		http.Error(w, "Erro ao buscar clientes: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(c); err != nil {
		http.Error(w, "Erro ao codificar a resposta: "+err.Error(), http.StatusInternalServerError)
		return
	}
}

func GetCliente(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	var cliente models.Cliente

	if err := database.DB.First(&cliente, id).Error; err != nil {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "Cliente não encontrado")
		return
	}

	err := json.NewEncoder(w).Encode(cliente)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Erro ao codificar cliente em JSON: %v", err)
		return
	}
}

func CreateCliente(w http.ResponseWriter, r *http.Request) {
	var novoCliente models.Cliente

	if err := json.NewDecoder(r.Body).Decode(&novoCliente); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Erro ao ler o corpo da requisição: %v", err)
		return
	}

	if novoCliente.NomeCliente == "" {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "O nome não deve ser vazio")
		return
	}

	if len(novoCliente.TelefoneCliente) < 10 || len(novoCliente.TelefoneCliente) > 12 {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "O telefone deve ter entre 10 e 12 caracteres")
		return
	}

	if err := database.DB.Create(&novoCliente).Error; err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Erro ao criar o novo cliente: %v", err)
		return
	}

	if err := json.NewEncoder(w).Encode(novoCliente); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Erro ao codificar o produto em JSON: %v", err)
		return
	}
}

func DeleteCliente(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var cliente models.Cliente
	result := database.DB.Delete(&cliente, id)
	if result.Error != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Erro ao excluir o cliente: %v", result.Error)
		return
	}

	if result.RowsAffected == 0 {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "Cliente não encontrado com o ID: %s", id)
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Cliente excluído com sucesso")
}

func UpdateCliente(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var cliente models.Cliente
	err := json.NewDecoder(r.Body).Decode(&cliente)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Erro ao decodificar o corpo da requisição: %v", err)
		return
	}

	if cliente.ClienteID == strconv.Itoa(0) {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "ID do cliente não fornecido ou inválido")
		return
	}

	result := database.DB.Model(&models.Cliente{}).Where("cliente_id = ?", id).Updates(&cliente)
	if result.Error != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Erro ao atualizar o cliente no banco de dados: %v", result.Error)
		return
	}

	if result.RowsAffected == 0 {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprint(w, "Cliente não encontrado")
		return
	}

	err = json.NewEncoder(w).Encode(cliente)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Erro ao codificar cliente em JSON: %v", err)
		return
	}
}

//vendedores
func GetVendedores(w http.ResponseWriter, r *http.Request) {
	var v []models.Vendedor
	if err := database.DB.Find(&v).Error; err != nil {
		http.Error(w, "Erro ao buscar vendedores: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(v); err != nil {
		http.Error(w, "Erro ao codificar a resposta: "+err.Error(), http.StatusInternalServerError)
		return
	}
}

func GetVendedor(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var vendedor models.Vendedor
	if err := database.DB.First(&vendedor, id).Error; err != nil {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "Vendedor não encontrado")
		return
	}

	err := json.NewEncoder(w).Encode(vendedor)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Erro ao codificar vendedor em JSON: %v", err)
		return
	}
}

func CreateVendedor(w http.ResponseWriter, r *http.Request) {
	var novoVendedor models.Vendedor

	if err := json.NewDecoder(r.Body).Decode(&novoVendedor); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Erro ao ler o corpo da requisição: %v", err)
		return
	}
	if novoVendedor.NomeVendedor == "" {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "O nome não deve ser vazio")
		return
	}

	if err := database.DB.Create(&novoVendedor).Error; err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Erro ao criar o novo cliente: %v", err)
		return
	}

	if err := json.NewEncoder(w).Encode(novoVendedor); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Erro ao codificar o produto em JSON: %v", err)
		return
	}
}

func DeleteVendedor(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var vendedor models.Vendedor
	result := database.DB.Delete(&vendedor, id)
	if result.Error != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Erro ao excluir o vendedor: %v", result.Error)
		return
	}

	if result.RowsAffected == 0 {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "Vendedor não encontrado com o ID: %s", id)
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "vendedor excluído com sucesso")
}

func UpdateVendedor(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var vendedor models.Vendedor
	err := json.NewDecoder(r.Body).Decode(&vendedor)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Erro ao decodificar o corpo da requisição: %v", err)
		return
	}

	if vendedor.VendedorID == strconv.Itoa(0) {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "ID do vendedor não fornecido ou inválido")
		return
	}

	result := database.DB.Model(&models.Vendedor{}).Where("vendedor_id = ?", id).Updates(&vendedor)
	if result.Error != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Erro ao atualizar o vendedor no banco de dados: %v", result.Error)
		return
	}

	if result.RowsAffected == 0 {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprint(w, "vendedor não encontrado")
		return
	}

	err = json.NewEncoder(w).Encode(vendedor)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Erro ao codificar o vendedor em JSON: %v", err)
		return
	}
}

//pedidos
func GetPedidos(w http.ResponseWriter, r *http.Request) {
	var p []models.Pedido
	if err := database.DB.Find(&p).Error; err != nil {
		http.Error(w, "Erro ao buscar clientes: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(p); err != nil {
		http.Error(w, "Erro ao codificar a resposta: "+err.Error(), http.StatusInternalServerError)
		return
	}
}
func GetPedido(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var pedido models.Pedido
	if err := database.DB.First(&pedido, id).Error; err != nil {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "Pedido não encontrado")
		return
	}

	err := json.NewEncoder(w).Encode(pedido)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Erro ao codificar produto em JSON: %v", err)
		return
	}
}

func CreatePedido(w http.ResponseWriter, r *http.Request) {
	var novoPedido models.Pedido

	if err := json.NewDecoder(r.Body).Decode(&novoPedido); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Erro ao ler o corpo da requisição: %v", err)
		return
	}

	novoPedido.DataPedido = time.Now()

	saldoCliente, err := obterSaldoCliente(uint(novoPedido.ClienteID))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Erro ao obter o saldo do cliente: %v", err)
		return
	}

	clienteID := uint(novoPedido.ClienteID)
	existeCliente, err := verificaClienteExistente(clienteID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Erro ao verificar o cliente: %v", err)
		return
	}

	if !existeCliente {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "ID do cliente inválido")
		return
	}

	vendedorID := uint(novoPedido.VendedorID)
	existeVendedor, err := verificaVendedorExistente(vendedorID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Erro ao verificar o Vendedor: %v", err)
		return
	}

	if !existeVendedor {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "ID do vendedor inválido")
		return
	}

	if strings.ToUpper(novoPedido.StatusPedido) != "EM ANDAMENTO" &&
		strings.ToUpper(novoPedido.StatusPedido) != "ENVIADO" &&
		strings.ToUpper(novoPedido.StatusPedido) != "CONCLUIDO" {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Status inválido")
		return
	}

	if saldoCliente < novoPedido.ValorPedido {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "Saldo do cliente insuficiente para o pedido")
		return
	}

	if err := database.DB.Create(&novoPedido).Error; err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Erro ao criar o novo cliente: %v", err)
		return
	}

	if err := atualizarSaldoCliente(uint(novoPedido.ClienteID), novoPedido.ValorPedido); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Erro ao atualizar o saldo do cliente: %v", err)
		return
	}

	if err := json.NewEncoder(w).Encode(novoPedido); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Erro ao codificar o produto em JSON: %v", err)
		return
	}
}

func DeletePedido(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var pedido models.Pedido
	result := database.DB.Delete(&pedido, id)
	if result.Error != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Erro ao excluir o pedido: %v", result.Error)
		return
	}

	if result.RowsAffected == 0 {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "Pedido não encontrado com o ID: %s", id)
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Pedido excluído com sucesso")
}

func UpdatePedido(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var pedido models.Pedido
	err := json.NewDecoder(r.Body).Decode(&pedido)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Erro ao decodificar o corpo da requisição: %v", err)
		return
	}

	if pedido.PedidoID == strconv.Itoa(0) {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "ID do pedido não fornecido ou inválido")
		return
	}

	result := database.DB.Model(&models.Pedido{}).Where("pedido_id = ?", id).Updates(&pedido)
	if result.Error != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Erro ao atualizar o pedido no banco de dados: %v", result.Error)
		return
	}

	if result.RowsAffected == 0 {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprint(w, "pedido não encontrado")
		return
	}

	err = json.NewEncoder(w).Encode(pedido)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Erro ao codificar o pedido em JSON: %v", err)
		return
	}
}

func verificaClienteExistente(clienteID uint) (bool, error) {
	var cliente models.Cliente
	result := database.DB.First(&cliente, clienteID)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return false, nil // Cliente não encontrado
		}
		return false, result.Error //erro ao consultar o banco
	}
	return true, nil // Se o cliente for encontrado
}

func verificaVendedorExistente(vendedorID uint) (bool, error) {
	var vendedor models.Vendedor
	result := database.DB.First(&vendedor, vendedorID)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return false, nil
		}
		return false, result.Error
	}
	return true, nil
}

func obterSaldoCliente(clienteID uint) (float64, error) {
	var cliente models.Cliente
	result := database.DB.First(&cliente, clienteID)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return 0.00000000, nil
		}
	}
	fmt.Println(result)
	return cliente.Saldo, nil
}

func atualizarSaldoCliente(clienteID uint, valorPedido float64) error {
	var cliente models.Cliente
	result := database.DB.Where("cliente_id = ?", clienteID).First(&cliente)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return fmt.Errorf("cliente com ID %d não encontrado", clienteID)
		}
		return result.Error
	}

	cliente.Saldo -= valorPedido
	if err := database.DB.Where("cliente_id = ?", clienteID).Save(&cliente).Error; err != nil {
		return fmt.Errorf("erro ao atualizar o saldo do cliente: %v", err)
	}

	return nil
}
