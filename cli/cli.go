package cli

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/ramin0x53/rcli/config"
)

type Options struct {
	Command string
	Key     string
	Values  []string
}

//List of all defined commands
var Commands = []string{"set", "get", "del", "lpop", "lpush", "lrange", "hmset", "hgetall", "hget", "hdel", "configcl"}

func commandCheck(cmd string) bool {
	for _, v := range Commands {
		if cmd == v {
			return true
		}
	}
	return false
}

func help() {
	sp := strings.Repeat(" ", 5)
	midsp := 15
	fmt.Println("Commands:")
	fmt.Println(sp, "configcl", strings.Repeat(" ", midsp-len("configcl")), "rcli configcl [ip:port] [password] [db(integer)]")
	fmt.Println(sp, "set", strings.Repeat(" ", midsp-len("set")), "rcli set key value")
	fmt.Println(sp, "get", strings.Repeat(" ", midsp-len("get")), "rcli get key")
	fmt.Println(sp, "del", strings.Repeat(" ", midsp-len("del")), "rcli del key")
	fmt.Println(sp, "lpop", strings.Repeat(" ", midsp-len("lpop")), "rcli lpop key")
	fmt.Println(sp, "lpush", strings.Repeat(" ", midsp-len("lpush")), "rcli lpush key value")
	fmt.Println(sp, "lrange", strings.Repeat(" ", midsp-len("lrange")), "rcli lrange key start stop")
	fmt.Println(sp, "hmset", strings.Repeat(" ", midsp-len("hmset")), "rcli hmset key field value")
	fmt.Println(sp, "hgetall", strings.Repeat(" ", midsp-len("hgetall")), "rcli hgetall key")
	fmt.Println(sp, "hget", strings.Repeat(" ", midsp-len("hget")), "rcli hget key field")
	fmt.Println(sp, "hdel", strings.Repeat(" ", midsp-len("hdel")), "rcli hdel key field")
}

func GetOptions() *Options {
	opts := Options{}

	if len(os.Args) <= 1 {
		fmt.Println("Error: Command arguments are empty")
		fmt.Println("Use \"rcli help\" for more information")
		os.Exit(0)
	}

	opts.Command = strings.ToLower(os.Args[1])

	if opts.Command == "configcl" {
		if len(os.Args) == 4 {
			config.WriteConfig(os.Args[2], os.Args[3], "0")
		} else if len(os.Args) == 3 {
			config.WriteConfig(os.Args[2], "", "0")
		} else if len(os.Args) == 5 {
			config.WriteConfig(os.Args[2], os.Args[3], os.Args[4])
		} else {
			log.Println("Wrong arguments for configcl")
		}

		os.Exit(0)
	}

	if opts.Command == "help" {
		help()
		os.Exit(0)
	}

	if !config.CheckConfig() {
		fmt.Println("First, enter the server IP and Password with the following command:")
		fmt.Println("    rcli configcl [ip:port] [password] [db(integer)]")
		os.Exit(0)
	}

	if !commandCheck(opts.Command) {
		fmt.Printf("Error: The \"%s\" command is not defined\n", opts.Command)
		fmt.Println("Use \"rcli help\" for more information")
		os.Exit(0)
	}

	opts.Key = os.Args[2]
	opts.Values = os.Args[3:]
	return &opts
}
