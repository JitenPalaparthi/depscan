package cmd

import (
	"errors"
	"fmt"

	"log"
	"path/filepath"

	"github.com/JitenPalaparthi/depscan/config"
	"github.com/JitenPalaparthi/depscan/implement"
	scnr "github.com/JitenPalaparthi/depscan/scanner"
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
	scanCmd.Flags().Uint8VarP(&depth, "depth", "d", 3, "the depth of directory recursion for file scans")
	scanCmd.Flags().StringVarP(&outFile, "out", "o", "output", "user has to provide output file name")

	rootCmd.AddCommand(scanCmd)
}

var scanCmd = &cobra.Command{
	Use:   "scan",
	Short: "scan scans a given repository",
	Long:  "scan scans a given repository provided by given path",
	Run: func(cmd *cobra.Command, args []string) {
		glog.Infoln("Current path is ", path)
		cnfg, err := config.New()
		if err != nil {
			log.Fatalln(err)
		}
		impl, err := implement.New(cnfg, path, fmt.Sprint(outFile, ".", format), depth) // create an instance of implement
		if err != nil {
			log.Fatalln(err)
		}
		fmt.Println(impl)
		err = impl.Feed() // feed required data for the implement object3
		if err != nil {
			log.Fatalln(err)
		}
		iscanners := make([]scnr.Scanner, 0)
		for _, v := range impl.Paths {
			depManager := cnfg.GetDepManagerByFileName(filepath.Base(v))
			if depManager != nil && isElementExist(impl.Exts, depManager.FileExt) {
				switch depManager.DepTool {
				case "pip":
					pip := new(implement.Pip)
					pip.FilePath = v
					iscanners = append(iscanners, pip)
				case "npm":
					npm := new(implement.Npm)
					npm.FilePath = v
					iscanners = append(iscanners, npm)
				case "gradle":
					gradle := new(implement.Gradle)
					gradle.FilePath = v
					iscanners = append(iscanners, gradle)
				case "maven":
					maven := new(implement.Maven)
					maven.FilePath = v
					iscanners = append(iscanners, maven)
				default:
					log.Println("Unimplemented tool")
				}
			}
		}
		deps, err := impl.ScanAll(iscanners...)
		if err != nil {
			log.Fatalln(err)
		}
		err = impl.Write(deps)
		if err != nil {
			log.Fatalln(err)
		}

		fmt.Println("Directory count", impl.DirCount, "\nFile Count", impl.FileCount)
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
