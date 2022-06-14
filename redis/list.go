package redis

import (
	"context"
	"fmt"
	"log"

	"github.com/go-redis/redis/v9"
	"github.com/ramin0x53/rcli/cli"
)

//Implementation of lpush command
func LpushRedis(ctx context.Context, rdb *redis.Client, data *cli.Options) {
	err := rdb.LPush(ctx, data.Key, data.Values[0]).Err()
	if err != nil {
		log.Println(err)
	}
}

//Implementation of lpop command
func LpopRedis(ctx context.Context, rdb *redis.Client, data *cli.Options) {
	err := rdb.LPop(ctx, data.Key).Err()
	if err != nil {
		log.Println(err)
	}
}

//Implementation of lrange command
func LrangeRedis(ctx context.Context, rdb *redis.Client, data *cli.Options) {

	val, err := rdb.LRange(ctx, data.Key, int64(IntConv(data.Values[0])), int64(IntConv(data.Values[1]))).Result()
	if err != nil {
		log.Println(err)
	}

	for _, txt := range val {
		fmt.Println(txt)
	}

}
