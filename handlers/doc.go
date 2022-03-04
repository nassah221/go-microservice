// Package classification of Product API
//
// Documentation for Product API
//
//	Schemes: http
//	BasePath: /
//	Version: 1.0.0
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
// swagger:meta
package handlers

import "github.com/go-microservice/data"

// Generic error message returned as a string
// swagger:response errorResponse
type errResponseWrapper struct {
	// Description of the error
	// in: body
	Body GenericError
}

// Validation errors defined as an array of strings
// swagger:response errorValidation
type errValidationWrapper struct {
	Body ValidationError
}

// A list of products
// swagger:response productsResponse
type productsResponseWrapper struct { //nolint
	// All current products
	// in: body
	Body []data.Product
}

// A single product
// swagger:response productResponse
type productResponseWrapper struct { //nolint
	// The added product
	// in: body
	Body data.Product
}

// swagger:parameters deleteProduct
type productIDParameterWrapper struct { //nolint
	// ID of the product to delete from the products list
	// in: path
	// required: true
	ID int `json:"id"`
}
