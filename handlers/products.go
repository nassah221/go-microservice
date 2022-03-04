package handlers

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/go-microservice/data"
	"github.com/gorilla/mux"
)

type Products struct {
	l *log.Logger
}

type KeyProduct struct{}

// NewProducts creates a products handler with the given logger
func NewProducts(l *log.Logger) *Products {
	return &Products{l}
}

// getProducts returns the products from the data store
func (p *Products) GetProducts(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle GET Products")

	// fetch the products from the datastore
	lp := data.GetProducts()

	// serialize the list to JSON
	err := lp.ToJSON(rw)
	if err != nil {
		http.Error(rw, "Unable to marshal json", http.StatusInternalServerError)
		return
	}
}

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
		http.Error(rw, "Unable cast request context into product", http.StatusInternalServerError)
		return
	}

	if err := data.UpdateProduct(id, &prod); err != nil {
		if err == data.ErrProductNotFound {
			http.Error(rw, "Product id not found", http.StatusNotFound)
			return
		}

		http.Error(rw, "Product id not found", http.StatusInternalServerError)
		return
	}
}

func (p Products) MiddlewareProductValidation(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		var prod data.Product

		if err := prod.FromJSON(r.Body); err != nil {
			p.l.Println(err)
			http.Error(rw, "Unable to encode JSON", http.StatusBadRequest)
			return
		}

		if err := prod.Validate(); err != nil {
			if err == data.ErrRegisterValidation {
				p.l.Println("[ERROR] Registering custom validator")
				http.Error(rw, "", http.StatusInternalServerError)
				return
			}
			p.l.Println("[ERROR] validating product")
			http.Error(rw, fmt.Sprintf("Error validating payload: %s", err), http.StatusBadRequest)
			return
		}

		ctx := context.WithValue(r.Context(), KeyProduct{}, prod)
		req := r.WithContext(ctx)

		next.ServeHTTP(rw, req)
	})
}
