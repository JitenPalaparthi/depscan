package gomod

import (
	"bufio"
	"os"
	"strings"

	"github.com/JitenPalaparthi/depscan/helper"
	scan "github.com/JitenPalaparthi/depscan/scanner"
	"github.com/golang/glog"
)

type Go struct {
	FilePaths []string
}

// This implementation is for composer.json file
func (g *Go) Scan() ([]scan.Dep, error) {
	gdeps := make([]scan.Dep, 0)
	duplicateDep := make(map[string][]string)
	inFile, err := os.Open(g.FilePaths[0])
	if err != nil {
		glog.Infoln(err)
		return nil, nil
	}
	defer inFile.Close()

	glog.Infoln("---------------XXXXX", inFile)
	scanner := bufio.NewScanner(inFile)
	for scanner.Scan() {
		line := scanner.Text()
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
			gdep.Source = g.FilePaths[0]
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
