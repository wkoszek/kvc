package cmd

import (
	"github.com/spf13/cobra"
)

// hmsetCmd represents the hmset command
var hmsetCmd = &cobra.Command{
	Use:     "hmset",
	Short:   "Sets multiple hash fields to multiple values",
	Long:    `Sets multiple hash fields to multiple values`,
	Example: "rcli hmset user name John",
	Args:    cobra.MinimumNArgs(2),
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx, rdb := RedisConfiguration()
		err := rdb.HMSet(ctx, args[0], args[1:]).Err()
		if err != nil {
			return err
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(hmsetCmd)
}
