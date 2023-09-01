package composer

import (
	"bufio"
	"os"
	"strings"

	"github.com/JitenPalaparthi/depscan/helper"
	scan "github.com/JitenPalaparthi/depscan/scanner"
	"github.com/golang/glog"
)

type Composer struct {
	FilePaths []string
}

// This implementation is for composer.json file
func (c *Composer) ScanFor() ([]scan.Dep, error) {
	gdeps := make([]scan.Dep, 0)
	inFile, err := os.Open(c.FilePaths[0])
	if err != nil {
		glog.Infoln(err)
		return nil, nil
	}
	defer inFile.Close()

	glog.Infoln("---------------XXXXX", inFile)
	scanner := bufio.NewScanner(inFile)
	found := false
	//lines := make([]string, 0)
	for scanner.Scan() {
		line := scanner.Text()
		//line = strings.TrimSpace(line)
		if strings.Contains(line, "}") {
			found = false
		}

		if strings.Contains(line, `"require"`) {
			glog.Infoln(line)
			found = true
			continue
		}
		if found {
			line = strings.Replace(line, ",", "", -1) //remove training comma
			line = strings.Replace(line, `"`, "", -1) //remove quotes in a string
			line = strings.TrimSpace(line)            // trim leading and traling spaces
			strs := strings.Split(line, ":")
			gdep := scan.Dep{}
			gdep.Direct = true
			gdep.Type = "composer"
			gdep.Name = strings.TrimSpace(strs[0])
			gdep.Version = strings.TrimSpace(strs[1])
			gdep.Source = c.FilePaths[0]
			gdeps = append(gdeps, gdep)
		}
	}
	return gdeps, nil
}

// This implementation is for composer.lock
func (c *Composer) Scan() ([]scan.Dep, error) {
	gdeps := make([]scan.Dep, 0)
	duplicateDep := make(map[string][]string)
	mp, err := helper.FileToMap(c.FilePaths[0])
	if err != nil {
		return nil, err
	}
	v1, ok1 := mp["packages"]
	if ok1 {
		for _, v2 := range v1.([]any) {
			gdep := scan.Dep{}
			gdep.Direct = true
			gdep.Type = "composer"
			gdep.Name = v2.(map[string]any)["name"].(string)
			gdep.Version = v2.(map[string]any)["version"].(string)
			gdep.Source = c.FilePaths[0]

			v4, ok4 := v2.(map[string]any)["require"]
			if ok4 {
				gdep.Dependencies = v4.(map[string]any)
			}
			v3, ok3 := duplicateDep[gdep.Name]

			if !ok3 {
				duplicateDep[gdep.Name] = append(duplicateDep[gdep.Name], gdep.Version)
				gdeps = append(gdeps, gdep)
			} else {
				if !helper.IsElementExist(v3, gdep.Version) {
					duplicateDep[gdep.Name] = append(duplicateDep[gdep.Name], gdep.Version)
					gdeps = append(gdeps, gdep)
				}
			}

		}
	}

	return gdeps, nil
}
