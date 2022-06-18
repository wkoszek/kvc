package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// lrangeCmd represents the lrange command
var lrangeCmd = &cobra.Command{
	Use:     "lrange",
	Short:   "Gets a range of elements from a list",
	Long:    `Gets a range of elements from a list`,
	Example: "rcli lrange names 0 10",
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx, rdb := RedisConfiguration()
		ArgsNumberCheck(3, len(args))
		val, err := rdb.LRange(ctx, args[0], int64(IntConv(args[1])), int64(IntConv(args[2]))).Result()
		if err != nil {
			return err
		}

		for _, txt := range val {
			fmt.Println(txt)
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(lrangeCmd)
}
