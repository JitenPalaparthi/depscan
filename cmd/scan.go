package cmd

import (
	"errors"
	"io/fs"
	"log"
	"path/filepath"

	"github.com/JitenPalaparthi/depscan/config"
	"github.com/golang/glog"
	"github.com/spf13/cobra"
)

var (
	path, format, outFile string
	depth                 uint8
	ErrSkipFile           = errors.New("skip file")
)

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
		fcount, dcount := 0, 0
		glog.Infoln("Current path is ", path)
		cnfg, err := config.New()
		log.Println(cnfg, err)
		filepath.WalkDir(path, func(p string, d fs.DirEntry, err error) error {
			if d.IsDir() {
				if isElementExist(cnfg.IgnoreDirs, d.Name()) {
					return filepath.SkipDir
				}
				dcount++
			} else {
				if isElementExist(cnfg.IgnoreFiles, d.Name()) {
					return ErrSkipFile
				}
				fcount++
			}
			log.Println(p)
			return nil
		})

	},
}

func isElementExist(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}
	return false
}
