package cmd

import (
	"github.com/spf13/cobra"
)

// hdelCmd represents the hdel command
var hdelCmd = &cobra.Command{
	Use:     "hdel",
	Short:   "Deletes one or more hash fields",
	Long:    `Deletes one or more hash fields`,
	Example: "rcli hdel user name",
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx, rdb := RedisConfiguration()
		ArgsNumberCheck(2, len(args))
		err := rdb.HDel(ctx, args[0], args[1:]...).Err()
		if err != nil {
			return err
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(hdelCmd)
}
