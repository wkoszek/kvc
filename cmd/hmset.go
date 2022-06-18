package cmd

import (
	"log"

	"github.com/spf13/cobra"
)

// hmsetCmd represents the hmset command
var hmsetCmd = &cobra.Command{
	Use:     "hmset",
	Short:   "Sets multiple hash fields to multiple values",
	Long:    `Sets multiple hash fields to multiple values`,
	Example: "rcli hmset user name John",
	Run: func(cmd *cobra.Command, args []string) {
		ctx, rdb := RedisConfiguration()
		ArgsNumberCheck(2, len(args))
		err := rdb.HMSet(ctx, args[0], args[1:]).Err()
		if err != nil {
			log.Println(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(hmsetCmd)
}
