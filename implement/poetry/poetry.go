package poetry

import (
	"os"

	"log"

	"github.com/BurntSushi/toml"
	scan "github.com/JitenPalaparthi/depscan/scanner"
)

type Poetry struct {
	FilePaths []string
}

func (p *Poetry) Scan() ([]scan.Dep, error) {
	gdeps := make([]scan.Dep, 0)
	inFile, err := os.ReadFile(p.FilePaths[0])
	if err != nil {
		log.Println(err)
		return nil, nil
	}

	deps := make(map[string]any)

	err = toml.Unmarshal(inFile, &deps)
	if err != nil {
		return nil, err
	}
	v1, ok1 := deps["package"]

	if ok1 {
		for _, v2 := range v1.([]map[string]any) {
			if len(v2) > 0 {
				name, ok2 := v2["name"]
				version, ok3 := v2["version"]
				if ok2 && ok3 {
					gdep := scan.Dep{}
					gdep.Direct = true
					gdep.Type = "poetry"
					gdep.Name = name.(string)
					gdep.Version = version.(string)
					gdep.Source = p.FilePaths[0]
					gdeps = append(gdeps, gdep)
				}
			}
		}
	}
	return gdeps, nil
}
