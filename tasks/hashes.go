package tasks

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
)

func HashesExample(client *redis.Client) error {
	ctx := context.Background()

	userKey := "user:1000"
	userData := map[string]interface{}{
		"id":    "1000",
		"name":  "Shaxboz",
		"email": "shaxboz@gmail.com",
	}
	err := client.HMSet(ctx, userKey, userData).Err()
	if err != nil {
		return err
	}

	err = client.HMSet(ctx, userKey, map[string]interface{}{
		"name":  "Abduazim",
		"email": "abduazim@gmail.com",
	}).Err()
	if err != nil {
		return err
	}

	fields, err := client.HGetAll(ctx, userKey).Result()
	if err != nil {
		return err
	}
	fmt.Println("User fields:", fields)

	return nil
}
