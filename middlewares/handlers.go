package middlewares

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	"github.com/jotafraga/go-rest-api/models"
)

func createConnection() *sqlx.DB {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("ERRO - O sistema não conseguiu encontrar o arquivo de variavéis de ambiente.")
	}

	db, err := sqlx.Connect("mysql", os.Getenv("MYSQL_URL"))
	if err != nil {
		log.Fatalf("ERRO - O sistema não conseguiu estabelecer conexão com o banco de dados.")
	}

	return db
}

//GetProduct ...
func GetProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	params := mux.Vars(r)

	id, err := strconv.Atoi(params["id"])
	if err != nil {
		res := models.Response{
			Message: "ERRO - Falha no ID! Verifique se o ID está correto e é um inteiro.",
		}

		json.NewEncoder(w).Encode(res)
		return
	}

	product, err := getProduct(int64(id))
	if err != nil {
		log.Fatalf("ERRO - Falha ao obter o produto solicitado. %v", err)
	}

	json.NewEncoder(w).Encode(product)
}

//GetProducts ...
func GetProducts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	users, err := getProducts()
	if err != nil {
		log.Fatalf("ERRO - Falha ao obter lista de produtos. %v", err)
	}

	json.NewEncoder(w).Encode(users)
}

//CreateProduct ...
func CreateProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	var product models.Product

	err := json.NewDecoder(r.Body).Decode(&product)
	if err != nil {
		res := models.Response{
			Message: "ERRO - Falha ao capturar corpo da requisição! Verifique se os parâmetros foram enviados corretamente.",
		}

		json.NewEncoder(w).Encode(res)
		return
	}

	insertID := insertProduct(product)

	res := models.Response{
		ID:      insertID,
		Message: "Produto Inserido com Sucesso!",
	}

	json.NewEncoder(w).Encode(res)
}

// UpdateProduct ...
func UpdateProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "PUT")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		res := models.Response{
			Message: "ERRO - Falha no ID! Verifique se o ID está correto e é um inteiro.",
		}

		json.NewEncoder(w).Encode(res)
		return
	}

	var product models.Product

	err = json.NewDecoder(r.Body).Decode(&product)
	if err != nil {
		res := models.Response{
			Message: "ERRO - Falha ao capturar corpo da requisição! Verifique se os parâmetros foram enviados corretamente.",
		}

		json.NewEncoder(w).Encode(res)
		return
	}

	updatedRows := updateProduct(int64(id), product)

	var msg string
	if updatedRows > 0 {
		msg = "Produto atualizado com sucesso!"
	} else {
		msg = "Nenhum produto foi atualizado. Verificar o ID do produto e parâmetros enviados!"
	}

	res := models.Response{
		ID:      int64(id),
		Message: msg,
	}

	json.NewEncoder(w).Encode(res)
}

// DeleteProduct ...
func DeleteProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	params := mux.Vars(r)

	id, err := strconv.Atoi(params["id"])
	if err != nil {
		res := models.Response{
			Message: "ERRO - Falha no ID! Verifique se o ID está correto e é um inteiro.",
		}

		json.NewEncoder(w).Encode(res)
		return
	}

	deletedRows := deleteProduct(int64(id))

	var msg string
	if deletedRows > 0 {
		msg = "Produto removido com sucesso!"
	} else {
		msg = "Nenhum produto foi removido. Verificar o ID do produto!"
	}

	res := models.Response{
		ID:      int64(id),
		Message: msg,
	}

	json.NewEncoder(w).Encode(res)
}

func getProduct(id int64) (models.Product, error) {
	db := createConnection()
	defer db.Close()

	var product models.Product

	sqlStatement := `SELECT * FROM products WHERE product_id=?`
	row := db.QueryRow(sqlStatement, id)
	err := row.Scan(&product.ProductID, &product.Name, &product.Description, &product.Price, &product.Amount)
	switch err {
	case sql.ErrNoRows:
		return product, nil
	case nil:
		return product, nil
	default:
		log.Fatalf("ERRO - Falha ao realizar a leitura do retorno do banco. %v", err)
	}

	return product, err
}

func getProducts() ([]models.Product, error) {
	db := createConnection()
	defer db.Close()

	var products []models.Product
	err := db.Select(&products, "SELECT * FROM products")
	if err != nil {
		log.Fatalf("ERRO - Falha ao executar Query para obter produtos. %v", err)
	}

	return products, err
}

func insertProduct(product models.Product) int64 {
	db := createConnection()
	defer db.Close()

	sqlStatement := `INSERT INTO products (name, description, price, amount) VALUES (?, ?, ?, ?)`
	res, err := db.Exec(sqlStatement, product.Name, product.Description, product.Price, product.Amount)
	if err != nil {
		log.Fatalf("ERRO - Falha ao executar Query de Inserção. %v", err)
	}

	productID, _ := res.LastInsertId()

	return productID
}

func updateProduct(id int64, product models.Product) int64 {
	db := createConnection()
	defer db.Close()

	sqlStatement := `UPDATE products SET name=?, description=?, price=?, amount=? WHERE product_id=?`
	res, err := db.Exec(sqlStatement, &product.Name, &product.Description, &product.Price, &product.Amount, id)
	if err != nil {
		log.Fatalf("ERRO - Falha ao executar Query para atualizar o produto. %v", err)
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		log.Fatalf("ERRO - Falha ao capturar o número de linhas afetadas pela Query anterior. %v", err)
	}

	return rowsAffected
}

func deleteProduct(id int64) int64 {
	db := createConnection()
	defer db.Close()

	sqlStatement := `DELETE FROM products WHERE product_id=?`
	res, err := db.Exec(sqlStatement, id)
	if err != nil {
		log.Fatalf("ERRO - Falha ao executar Query para remover o produto. %v", err)
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		log.Fatalf("ERRO - Falha ao capturar o número de linhas afetadas pela Query anterior. %v", err)
	}

	return rowsAffected
}
