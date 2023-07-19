package cmd

import (
	"fmt"
	"os"

	glogcobra "github.com/blocktop/go-glog-cobra"
	"github.com/spf13/cobra"
)

func init() {
	cobra.OnInitialize(initConfig)

	// rootCmd is defined by standard cobra configuration
	glogcobra.Init(rootCmd)
}

// initConfig is part of the standard rootCmd configuration that
// cobra creates.
func initConfig() {
	// This will also call flag.Parse() if you have not already.
	glogcobra.Parse(rootCmd)

	// glog will now have all the flags it needs
}

var rootCmd = &cobra.Command{
	Use:   "depscan",
	Short: "depscan is a dependency scanner",
	Long:  `A dependency scanner that scans repositories that are developed using different programming languages.`,
	Run: func(cmd *cobra.Command, args []string) {
		// Do Stuff Here
		fmt.Println("A dependency scanner that scans repositories that are developed using different programming languages.\nCurrently it supports Java(maven | gradle), Python(pip)\ndepscan  v0.0.5")
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
