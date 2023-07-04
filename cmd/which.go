package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(whichCmd)
}

var whichCmd = &cobra.Command{
	Use:   "which",
	Short: "currently it supports Java,Python,JavaScript",
	Long:  `currently it depscan supports Java(maven|gradle),Python(pip),JavaScript(npm->package-lock.json version 1 and 2)`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(`currently it depscan supports Java(maven|gradle),Python(pip),JavaScript(npm->package-lock.json version 1 and 2)`)
	},
}
