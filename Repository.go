package main

// Repository is abstraction of persistence layer
type Repository interface {
	PutShopToken(shopID string, token string) error
	PutCustomerToken(customerID string, token string) error
	GetShopToken(shopID string) (string, error)
	GetCustomerToken(customerID string) (string, error)
}
