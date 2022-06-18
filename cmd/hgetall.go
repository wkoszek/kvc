package cmd

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
)

// hgetallCmd represents the hgetall command
var hgetallCmd = &cobra.Command{
	Use:     "hgetall",
	Short:   "Gets all the fields and values stored in a hash at the specified key",
	Long:    `Gets all the fields and values stored in a hash at the specified key`,
	Example: "rcli hgetall user",
	Run: func(cmd *cobra.Command, args []string) {
		ctx, rdb := RedisConfiguration()
		ArgsCountCheck(1, len(args))
		val, err := rdb.HGetAll(ctx, args[0]).Result()
		if err != nil {
			log.Println(err)
		}

		for key, element := range val {
			fmt.Println(key, ": ", element)
		}
	},
}

func init() {
	rootCmd.AddCommand(hgetallCmd)
}
