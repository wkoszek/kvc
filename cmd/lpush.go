package cmd

import (
	"github.com/spf13/cobra"
)

// lpushCmd represents the lpush command
var lpushCmd = &cobra.Command{
	Use:     "lpush",
	Short:   "Prepends a value to a list",
	Long:    `Prepends a value to a list`,
	Example: "rcli lpush names John",
	Args:    cobra.MinimumNArgs(2),
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx, rdb := RedisConfiguration()
		err := rdb.LPush(ctx, args[0], args[1:]).Err()
		if err != nil {
			return err
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(lpushCmd)
}
