package redis

import (
	"context"
	"log"
	"strconv"

	"github.com/go-redis/redis/v9"
	"github.com/ramin0x53/rcli/cli"
	"github.com/ramin0x53/rcli/config"
)

//Convert string to integer
func IntConv(text string) int {
	s, err := strconv.Atoi(text)
	if err != nil {
		log.Println(err)
	}
	return s
}

func RedisRunner(data *cli.Options) {

	configData := config.ReadConfig()
	rdb := redis.NewClient(&redis.Options{
		Addr:     configData.Ip,
		Password: configData.Password,
		DB:       IntConv(configData.Db),
	})
	ctx := context.Background()

	switch data.Command {
	case "get":
		GetRedis(ctx, rdb, data)
	case "set":
		SetRedis(ctx, rdb, data)
	case "del":
		DelRedis(ctx, rdb, data)
	case "lpop":
		LpopRedis(ctx, rdb, data)
	case "lpush":
		LpushRedis(ctx, rdb, data)
	case "lrange":
		LrangeRedis(ctx, rdb, data)
	case "hmset":
		HmsetRedis(ctx, rdb, data)
	case "hgetall":
		HgetallRedis(ctx, rdb, data)
	case "hget":
		HgetRedis(ctx, rdb, data)
	case "hdel":
		HdelRedis(ctx, rdb, data)
	default:
		log.Println("Wrong command")
	}
}
