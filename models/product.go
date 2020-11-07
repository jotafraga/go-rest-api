package models

import (
	"encoding/json"
	"net/http"

	"github.com/jotafraga/go-rest-api/database"
)

/*Product ... */
type Product struct {
	ProductID   int     `json:"productId,omitempty"`
	Name        string  `json:"name,omitempty"`
	Description string  `json:"description,omitempty"`
	Price       float64 `json:"price,omitempty"`
	Amount      int     `json:"amount,omitempty"`
}

/*GetProduct ... */
// swagger:operation GET /product/{id} getProduct GetProduct
// ---
// summary: Retorna um produto específico cadastrado
// description: Ao receber um ID, a aplicação retorna o respectivo produto cadastrado no banco.
// responses:
//   "200":
//     "$ref": "#/responses/okResp"
//   "400":
//     "$ref": "#/responses/badReq"
//   "404":
//     "$ref": "#/responses/notFoundReq"
func GetProduct(w http.ResponseWriter, r *http.Request) {

}

/*GetProducts ... */
// swagger:operation GET /products getProduct GetProducts
// ---
// summary: Retornar a lista de produtos cadastrados
// description: Todos os produtos persistidos no banco serão retonados.
// responses:
//   "200":
//     "$ref": "#/responses/okResp"
//   "400":
//     "$ref": "#/responses/badReq"
//   "404":
//     "$ref": "#/responses/notFoundReq"
func GetProducts(w http.ResponseWriter, r *http.Request) {
	database := database.ConectDataBase()

	selectProducts, err := database.Query("SELECT * FROM products")
	if err != nil {
		panic(err.Error())
	}

	p := Product{}
	products := []Product{}

	for selectProducts.Next() {
		var id, amount int
		var name, description string
		var price float64

		err = selectProducts.Scan(&id, &name, &description, &price, &amount)
		if err != nil {
			panic(err.Error())
		}

		p.ProductID = id
		p.Name = name
		p.Description = description
		p.Price = price
		p.Amount = amount

		products = append(products, p)
	}

	defer database.Close()

	json.NewEncoder(w).Encode(products)
}

/*CreateProduct ... */
// swagger:operation POST /product/{id} createProduct CreateProduct
// ---
// summary: Retorna um produto específico cadastrado
// description: Ao receber um ID, a aplicação retorna o respectivo produto cadastrado no banco.
// responses:
//   "200":
//     "$ref": "#/responses/okResp"
//   "400":
//     "$ref": "#/responses/badReq"
//   "404":
//     "$ref": "#/responses/notFoundReq"
func CreateProduct(w http.ResponseWriter, r *http.Request) {

}

/*UpdateProduct ... */
func UpdateProduct(w http.ResponseWriter, r *http.Request) {

}

/*DeleteProduct ... */
func DeleteProduct(w http.ResponseWriter, r *http.Request) {

}

// Success response
// swagger:response okResp
type swaggRespOk struct {
	// in:body
	Body struct {
		// HTTP status code 200 - OK
		Code int `json:"code"`
	}
}

// Error Bad Request
// swagger:response badReq
type swaggReqBadRequest struct {
	// in:body
	Body struct {
		// HTTP status code 400 -  Bad Request
		Code int `json:"code"`
	}
}

// Error Not Found
// swagger:response notFoundReq
type swaggReqNotFound struct {
	// in:body
	Body struct {
		// HTTP status code 404 -  Not Found
		Code int `json:"code"`
	}
}
