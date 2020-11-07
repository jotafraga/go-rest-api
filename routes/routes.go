package routes

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jotafraga/go-rest-api/models"
)

func HandleRoutes() {
	router := mux.NewRouter()

	//ENDPOINTS DA APLICAÇÃO
	router.HandleFunc("/products", models.GetProducts).Methods("GET")
	router.HandleFunc("/product/{id}", models.GetProduct).Methods("GET")
	router.HandleFunc("/product/{id}", models.CreateProduct).Methods("POST")
	router.HandleFunc("/product/{id}", models.UpdateProduct).Methods("PUT")
	router.HandleFunc("/product/{id}", models.DeleteProduct).Methods("DELETE")

	//DOCUMENTAÇÃO APLICAÇÃO
	router.PathPrefix("/").Handler(http.FileServer(http.Dir("./swaggerui/")))

	log.Fatal(http.ListenAndServe(":8000", router))
}
