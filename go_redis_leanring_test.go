package main

import (
	"fmt"
	"testing"

	"github.com/go-redis/redis/v7"
)

func TestNewClient(t *testing.T) {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	pong, err := client.Ping().Result()
	fmt.Println(pong, err)
	if err != nil {
		t.Errorf("An error occurred: %s\n", err.Error())
	}
}

func TestClient(t *testing.T) {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
	err := client.Set("someKey", "value", 0).Err()

	if err != nil {
		t.Errorf("Error occurred: %s\n", err.Error())
	}

	val, err := client.Get("someKey").Result()
	if err != nil {
		t.Errorf("Error occurred: %s\n", err.Error())
	}

	t.Log("someKey", val)

	val2, err := client.Get("otherKey").Result()
	if err == redis.Nil {
		t.Log("otherKey does not exist")
	} else if err == nil {
		t.Errorf("Error occurred: %s\n", err.Error())
	}
	t.Log("otherKey", val2)
}

func TestAppend(t *testing.T) {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
	client.Set("someValue2", "2", 0)
	client.Append("someValue2", "1")
	val := client.Get("someValue2")
	t.Log("Appended", val)
}
