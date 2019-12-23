package main

import "testing"

type testRepository struct {
	shops             map[string]string
	customers         map[string]string
	shopsGetCount     map[string]int
	customersGetCount map[string]int
}

func (t *testRepository) PutShopToken(shopID string, token string) error {
	t.shops[shopID] = token
	return nil
}

func (t *testRepository) PutCustomerToken(customerID string, token string) error {
	t.customers[customerID] = token
	return nil
}

func (t *testRepository) GetShopToken(shopID string) (string, error) {
	t.shopsGetCount[shopID]++
	return t.shops[shopID], nil
}

func (t *testRepository) GetCustomerToken(customerID string) (string, error) {
	t.customersGetCount[customerID]++
	return t.customers[customerID], nil
}

type testAPIWrapper struct {
	lastSentMessages map[string]string
}

func (t *testAPIWrapper) Send(token string, message string) error {
	t.lastSentMessages[token] = message
	return nil
}

func newTestRepository() (repository *testRepository) {
	repository = new(testRepository)
	repository.shops = make(map[string]string)
	repository.customers = make(map[string]string)
	repository.shopsGetCount = make(map[string]int)
	repository.customersGetCount = make(map[string]int)
	return
}

func newTestAPIWrapper() (wrapper *testAPIWrapper) {
	wrapper = new(testAPIWrapper)
	wrapper.lastSentMessages = make(map[string]string)
	return
}

func TestSend(t *testing.T) {
	// given
	shopID := "200"
	shopToken := "acbdef"
	customerID := "john"
	customerToken := "123abc"
	var service = new(NotificationService)
	repository := newTestRepository()
	apiWrapper := newTestAPIWrapper()
	service.TokenRepository = repository
	service.APIWrapper = apiWrapper

	// when
	service.subscribeShop(shopID, shopToken)
	service.subscribeCustomer(customerID, customerToken)
	service.sendShop(shopID, "Hello")
	service.sendCustomer(customerID, "Hi")
	service.sendCustomer(customerID, "Bye")

	// then
	if _, exist := repository.shops[shopID]; !exist {
		t.Error("Repository doesn't have shop")
	}
	if _, exist := repository.customers[customerID]; !exist {
		t.Error("Repository doesn't have customer")
	}
	if repository.shopsGetCount[shopID] == 0 {
		t.Errorf("Repository never have returned shop with id %s", shopID)
	}
	if repository.customersGetCount[customerID] == 0 {
		t.Errorf("Repository never have returned customer with id %s", customerID)
	}
	if apiWrapper.lastSentMessages[shopToken] != "Hello" {
		t.Errorf("Incorrect last message: '%s'", apiWrapper.lastSentMessages[shopToken])
	}
	if apiWrapper.lastSentMessages[customerToken] != "Bye" {
		t.Errorf("Incorrect last message: '%s'", apiWrapper.lastSentMessages[customerToken])
	}
}
