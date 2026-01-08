package utils

import (
	"context"
	"fmt"
	"os"
	"github.com/go-redis/redis/v8"
)

var RedisClient *redis.Client 
var Ctx = context.Background()

func InitRedis(){
	redisHost := os.Getenv("REDIS_HOST")
	if redisHost == " "{
		redisHost = "localhost"
	}
	RedisClient = redis.NewClient(&redis.Options{
		Addr :redisHost, //default port 6379
		Password : "", // no password set
		DB : 0,
	})

	_, err := RedisClient.Ping(Ctx).Result()
	if err != nil {
		panic(fmt.Sprintf("Failed to connect to Redis: %v", err))
	}
	fmt.Println("Connected to Redis successfully")
}