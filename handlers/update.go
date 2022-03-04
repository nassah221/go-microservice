package handlers

import (
	"net/http"
	"strconv"

	"github.com/go-microservice/data"
	"github.com/gorilla/mux"
)

// swagger:route PUT /products/{id} product updateProduct
// Updates a product
// responses:
//  201: noContent
//  404: errorResponse
//  422: errorValidation

// UpdateProduct updates the product
func (p *Products) UpdateProduct(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle PUT Products")

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(rw, "Cannot parse id from URL Path", http.StatusBadRequest)
		return
	}

	v := r.Context().Value(KeyProduct{})
	prod, ok := v.(data.Product)
	if !ok {
		p.l.Panic("unable to cast context value into product")
	}

	if err := data.UpdateProduct(id, &prod); err != nil {
		if err == data.ErrProductNotFound {
			rw.WriteHeader(http.StatusNotFound)
			data.ToJSON(&GenericError{Message: "Product not found"}, rw)
			return
		}

		p.l.Panicf("Unexpected error: %v", err)
	}
}
