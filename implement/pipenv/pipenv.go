package pipenv

import (
	"encoding/json"
	"os"

	"log"

	scan "github.com/JitenPalaparthi/depscan/scanner"
)

type Pipenv struct {
	FilePaths []string
}

func (p *Pipenv) Scan() ([]scan.Dep, error) {
	gdeps := make([]scan.Dep, 0)
	inFile, err := os.ReadFile(p.FilePaths[0])
	if err != nil {
		log.Println(err)
		return nil, nil
	}

	deps := make(map[string]any)

	err = json.Unmarshal(inFile, &deps)
	if err != nil {
		return nil, err
	}

	val1, ok1 := deps["default"]
	if ok1 { // if it finds default key then there are package details
		switch val1.(type) {
		case map[string]any:
			for k2, v2 := range val1.(map[string]any) {
				switch v2.(type) {
				case map[string]any:
					//fmt.Println(v2, "Type:", reflect.TypeOf(v2))
					gdep := scan.Dep{}
					gdep.Direct = true
					gdep.Type = "pipenv"
					gdep.Name = k2 //name.(string)
					version := v2.(map[string]any)["version"]
					if version != nil {
						gdep.Version = version.(string)
					}
					gdep.Source = p.FilePaths[0]
					gdeps = append(gdeps, gdep)
				}

			}
		}

	}

	return gdeps, nil
}
