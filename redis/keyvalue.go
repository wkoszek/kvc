package redis

import (
	"context"
	"fmt"
	"log"

	"github.com/go-redis/redis/v9"
	"github.com/ramin0x53/rcli/cli"
)

//Implementation of get command
func GetRedis(ctx context.Context, rdb *redis.Client, data *cli.Options) {

	val, err := rdb.Get(ctx, data.Key).Result()
	if err != nil {
		log.Println(err)
	}
	fmt.Println(val)

}

//Implementation of set command
func SetRedis(ctx context.Context, rdb *redis.Client, data *cli.Options) {
	err := rdb.Set(ctx, data.Key, data.Values[0], 0).Err()
	if err != nil {
		log.Println(err)
	}
}

//Implementation of del command
func DelRedis(ctx context.Context, rdb *redis.Client, data *cli.Options) {
	err := rdb.Del(ctx, data.Key).Err()
	if err != nil {
		log.Println(err)
	}
}
