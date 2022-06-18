package cmd

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
)

// getCmd represents the get command
var getCmd = &cobra.Command{
	Use:                "get",
	Short:              "Gets the value of a key",
	Long:               `Gets the value of a key`,
	Example:            "rcli get name",
	DisableFlagParsing: false,
	Run: func(cmd *cobra.Command, args []string) {
		ctx, rdb := RedisConfiguration()
		ArgsCountCheck(1, len(args))
		val, err := rdb.Get(ctx, args[0]).Result()
		if err != nil {
			log.Println(err)
		}
		fmt.Println(val)
	},
}

func init() {
	rootCmd.AddCommand(getCmd)
}
