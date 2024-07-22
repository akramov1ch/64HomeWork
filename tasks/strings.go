package tasks

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
)

func StringsExample(client *redis.Client) error {
	ctx := context.Background()

	err := client.Set(ctx, "key", "value", 0).Err()
	if err != nil {
		return err
	}

	val, err := client.Get(ctx, "key").Result()
	if err != nil {
		return err
	}
	fmt.Println("Key:", val)

	err = client.Set(ctx, "counter", 0, 0).Err()
	if err != nil {
		return err
	}

	err = client.Incr(ctx, "counter").Err()
	if err != nil {
		return err
	}

	counter, err := client.Get(ctx, "counter").Result()
	if err != nil {
		return err
	}
	fmt.Println("Counter:", counter)

	return nil
}
