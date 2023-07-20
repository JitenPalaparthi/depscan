package gradle

import (
	"fmt"
	"strings"

	"github.com/JitenPalaparthi/depscan/helper"
	scan "github.com/JitenPalaparthi/depscan/scanner"
	"github.com/golang/glog"
)

type Gradle struct {
	FilePaths []string
}

// This implementation is based on gradle.build
// The implmentation has been changed from gradle.build to dependencies.lock file
// func (g *Gradle) Scan() ([]scan.Dep, error) {
// 	gdeps := make([]scan.Dep, 0)
// 	inFile, err := os.Open(g.FilePaths[0])
// 	if err != nil {
// 		glog.Infoln(err)
// 		return nil, nil
// 	}
// 	defer inFile.Close()

// 	glog.Infoln("---------------XXXXX", inFile)
// 	scanner := bufio.NewScanner(inFile)
// 	found := false
// 	//lines := make([]string, 0)
// 	for scanner.Scan() {
// 		line := scanner.Text()
// 		//line = strings.TrimSpace(line)
// 		if strings.Contains(line, "}") {
// 			found = false
// 		}

// 		if strings.Contains(strings.ReplaceAll(line, " ", ""), "dependencies{") {
// 			glog.Infoln(line)
// 			found = true
// 			continue
// 		}
// 		if found {
// 			line = strings.Replace(line, `'`, `"`, -1) //fix for single quote
// 			fi := strings.Index(line, `"`)
// 			li := strings.LastIndex(line, `"`)
// 			//lines = append(lines, line[fi+1:li])
// 			gdep := scan.Dep{}
// 			gdep.Direct = true
// 			gdep.Type = "gradle"
// 			gdep.Name = line[fi+1 : li]
// 			//todo discuss and implement
// 			//gdep.Version = strs[1] // Unable to determine version. Some lines do not have version
// 			gdep.Source = g.FilePaths[0]
// 			gdeps = append(gdeps, gdep)
// 		}
// 	}
// 	return gdeps, nil
// }

func (g *Gradle) Scan() ([]scan.Dep, error) {
	gdeps := make([]scan.Dep, 0)

	mp, err := helper.FileToMap(g.FilePaths[0])
	if err != nil {
		glog.Infoln(err)
		return nil, err
	}

	depMap := make(map[string][]string)
	for _, v1 := range mp {
		for k2, v2 := range v1.(map[string]any) {
			str := fmt.Sprint(v2)
			str = strings.Replace(str, "map[", "", -1)
			str = strings.Replace(str, "]", "", -1)
			str = strings.Replace(str, "locked:", "", -1)
			glog.Infoln(k2, "----->>>", str)
			v3, ok3 := depMap[k2]
			isDuplicate := true
			if ok3 {
				if !helper.IsElementExist(v3, str) {
					depMap[k2] = append(depMap[k2], str)

				}
			} else {
				depMap[k2] = append(depMap[k2], str)
				isDuplicate = false
			}
			if !isDuplicate {
				isDuplicate = false
				gdep := scan.Dep{}
				gdep.Direct = true
				gdep.Type = "gradle"
				gdep.Name = k2
				gdep.Version = str
				gdep.Source = g.FilePaths[0]
				gdeps = append(gdeps, gdep)
			}
		}
	}

	return gdeps, nil
}
