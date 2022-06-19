package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// getCmd represents the get command
var getCmd = &cobra.Command{
	Use:                "get",
	Short:              "Gets the value of a key",
	Long:               `Gets the value of a key`,
	Example:            "rcli get name",
	DisableFlagParsing: false,
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx, rdb := RedisConfiguration()
		ArgsNumberCheck(1, len(args))
		val, err := rdb.Get(ctx, args[0]).Result()
		if err != nil {
			return err
		}
		fmt.Println(val)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(getCmd)
}
