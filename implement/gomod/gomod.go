package gomod

import (
	"bufio"
	"os"
	"path/filepath"
	"strings"

	"github.com/JitenPalaparthi/depscan/helper"
	"github.com/JitenPalaparthi/depscan/implement"
	scan "github.com/JitenPalaparthi/depscan/scanner"
	"github.com/golang/glog"
)

type Go struct {
	FilePaths []string
}

// This implementation is for composer.json file
func (g *Go) Scan() ([]scan.Dep, error) {
	gdeps := make([]scan.Dep, 0)
	duplicateDep := make(map[string][]string) // to ensure there are no duplicates
	gosumFile, gomodFile := "", ""

	if len(g.FilePaths) < 2 {
		return nil, implement.ErrInvalidNoOfFiles
	} else if filepath.Base(g.FilePaths[0]) == "go.sum" {
		gosumFile = g.FilePaths[0]
		gomodFile = g.FilePaths[1]
	} else if filepath.Base(g.FilePaths[0]) == "go.mod" {
		gosumFile = g.FilePaths[1]
		gomodFile = g.FilePaths[0]
	}

	// read gomodFile
	inFile1, err := os.Open(gomodFile)
	if err != nil {
		glog.Infoln(err)
		return nil, nil
	}
	scanner1 := bufio.NewScanner(inFile1)
	for scanner1.Scan() {

		line := scanner1.Text()
		if !strings.Contains(line, "// indirect") {
			continue
		}
		pkg, ver := "", ""
		isPkg, isVer := true, false
		for _, v := range line {
			if isPkg {
				if string(v) == " " {
					isPkg = false
					isVer = true
				} else {
					if string(v) != "\t" {
						pkg += string(v)
					}
				}
			} else if isVer {
				if string(v) == " " {
					isVer = false
					break
				} else {
					ver += string(v)
				}
			}
		}

		gdep := scan.Dep{}
		gdep.Version = ver
		gdep.Direct = false
		gdep.Name = pkg
		gdep.Type = "mod"
		gdep.Source = gosumFile
		gdeps = append(gdeps, gdep)

	}

	// read gosumFile
	inFile2, err := os.Open(gosumFile)

	if err != nil {
		glog.Infoln(err)
		return nil, nil
	}
	defer inFile2.Close()

	glog.Infoln("---------------XXXXX", inFile2)
	scanner2 := bufio.NewScanner(inFile2)
	for scanner2.Scan() {
		line := scanner2.Text()
		strs := strings.Split(line, " ")
		if len(strs) >= 2 {
			gdep := scan.Dep{}
			gdep.Direct = true
			gdep.Type = "mod"
			gdep.Name = strings.TrimSpace(strs[0])
			//version := strs[1]
			if strings.TrimSpace(strs[1])[0] == 'v' { // it should start with v
				version := strings.Split(strings.TrimSpace(strs[1]), "/")[0]
				gdep.Version = version
			}
			gdep.Source = gosumFile
			v1, ok1 := duplicateDep[gdep.Name]

			if !ok1 {
				duplicateDep[gdep.Name] = append(duplicateDep[gdep.Name], gdep.Version)
				gdeps = append(gdeps, gdep)
			} else {
				if !helper.IsElementExist(v1, gdep.Version) {
					duplicateDep[gdep.Name] = append(duplicateDep[gdep.Name], gdep.Version)
					gdeps = append(gdeps, gdep)
				}
			}
		}
	}
	return gdeps, nil
}

//github.com/BurntSushi/toml
