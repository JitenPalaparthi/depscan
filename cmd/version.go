package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "depscan v0.0.4",
	Long:  `depscan v0.0.4`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("depscan  v0.0.4")
	},
}
