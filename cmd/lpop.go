package cmd

import (
	"log"

	"github.com/spf13/cobra"
)

// lpopCmd represents the lpop command
var lpopCmd = &cobra.Command{
	Use:     "lpop",
	Short:   "Removes and gets the first element in a list",
	Long:    `Removes and gets the first element in a list`,
	Example: "rcli lpop names",
	Run: func(cmd *cobra.Command, args []string) {
		ctx, rdb := RedisConfiguration()
		ArgsNumberCheck(1, len(args))
		err := rdb.LPop(ctx, args[0]).Err()
		if err != nil {
			log.Println(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(lpopCmd)
}
