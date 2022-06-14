package main

import (
	"github.com/ramin0x53/rcli/cli"
	"github.com/ramin0x53/rcli/redis"
)

func main() {
	args := cli.GetOptions()
	redis.RedisRunner(args)
}
