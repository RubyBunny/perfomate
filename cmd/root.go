package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{Use: "perfomate"}

var InputPath string
var OutputPath string

func init() {
	rootCmd.AddCommand(perfomanceCmd)
	rootCmd.AddCommand(selfCmd)

	rootCmd.PersistentFlags().StringVarP(&InputPath, "input-path", "i", "", "input raw file path")
	rootCmd.PersistentFlags().StringVarP(&OutputPath, "output-path", "o", "./", "output path (application doesn't create folders)")

	rootCmd.MarkPersistentFlagRequired("input-path")
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
	}
}
