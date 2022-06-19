package cmd

import (
	"github.com/spf13/cobra"
)

// setCmd represents the set command
var setCmd = &cobra.Command{
	Use:     "set",
	Short:   "This command sets the value at the specified key",
	Long:    `This command sets the value at the specified key`,
	Example: "rcli set name John",
	Args:    cobra.ExactArgs(2),
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx, rdb := RedisConfiguration()
		err := rdb.Set(ctx, args[0], args[1], 0).Err()
		if err != nil {
			return err
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(setCmd)
}
