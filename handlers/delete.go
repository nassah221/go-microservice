package handlers

import (
	"net/http"
	"strconv"

	"github.com/go-microservice/data"
	"github.com/gorilla/mux"
)

// swagger:route DELETE /products/{id} products deleteProduct
// Deletes a product
// responses:
//  201: noContent

// getProducts returns the products from the data store
func (p *Products) DeleteProduct(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle DELETE Product")

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(rw, "Unable to parse id", http.StatusBadRequest)
		return
	}

	if err := data.DeleteProduct(id); err != nil {
		http.Error(rw, "Product not found", http.StatusNotFound)
		return
	}
}
