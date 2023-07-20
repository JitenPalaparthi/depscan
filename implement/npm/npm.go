package npm

import (
	"fmt"
	"path/filepath"
	"reflect"
	"strings"

	"github.com/JitenPalaparthi/depscan/helper"
	"github.com/JitenPalaparthi/depscan/implement"
	scan "github.com/JitenPalaparthi/depscan/scanner"
	"github.com/golang/glog"
)

type Npm struct {
	FilePaths []string
}

type NpmVersion1DataFormat struct {
	Version  string `json:"version"`
	Resolved string `json:"resolved"`
	Dev      bool   `json:"dev"`
}

func (n *Npm) Scan() ([]scan.Dep, error) {
	gdeps := make([]scan.Dep, 0)
	depMap := make(map[string]any)

	// a path can have maximum two files w.r.t npm.
	// 1- package-lock.json
	// 2- package.json

	// if there are both the files? then process first package.json
	// keep values in a map and change direct == true if there is an entry
	packageFile := ""
	packagelockFile := ""
	for _, filePath := range n.FilePaths {
		fileName := filepath.Base(filePath)
		if fileName == "package.json" {
			packageFile = filePath
		} else {
			packagelockFile = filePath
		}
	}

	if len(n.FilePaths) == 2 {
		if mp, err := helper.FileToMap(packageFile); err != nil {
			return nil, err
		} else {
			depMap = mp["dependencies"].(map[string]any)
		}

		if mp, err := helper.FileToMap(packagelockFile); err != nil {
			return nil, err
		} else {
			if mp["lockfileVersion"].(float64) == 1 {
				glog.Infoln("This is lockfileVersion:1")
				fmt.Println(reflect.TypeOf(mp))
				data := mp["dependencies"]
				switch data := data.(type) {
				case map[string]any:
					for k1, v1 := range data {
						gdep := scan.Dep{}

						// to fetch dependencies
						dependencies, ok := v1.(map[string]any)["dependencies"]
						if ok {
							glog.Infoln("<<<<-------Sub-Dependencies--------->>>>")
							glog.Infoln(dependencies)
							gdep.Dependencies = dependencies.(map[string]any)
						}

						//fmt.Println("Key-->", k1) // "Type of Value:", reflect.TypeOf(v))
						//if strings.Contains(k1, "node_modules/") {
						//k1 = strings.Replace(k1, "node_modules/", "", 1)
						li := strings.LastIndex(k1, "node_modules/")
						if li != -1 {
							li = li + len("node_modules/")
							k1 = k1[li:]
						}
						//}

						gdep.Name = k1
						//k1s := strings.Split(k1, "/") //
						yes := false
						for key := range depMap {
							if strings.HasSuffix(k1, key) {
								yes = true
								break
							}
						}
						if yes {
							gdep.Direct = true
						} else {
							gdep.Direct = false
						}

						gdep.Type = "npm"
						gdep.Source = packagelockFile
						isDev := false
						for k2, v2 := range v1.(map[string]any) {
							if k2 == "version" {
								gdep.Version = v2.(string)
							}
							if k2 == "dev" {
								isDev = true
							}
						}
						if !isDev {
							gdeps = append(gdeps, gdep)
						}
					}

				}
				//fmt.Println(gdeps)
				//todo logic for version 1
			} else if mp["lockfileVersion"].(float64) == 2 || mp["lockfileVersion"].(float64) == 3 {
				glog.Infoln("This is lockfileVersion:", mp["lockfileVersion"].(float64))
				data := mp["packages"]
				//fmt.Println(reflect.TypeOf(data))
				switch data := data.(type) {
				case map[string]any:
					for k1, v1 := range data {
						if k1 == "devDependencies" || k1 == "" {
							continue
						}
						gdep := scan.Dep{}
						dependencies, ok := v1.(map[string]any)["dependencies"]
						if ok {
							glog.Infoln("<<<<-------Sub-Dependencies--------->>>>")
							glog.Infoln(dependencies)
							gdep.Dependencies = dependencies.(map[string]any)

						}
						//fmt.Println("Key-->", k1) // "Type of Value:", reflect.TypeOf(v))
						//k1 = strings.Replace(k1, "node_modules/", "", 1)
						li := strings.LastIndex(k1, "node_modules/")
						if li != -1 {
							li = li + len("node_modules/")
							k1 = k1[li:]
						}
						gdep.Name = k1
						yes := false
						for key := range depMap {
							if strings.HasSuffix(k1, key) {
								yes = true
								break
							}
						}
						if yes {
							gdep.Direct = true
						} else {
							gdep.Direct = false
						}
						gdep.Type = "npm"
						gdep.Source = packagelockFile
						//isDev := false

						for k2, v2 := range v1.(map[string]any) {
							if k2 == "version" {
								gdep.Version = v2.(string)
							}
						}
						//if !isDev {
						gdeps = append(gdeps, gdep)
						//}

					}
				}
			} else {
				return nil, implement.ErrUnsupportedNPMVersion
			}
		}
	} else if len(n.FilePaths) == 1 && packagelockFile != "" {
		if mp, err := helper.FileToMap(packagelockFile); err != nil {
			return nil, err
		} else {
			if mp["lockfileVersion"].(float64) == 1 {
				glog.Infoln("This is lockfileVersion:1")
				fmt.Println(reflect.TypeOf(mp))
				data := mp["dependencies"]
				switch data := data.(type) {
				case map[string]any:
					for k1, v1 := range data {
						gdep := scan.Dep{}
						//fmt.Println("Key-->", k1) // "Type of Value:", reflect.TypeOf(v))
						dependencies, ok := v1.(map[string]any)["dependencies"]
						if ok {
							glog.Infoln("<<<<-------Sub-Dependencies--------->>>>")
							glog.Infoln(dependencies)
							gdep.Dependencies = dependencies.(map[string]any)

						}

						//fmt.Println("Key-->", k1) // "Type of Value:", reflect.TypeOf(v))
						//if strings.Contains(k1, "node_modules/") {
						//k1 = strings.Replace(k1, "node_modules/", "", 1)
						//}
						li := strings.LastIndex(k1, "node_modules/")
						if li != -1 {
							li = li + len("node_modules/")
							k1 = k1[li:]
						}
						gdep.Name = k1
						//k1s := strings.Split(k1, "/") //
						yes := false
						for key := range depMap {
							if strings.HasSuffix(k1, key) {
								yes = true
								break
							}
						}
						if yes {
							gdep.Direct = true
						} else {
							gdep.Direct = false
						}
						gdep.Type = "npm"
						gdep.Source = packagelockFile
						isDev := false
						for k2, v2 := range v1.(map[string]any) {
							if k2 == "version" {
								gdep.Version = v2.(string)
							}
							if k2 == "dev" {
								isDev = true
							}
						}
						if !isDev {
							gdeps = append(gdeps, gdep)
						}
					}

				}
				//fmt.Println(gdeps)
				//todo logic for version 1
			} else if mp["lockfileVersion"].(float64) == 2 || mp["lockfileVersion"].(float64) == 3 {
				glog.Infoln("This is lockfileVersion:", mp["lockfileVersion"].(float64))
				data := mp["packages"]
				//fmt.Println(reflect.TypeOf(data))
				switch data := data.(type) {
				case map[string]any:
					for k1, v1 := range data {
						if k1 == "devDependencies" || k1 == "" {
							continue
						}
						gdep := scan.Dep{}
						//fmt.Println("Key-->", k1) // "Type of Value:", reflect.TypeOf(v))
						dependencies, ok := v1.(map[string]any)["dependencies"]
						if ok {
							glog.Infoln("<<<<-------Sub-Dependencies--------->>>>")
							glog.Infoln(dependencies)
							gdep.Dependencies = dependencies.(map[string]any)
						}

						//fmt.Println("Key-->", k1) // "Type of Value:", reflect.TypeOf(v))
						//if strings.Contains(k1, "node_modules/") {
						//k1 = strings.Replace(k1, "node_modules/", "", 1)
						//}
						li := strings.LastIndex(k1, "node_modules/")
						if li != -1 {
							li = li + len("node_modules/")
							k1 = k1[li:]
						}
						gdep.Name = k1
						//k1s := strings.Split(k1, "/") //
						yes := false
						for key := range depMap {
							if strings.HasSuffix(k1, key) {
								yes = true
								break
							}
						}
						if yes {
							gdep.Direct = true
						} else {
							gdep.Direct = false
						}
						gdep.Type = "npm"
						gdep.Source = packagelockFile
						//isDev := false

						for k2, v2 := range v1.(map[string]any) {
							if k2 == "version" {
								gdep.Version = v2.(string)
							}
						}
						//if !isDev {
						gdeps = append(gdeps, gdep)
						//}

					}
				}
			} else {
				return nil, implement.ErrUnsupportedNPMVersion
			}
		}
	} else if len(n.FilePaths) == 1 && packageFile != "" {
		if mp, err := helper.FileToMap(packageFile); err != nil {
			return nil, err
		} else {
			depMap = mp["dependencies"].(map[string]any)
		}

		for k1, v1 := range depMap {
			gdep := scan.Dep{}
			gdep.Name = k1
			gdep.Direct = true
			gdep.Type = "npm"
			gdep.Source = packageFile
			gdep.Version = fmt.Sprint(v1)
			gdeps = append(gdeps, gdep)

		}
	} else {
		return nil, implement.ErrNoPackageFileFound
	}

	return gdeps, nil
}
