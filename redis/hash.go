package redis

import (
	"context"
	"fmt"
	"log"

	"github.com/go-redis/redis/v9"
	"github.com/ramin0x53/rcli/cli"
)

//Implementation of hmset command
func HmsetRedis(ctx context.Context, rdb *redis.Client, data *cli.Options) {
	err := rdb.HMSet(ctx, data.Key, data.Values).Err()
	if err != nil {
		log.Println(err)
	}
}

//Implementation of hgetall command
func HgetallRedis(ctx context.Context, rdb *redis.Client, data *cli.Options) {

	val, err := rdb.HGetAll(ctx, data.Key).Result()
	if err != nil {
		log.Println(err)
	}

	for key, element := range val {
		fmt.Println(key, ": ", element)
	}

}

//Implementation of hget command
func HgetRedis(ctx context.Context, rdb *redis.Client, data *cli.Options) {

	val, err := rdb.HGet(ctx, data.Key, data.Values[0]).Result()
	if err != nil {
		log.Println(err)
	}
	fmt.Println(val)

}

//Implementation of hdel command
func HdelRedis(ctx context.Context, rdb *redis.Client, data *cli.Options) {
	err := rdb.HDel(ctx, data.Key, data.Values...).Err()
	if err != nil {
		log.Println(err)
	}
}
