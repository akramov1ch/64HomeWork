package main

import (
	"fmt"
	"log"

	"64HW/tasks"
)

func main() {
	InitRedis()
	client := GetRedisClient()
	defer client.Close()

	if err := tasks.StringsExample(client); err != nil {
		log.Fatalf("Strings example failed: %v", err)
	}
	fmt.Println("Strings example completed successfully")

	if err := tasks.HashesExample(client); err != nil {
		log.Fatalf("Hashes example failed: %v", err)
	}
	fmt.Println("Hashes example completed successfully")

	if err := tasks.ListsExample(client); err != nil {
		log.Fatalf("Lists example failed: %v", err)
	}
	fmt.Println("Lists example completed successfully")
}
