package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// @title Orders API
// @version 1.0
// @description Simple Orders API
// @termOfServices https://hackitv8.com/tos
// @contact.name API Support
// @contact.email api@hacktiv8.com
// @license.name Apache 2.0
// @license.url http://hacktiv8.com
// @host http://localhost:8080
// @BasePath /
func main() {
	router := mux.NewRouter()

	router.HandleFunc("/orders", getOrders)
	log.Fatal(http.ListenAndServe(":8080", router))
}

// GetOrders godoc
// @Summary list of orders
// @Description this endpoint will return array of orders
// @Tags orders
// @Accept json
// @Produce json
// @Success 200 {array} Order
// @Router /orders [get]
func getOrders(rw http.ResponseWriter, r *http.Request) {

}
