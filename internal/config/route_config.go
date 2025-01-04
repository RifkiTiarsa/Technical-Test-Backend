package config

const (
	ApiGroup = "/api/v1"

	// Auth Route
	Register = "/auth/register"
	Login    = "/auth/login"

	// Product Route
	GetProduct     = "/products"
	GetProductByID = "/products/:id"

	// Cart Route
	AddCart = "/carts"

	// Checkout Route
	CreateCheckout = "/checkout"
)
