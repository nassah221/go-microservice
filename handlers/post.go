package handlers

import (
	"net/http"

	"github.com/go-microservice/data"
)

// swagger:route POST /products products addProduct
// Adds a product
// responses:
//  200: productResponse
//  422: errorValidation
//  501: errorResponse
// AddProduct adds a product to the products list
func (p *Products) AddProduct(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle POST Products")

	v := r.Context().Value(KeyProduct{})
	prod, ok := v.(data.Product)
	if !ok {
		http.Error(rw, "Unable cast request context into product", http.StatusInternalServerError)
		return
	}

	data.AddProduct(&prod)
}
