package cmd

import (
	"log"

	"github.com/spf13/cobra"
)

// hdelCmd represents the hdel command
var hdelCmd = &cobra.Command{
	Use:     "hdel",
	Short:   "Deletes one or more hash fields",
	Long:    `Deletes one or more hash fields`,
	Example: "rcli hdel user name",
	Run: func(cmd *cobra.Command, args []string) {
		ctx, rdb := RedisConfiguration()
		ArgsCountCheck(2, len(args))
		err := rdb.HDel(ctx, args[0], args[1:]...).Err()
		if err != nil {
			log.Println(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(hdelCmd)
}
