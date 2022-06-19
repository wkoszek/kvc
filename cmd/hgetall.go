package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// hgetallCmd represents the hgetall command
var hgetallCmd = &cobra.Command{
	Use:     "hgetall",
	Short:   "Gets all the fields and values stored in a hash at the specified key",
	Long:    `Gets all the fields and values stored in a hash at the specified key`,
	Example: "rcli hgetall user",
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx, rdb := RedisConfiguration()
		ArgsNumberCheck(1, len(args))
		val, err := rdb.HGetAll(ctx, args[0]).Result()
		if err != nil {
			return err
		}

		for key, element := range val {
			fmt.Println(key, ": ", element)
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(hgetallCmd)
}
