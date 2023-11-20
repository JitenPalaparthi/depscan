package cmd

import (
	"errors"
	"fmt"
	"strings"

	"path/filepath"

	"github.com/JitenPalaparthi/depscan/config"
	"github.com/JitenPalaparthi/depscan/implement"
	composerp "github.com/JitenPalaparthi/depscan/implement/composer"
	gomodp "github.com/JitenPalaparthi/depscan/implement/gomod"
	gradlep "github.com/JitenPalaparthi/depscan/implement/gradle"
	mavenp "github.com/JitenPalaparthi/depscan/implement/maven"
	npmp "github.com/JitenPalaparthi/depscan/implement/npm"
	nugetp "github.com/JitenPalaparthi/depscan/implement/nuget"
	pipp "github.com/JitenPalaparthi/depscan/implement/pip"
	pipenvp "github.com/JitenPalaparthi/depscan/implement/pipenv"
	poetryp "github.com/JitenPalaparthi/depscan/implement/poetry"

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
	scanCmd.Flags().StringVarP(&outFile, "out", "o", "", "user has to provide output file name")

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
		var (
			impl *implement.Implement
		)

		// if not output file is given , there need not to be any output as a file, print it int the stdout
		if outFile == "" {
			format = ""
			impl, err = implement.New(cnfg, path, "", depth) // create an instance of implement
			if err != nil {
				glog.Errorln(err)
				return
			}
			glog.Infoln("implement object:", impl)

		} else {
			impl, err = implement.New(cnfg, path, fmt.Sprint(strings.TrimSuffix(outFile, filepath.Ext(outFile)), ".", format), depth) // create an instance of implement
			if err != nil {
				glog.Errorln(err)
				return
			}
			glog.Infoln("implement object:", impl)
		}

		err = impl.Feed() // Feed method feeds required data for the implement object3
		if err != nil {
			glog.Errorln(err)
			return
		}

		glog.Infoln("implement object after feed:", impl)

		iscanners := make([]scnr.Scanner, 0)
		var (
			pip      *pipp.Pip
			pipenv   *pipenvp.Pipenv
			poetry   *poetryp.Poetry
			npm      *npmp.Npm
			gradle   *gradlep.Gradle
			maven    *mavenp.Maven
			composer *composerp.Composer
			gomod    *gomodp.Go
			nuget    *nugetp.Nuget
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
				case "pypi":
					pip = new(pipp.Pip)
					pip.FilePaths = append(pip.FilePaths, value...)
					iscanners = append(iscanners, pip)
					glog.Infoln("Found pip as dependency manager.The Filepath is ", value[0])
				case "pipenv":
					pipenv = new(pipenvp.Pipenv)
					pipenv.FilePaths = append(pipenv.FilePaths, value...)
					iscanners = append(iscanners, pipenv)
					glog.Infoln("Found pipenv as dependency manager.The Filepath is ", value[0])

				case "poetry":
					poetry = new(poetryp.Poetry)
					poetry.FilePaths = append(poetry.FilePaths, value...)
					iscanners = append(iscanners, poetry)
					glog.Infoln("Found poetry as dependency manager.The Filepath is ", value[0])

				case "npm":
					npm = new(npmp.Npm)
					npm.FilePaths = append(npm.FilePaths, value...)
					iscanners = append(iscanners, npm)
					glog.Infoln("Found npm as dependency manager.The Filepath is ", value[0])

				case "gradle":
					gradle = new(gradlep.Gradle)
					gradle.FilePaths = append(gradle.FilePaths, value...)
					iscanners = append(iscanners, gradle)
					glog.Infoln("Found gradle as dependency manager.The Filepath is ", value[0])

				case "maven":
					maven = new(mavenp.Maven)
					maven.FilePaths = append(maven.FilePaths, value...)
					iscanners = append(iscanners, maven)
					glog.Infoln("Found maven as dependency manager.The Filepath is ", value[0])
				case "composer":
					composer = new(composerp.Composer)
					composer.FilePaths = append(composer.FilePaths, value...)
					iscanners = append(iscanners, composer)
					glog.Infoln("Found composer as dependency manager.The Filepath is ", value[0])
				case "mod":
					gomod = new(gomodp.Go)
					gomod.FilePaths = append(gomod.FilePaths, value...)
					iscanners = append(iscanners, gomod)
					glog.Infoln("Found go mod as dependency manager.The Filepath is ", value[0])
				case "nuget":
					nuget = new(nugetp.Nuget)
					nuget.FilePaths = append(nuget.FilePaths, value...)
					iscanners = append(iscanners, nuget)
					glog.Infoln("Found .net nuget as dependency manager.The Filepath is ", value[0])

				default:
					glog.Infoln("Unimplemented tool")
				}
			}
		}

		glog.Infoln("There are/is ", len(iscanners), "of scanners to scan")
		deps, err := impl.ScanAll(iscanners...)
		if err != nil {
			glog.Info(err)
			//glog.Errorln(err)
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
