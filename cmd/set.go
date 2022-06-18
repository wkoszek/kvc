package cmd

import (
	"log"

	"github.com/spf13/cobra"
)

// setCmd represents the set command
var setCmd = &cobra.Command{
	Use:     "set",
	Short:   "This command sets the value at the specified key",
	Long:    `This command sets the value at the specified key`,
	Example: "rcli set name John",
	Run: func(cmd *cobra.Command, args []string) {
		ctx, rdb := RedisConfiguration()
		err := rdb.Set(ctx, args[0], args[1], 0).Err()
		if err != nil {
			log.Println(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(setCmd)
}
