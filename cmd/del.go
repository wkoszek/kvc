package cmd

import (
	"github.com/spf13/cobra"
)

// delCmd represents the del command
var delCmd = &cobra.Command{
	Use:     "del",
	Short:   "This command deletes the key, if it exists",
	Long:    `This command deletes the key, if it exists`,
	Example: "rcli del name",
	Args:    cobra.MinimumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx, rdb := RedisConfiguration()
		err := rdb.Del(ctx, args...).Err()
		if err != nil {
			return err
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(delCmd)
}
