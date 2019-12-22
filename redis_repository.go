package main

import "github.com/go-redis/redis"

// RedisRepository is implementation of Repository interface using redis
type RedisRepository struct {
	client *redis.Client
}

// PutShopToken saves shop id - token pair to
func (r *RedisRepository) PutShopToken(shopID string, token string) error {
	return r.client.Set("shop:"+shopID, token, 0).Err()
}

// PutCustomerToken saves customer id - token pair
func (r *RedisRepository) PutCustomerToken(customerID string, token string) error {
	return r.client.Set("customer:"+customerID, token, 0).Err()
}

// GetShopToken retrieves token matches given shop id
func (r *RedisRepository) GetShopToken(shopID string) (string, error) {
	value, err := r.client.Get("shop:" + shopID).Result()
	if err != nil {
		return "", err
	}
	return value, nil
}

// GetCustomerToken retrieves token matches given customer id
func (r *RedisRepository) GetCustomerToken(customerID string) (string, error) {
	value, err := r.client.Get("customer:" + customerID).Result()
	if err != nil {
		return "", err
	}
	return value, nil
}

// NewRedisRepository contructs new redis client with given info.
func NewRedisRepository(addr string, password string, db int) Repository {
	repository := new(RedisRepository)
	repository.client = redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password,
		DB:       db,
	})
	return repository
}
