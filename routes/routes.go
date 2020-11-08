package routes

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jotafraga/go-rest-api/middlewares"
)

//HandleRoutes é o método responsável por lidar com as rotas da aplicação
func HandleRoutes() {
	router := mux.NewRouter()

	// swagger:operation GET /products getProduct GetProducts
	// ---
	// summary: Retornar a lista de produtos cadastrados
	// description: Todos os produtos persistidos no banco serão retonados.
	// responses:
	//   "200":
	//     "$ref": "#/responses/productRes"
	//   "400":
	//     "$ref": "#/responses/badReq"
	//   "404":
	//     "$ref": "#/responses/notFoundReq"
	router.HandleFunc("/products", middlewares.GetProducts).Methods("GET")

	// swagger:operation GET /product/{id} getProduct GetProduct
	// ---
	// summary: Retorna um produto específico cadastrado
	// description: Ao receber um ID, a aplicação retorna o respectivo produto cadastrado no banco.
	// parameters:
	// - name: id
	//   description: id of the product
	//   in: path
	//   required: true
	// responses:
	//   "200":
	//     "$ref": "#/responses/productRes"
	//   "400":
	//     "$ref": "#/responses/badReq"
	//   "404":
	//     "$ref": "#/responses/notFoundReq"
	router.HandleFunc("/product/{id}", middlewares.GetProduct).Methods("GET")

	// swagger:operation POST /product createProduct CreateProduct
	// ---
	// summary: Realiza a inserção do produto desejado
	// description: Método recebe os dados do produtos e realiza o cadastro do mesmo.
	// parameters:
	// - name: product
	//   description: product to add to the list of products
	//   in: body
	//   required: true
	//   schema:
	//     "$ref": "#/definitions/Product"
	// responses:
	//   "200":
	//     "$ref": "#/responses/productRes"
	//   "400":
	//     "$ref": "#/responses/badReq"
	//   "404":
	//     "$ref": "#/responses/notFoundReq"
	router.HandleFunc("/product", middlewares.CreateProduct).Methods("POST")

	// swagger:operation PUT /product/{id} updateProduct UpdateProduct
	// ---
	// summary: Atualiza um produto específico cadastrado
	// description: Ao receber um ID e o JSON contendo os dados a serem alterados, a aplicação atualiza o respectivo produto cadastrado no banco.
	// parameters:
	// - name: id
	//   description: id of the product
	//   in: path
	//   required: true
	// - name: product
	//   description: parameters of a product to be changed
	//   in: body
	//   required: true
	//   schema:
	//     "$ref": "#/definitions/Product"
	// responses:
	//   "200":
	//     "$ref": "#/responses/productRes"
	//   "400":
	//     "$ref": "#/responses/badReq"
	//   "404":
	//     "$ref": "#/responses/notFoundReq"
	router.HandleFunc("/product/{id}", middlewares.UpdateProduct).Methods("PUT")

	// swagger:operation DELETE /product/{id} deleteProduct DeleteProduct
	// ---
	// summary: Remove um produto específico
	// description: Ao receber um ID, a aplicação remove o respectivo produto cadastrado no banco.
	// parameters:
	// - name: id
	//   description: id of the product
	//   in: path
	//   required: true
	// responses:
	//   "200":
	//     "$ref": "#/responses/productRes"
	//   "400":
	//     "$ref": "#/responses/badReq"
	//   "404":
	//     "$ref": "#/responses/notFoundReq"
	router.HandleFunc("/product/{id}", middlewares.DeleteProduct).Methods("DELETE")

	//Endereço que contém a documentação da aplicação, gerada pelo Swagger
	router.PathPrefix("/").Handler(http.FileServer(http.Dir("./swaggerui/")))

	//Disponibilização da aplicação
	log.Fatal(http.ListenAndServe(":8000", router))
}
