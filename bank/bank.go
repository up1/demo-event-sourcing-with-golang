package bank

import (
	"errors"
	"os"
	"strconv"

	"github.com/go-redis/redis"
)

type BankAccount struct {
	Id      string
	Name    string
	Balance int
}

func initRedis() *redis.Client {
	redisUrl := os.Getenv("REDIS_URL")

	if redisUrl == "" {
		redisUrl = "redis:6379"
	}

	return redis.NewClient(&redis.Options{
		Addr:     redisUrl,
		Password: "",
		DB:       0,
	})
}

func FetchAccount(id string) (*BankAccount, error) {
	cmd := initRedis().HGetAll(id)
	if err := cmd.Err(); err != nil {
		return nil, err
	}

	data := cmd.Val()
	if len(data) == 0 {
		return nil, nil
	} else {
		return ToAccount(data)
	}
}

func UpdateAccount(id string, data map[string]interface{}) (*BankAccount, error) {
	cmd := initRedis().HMSet(id, data)

	if err := cmd.Err(); err != nil {
		return nil, err
	} else {
		return FetchAccount(id)
	}
}

func ToAccount(m map[string]string) (*BankAccount, error) {
	balance, err := strconv.Atoi(m["Balance"])
	if err != nil {
		return nil, err
	}

	if _, ok := m["Id"]; !ok {
		return nil, errors.New("Missing account ID")
	}

	return &BankAccount{
		Id:      m["Id"],
		Name:    m["Name"],
		Balance: balance,
	}, nil
}
