package implement

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"reflect"

	scan "github.com/JitenPalaparthi/depscan/scanner"
	"github.com/golang/glog"
)

type Npm struct {
	FilePaths []string
}

// type NpmLockfileVerson2 struct {
// 	Packages map[string]struct {
// 		Version  string `json:"version"`
// 		Resolved string `json:"resolved"`
// 		Dev      bool   `json:"dev"`
// 	}
// }

type NpmVersion1DataFormat struct {
	Version  string `json:"version"`
	Resolved string `json:"resolved"`
	Dev      bool   `json:"dev"`
}

func (n *Npm) Scan() ([]scan.Dep, error) {
	gdeps := make([]scan.Dep, 0)

	bytes, err := ioutil.ReadFile(n.FilePaths[0])
	if err != nil {
		return nil, err
	}
	mp := make(map[string]any)
	err = json.Unmarshal(bytes, &mp)
	if err != nil {
		fmt.Println(err)
	}

	if mp["lockfileVersion"].(float64) == 1 {
		glog.Infoln("This is lockfileVersion:1")
		fmt.Println(reflect.TypeOf(mp))
		data := mp["dependencies"]
		switch data := data.(type) {
		case map[string]any:
			for k1, v1 := range data {
				gdep := scan.Dep{}
				//fmt.Println("Key-->", k1) // "Type of Value:", reflect.TypeOf(v))
				gdep.Name = k1
				gdep.Direct = true
				gdep.Type = "npm"
				gdep.Source = n.FilePaths[0]
				isDev := false
				for k2, v2 := range v1.(map[string]any) {
					if k2 == "version" {
						gdep.Version = v2.(string)
					}
					// if k2 == "revision" {

					// }
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
	} else if mp["lockfileVersion"].(float64) == 2 {
		glog.Infoln("This is lockfileVersion:2")
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
				gdep.Name = k1
				gdep.Direct = true
				gdep.Type = "npm"
				gdep.Source = n.FilePaths[0]
				isDev := false

				for k2, v2 := range v1.(map[string]any) {
					if k2 == "version" {
						gdep.Version = v2.(string)
					}
					// if k2 == "revision" {

					// }
					if k2 == "dev" {
						isDev = true
					}
				}
				if !isDev {
					gdeps = append(gdeps, gdep)
				}

			}
		}
	} else {
		return nil, ErrUnsupportedNPMVersion
	}

	//data := mp["pakages"]
	//fmt.Println(reflect.TypeOf(mp), mp)

	// for key, val := range data {
	// 	fmt.Println(key, "---------->")
	// 	fmt.Println(val)
	// 	fmt.Println("-----------------------------------------------------")
	// }
	//mp["packages"]
	return gdeps, nil
}
