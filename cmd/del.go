package cmd

import (
	"log"

	"github.com/spf13/cobra"
)

// delCmd represents the del command
var delCmd = &cobra.Command{
	Use:     "del",
	Short:   "This command deletes the key, if it exists",
	Long:    `This command deletes the key, if it exists`,
	Example: "rcli del name",
	Run: func(cmd *cobra.Command, args []string) {
		ctx, rdb := RedisConfiguration()
		ArgsNumberCheck(1, len(args))
		err := rdb.Del(ctx, args...).Err()
		if err != nil {
			log.Println(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(delCmd)
}
