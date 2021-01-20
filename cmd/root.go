package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "hoge",
	Short: "fuga",
	Long:  "hogefuga",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("root command")

	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
