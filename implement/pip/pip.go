package pip

import (
	"bufio"
	"os"
	"strings"

	"log"

	scan "github.com/JitenPalaparthi/depscan/scanner"
)

type Pip struct {
	FilePaths []string
}

func (p *Pip) Scan() ([]scan.Dep, error) {
	gdeps := make([]scan.Dep, 0)
	inFile, err := os.Open(p.FilePaths[0])
	if err != nil {
		log.Println(err)
		return nil, nil
	}
	defer inFile.Close()

	scanner := bufio.NewScanner(inFile)
	for scanner.Scan() {
		line := scanner.Text()
		line = strings.TrimSpace(line)
		if line[0] == '#' {
			continue
		}
		strs := strings.Split(line, "==")
		if len(strs) == 2 {
			gdep := scan.Dep{}
			gdep.Direct = true
			gdep.Type = "pip"
			gdep.Name = strs[0]
			gdep.Version = strs[1]
			gdep.Source = p.FilePaths[0]
			gdeps = append(gdeps, gdep)
		} else if !strings.Contains(line, "==") { // There are entries without == as well.i.e no version information
			gdep := scan.Dep{}
			gdep.Direct = true
			gdep.Type = "pip"
			gdep.Name = strings.TrimSpace(line)
			gdep.Version = ""
			gdep.Source = p.FilePaths[0]
			gdeps = append(gdeps, gdep)
		}
		// the actual logic goes here
	}

	return gdeps, nil
}
