package main

import (
	"context"

	"github.com/redis/go-redis/v9"
)

func main() {
	ctx := context.Background()

	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	err := rdb.FlushDB(ctx).Err()
	if err != nil {
		panic(err)
	}

	err = rdb.Set(ctx, "Hello", "World", 0).Err()
	if err != nil {
		panic(err)
	}

	err = rdb.HSet(ctx, "user1", "name", "Bob", "Age", 25).Err()
	if err != nil {
		panic(err)
	}
}
