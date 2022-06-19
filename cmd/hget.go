package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// hgetCmd represents the hget command
var hgetCmd = &cobra.Command{
	Use:     "hget",
	Short:   "Gets the value of a hash field stored at the specified key",
	Long:    `Gets the value of a hash field stored at the specified key`,
	Example: "rcli hget user name",
	Args:    cobra.ExactArgs(2),
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx, rdb := RedisConfiguration()
		val, err := rdb.HGet(ctx, args[0], args[1]).Result()
		if err != nil {
			return err
		}
		fmt.Println(val)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(hgetCmd)
}
