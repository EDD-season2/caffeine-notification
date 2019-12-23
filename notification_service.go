package main

import "fmt"

// NotificationService provides make api call with given message and shop/customer id
type NotificationService struct {
	TokenRepository Repository
	APIWrapper      NotificationAPIWrapper
}

func (s *NotificationService) subscribeShop(shopID string, token string) {
	s.TokenRepository.PutShopToken(shopID, token)
}

func (s *NotificationService) subscribeCustomer(customerID string, token string) {
	s.TokenRepository.PutCustomerToken(customerID, token)
}

// ShopTokenNotFoundError is occcurred when token matches given shop id is not found
type ShopTokenNotFoundError struct {
	cause string
}

func (e *ShopTokenNotFoundError) Error() string {
	return fmt.Sprintf("No token is found with shop id '%s'", e.cause)
}

// CustomerTokenNotFoundError is occcurred when token matches given customer id is not found
type CustomerTokenNotFoundError struct {
	cause string
}

func (e *CustomerTokenNotFoundError) Error() string {
	return fmt.Sprintf("No token is found with customer id '%s'", e.cause)
}

func (s *NotificationService) sendShop(shopID string, message string) error {
	token, err := s.TokenRepository.GetShopToken(shopID)
	if token == "" {
		return &ShopTokenNotFoundError{shopID}
	} else if err != nil {
		return err
	}
	err = s.APIWrapper.Send(token, message)
	if err != nil {
		return err
	}
	return nil
}

func (s *NotificationService) sendCustomer(customerID string, message string) error {
	token, err := s.TokenRepository.GetCustomerToken(customerID)
	if token == "" {
		return &CustomerTokenNotFoundError{customerID}
	} else if err != nil {
		return err
	}
	err = s.APIWrapper.Send(token, message)
	if err != nil {
		return err
	}
	return nil
}

func NewNotificationService(apiWrapper NotificationAPIWrapper, repository Repository) *NotificationService {
	service := new(NotificationService)
	service.APIWrapper = apiWrapper
	service.TokenRepository = repository
	return service
}
