package cmd

import (
	"github.com/spf13/cobra"
)

var path string
var format string
var depth uint8
var outFile string

func init() {

	scanCmd.Flags().StringVarP(&path, "path", "p", ".", "user has to provide path.Ideally this is a git repository path")
	scanCmd.Flags().StringVarP(&format, "format", "f", "json", "output file format. We support two formats json|yaml")
	scanCmd.Flags().Uint8VarP(&depth, "depth", "d", 1, "the depth of directory recursion for file scans")
	scanCmd.Flags().StringVarP(&outFile, "out", "o", "output.json", "user has to provide output file name")

	rootCmd.AddCommand(scanCmd)
}

var scanCmd = &cobra.Command{
	Use:   "scan",
	Short: "scan scans a given repository",
	Long:  "scan scans a given repository provided by given path",
	Run: func(cmd *cobra.Command, args []string) {
	},
}
