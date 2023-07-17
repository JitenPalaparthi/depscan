package cmd

import (
	"errors"
	"fmt"
	"strings"

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
			glog.Errorln(err)
			return
		}

		glog.Infoln("config object:", cnfg)
		// What if outfile has format?
		impl, err := implement.New(cnfg, path, fmt.Sprint(strings.TrimSuffix(outFile, filepath.Ext(outFile)), ".", format), depth) // create an instance of implement
		if err != nil {
			glog.Errorln(err)
			return
		}
		glog.Infoln("implement object:", impl)

		err = impl.Feed() // feed required data for the implement object3
		if err != nil {
			glog.Errorln(err)
			return
		}

		// fmt.Println("<><>><><><><><<><><")
		// fmt.Println(impl.PathSets)
		// fmt.Println("<><><><><>>><><><<<>")

		glog.Infoln("implement object after feed:", impl)

		iscanners := make([]scnr.Scanner, 0)
		var (
			pip    *implement.Pip
			npm    *implement.Npm
			gradle *implement.Gradle
			maven  *implement.Maven
		)
		for _, value := range impl.PathSets {
			depManager := cnfg.GetDepManagerByFileName(filepath.Base(value[0]))
			// The problem is .java files are inside many directories.So the tool might have not found
			//.java file. Unless it finds a file with .java or .py or .js extention it cannot blindly go and
			// do the process.Due to .java file directory depth , it is unable to find for gradle and maven.
			// Hence the below logic is commented.
			//if depManager != nil && helper.IsElementExist(impl.Exts, depManager.FileExt) {
			if depManager != nil {
				switch depManager.DepTool {
				case "pip":
					pip = new(implement.Pip)
					pip.FilePaths = append(pip.FilePaths, value...)
					iscanners = append(iscanners, pip)
					glog.Infoln("Found pip as dependency manager.The Filepath is ", value[0])

				case "npm":
					npm = new(implement.Npm)
					npm.FilePaths = append(npm.FilePaths, value...)
					iscanners = append(iscanners, npm)
					glog.Infoln("Found npm as dependency manager.The Filepath is ", value[0])

				case "gradle":
					gradle = new(implement.Gradle)
					gradle.FilePaths = append(gradle.FilePaths, value...)
					iscanners = append(iscanners, gradle)
					glog.Infoln("Found gradle as dependency manager.The Filepath is ", value[0])

				case "maven":
					maven = new(implement.Maven)
					maven.FilePaths = append(maven.FilePaths, value...)
					iscanners = append(iscanners, maven)
					glog.Infoln("Found maven as dependency manager.The Filepath is ", value[0])

				default:
					glog.Infoln("Unimplemented tool")
				}
			}
		}
		// for _, v := range impl.Paths {
		// 	depManager := cnfg.GetDepManagerByFileName(filepath.Base(v))
		// 	// The problem is .java files are inside many directories.So the tool might have not found
		// 	//.java file. Unless it finds a file with .java or .py or .js extention it cannot blindly go and
		// 	// do the process.Due to .java file directory depth , it is unable to find for gradle and maven.
		// 	// Hence the below logic is commented.
		// 	//if depManager != nil && helper.IsElementExist(impl.Exts, depManager.FileExt) {
		// 	if depManager != nil {
		// 		switch depManager.DepTool {
		// 		case "pip":
		// 			if pip == nil {
		// 				pip = new(implement.Pip)
		// 			}
		// 			pip.FilePaths = append(pip.FilePaths, v)
		// 			iscanners = append(iscanners, pip)
		// 			glog.Infoln("Found pip as dependency manager.The Filepath is ", v)

		// 		case "npm":
		// 			if npm == nil {
		// 				npm = new(implement.Npm)
		// 			}
		// 			npm.FilePaths = append(npm.FilePaths, v)
		// 			iscanners = append(iscanners, npm)
		// 			glog.Infoln("Found npm as dependency manager.The Filepath is ", v)

		// 		case "gradle":
		// 			if gradle == nil {
		// 				gradle = new(implement.Gradle)
		// 			}
		// 			gradle.FilePaths = append(gradle.FilePaths, v)
		// 			iscanners = append(iscanners, gradle)
		// 			glog.Infoln("Found gradle as dependency manager.The Filepath is ", v)

		// 		case "maven":
		// 			if maven == nil {
		// 				maven = new(implement.Maven)
		// 			}
		// 			maven.FilePaths = append(maven.FilePaths, v)
		// 			iscanners = append(iscanners, maven)
		// 			glog.Infoln("Found maven as dependency manager.The Filepath is ", v)

		// 		default:
		// 			glog.Infoln("Unimplemented tool")
		// 		}
		// 	}
		// }
		glog.Infoln("There are/is ", len(iscanners), "of scanners to scan")
		deps, err := impl.ScanAll(iscanners...)
		if err != nil {
			glog.Errorln(err)
			return
		}
		err = impl.Write(deps)
		if err != nil {
			glog.Errorln(err)
			return
		}
		glog.Info("Directory count:", impl.DirCount, "\nFile Count:", impl.FileCount)
	},
}
