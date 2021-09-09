package redisclient

import (
	"context"
	"fmt"

	"github.com/go-redis/redis/v8"
)

// var (
// 	ctx context.Context
// 	rdb *redis.Client
// )

type RedisClient struct {
	ctx context.Context
	rdb *redis.Client
}

//init Redis connection
func (client *RedisClient) Init(addr string, ctx context.Context) *RedisClient {
	if len(addr) == 0 {
		addr = "localhost:6379"
		fmt.Println("use default addr :", addr)
	}
	rdb := redis.NewClient(&redis.Options{
		Addr:     addr, //localhost:6379
		Password: "",   // no password set
		DB:       0,    // use default DB
	})
	client.ctx = ctx
	client.rdb = rdb
	return client
}

//set redis key value
func (client *RedisClient) Set(k string, v string) {
	err := client.rdb.Set(client.ctx, k, v, 0).Err()
	if err != nil {
		panic(err)
	}
}

//get redis value by key
func (client *RedisClient) Get(k string) string {
	v, err := client.rdb.Get(client.ctx, k).Result()
	if err == redis.Nil {
		fmt.Println("key ", k, " does not exist")
	} else if err != nil {
		panic(err)
	} else {
		fmt.Println(k, ":", v)
	}
	return v
}

func (client *RedisClient) HSet(key string, data ...interface{}) {

}
