package tasks

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
)

func ListsExample(client *redis.Client) error {
	ctx := context.Background()

	listKey := "mylist"
	elements := []string{"one", "two", "three"}
	for _, el := range elements {
		err := client.LPush(ctx, listKey, el).Err()
		if err != nil {
			return err
		}
	}

	list, err := client.LRange(ctx, listKey, 0, -1).Result()
	if err != nil {
		return err
	}
	fmt.Println("List elements:", list)

	err = client.LRem(ctx, listKey, 0, "two").Err()
	if err != nil {
		return err
	}

	list, err = client.LRange(ctx, listKey, 0, -1).Result()
	if err != nil {
		return err
	}
	fmt.Println("List elements after removal:", list)

	return nil
}
