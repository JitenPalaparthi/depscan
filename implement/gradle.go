package implement

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	scan "github.com/JitenPalaparthi/depscan/scanner"
)

type Gradle struct {
	FilePath string
}

func (g *Gradle) Scan() ([]scan.Dep, error) {
	gdeps := make([]scan.Dep, 0)
	inFile, err := os.Open(g.FilePath)
	if err != nil {
		log.Println(err)
		return nil, nil
	}
	defer inFile.Close()

	scanner := bufio.NewScanner(inFile)
	found := false
	//lines := make([]string, 0)
	for scanner.Scan() {
		line := scanner.Text()
		//line = strings.TrimSpace(line)
		if strings.Contains(line, "}") {
			found = false
		}
		if strings.Contains(line, "dependencies {") {
			fmt.Println(line)
			found = true
		}
		if found && !strings.Contains(line, "dependencies {") {
			line = strings.Replace(line, `'`, `"`, -1) //fix for single quote
			fi := strings.Index(line, `"`)
			li := strings.LastIndex(line, `"`)
			//lines = append(lines, line[fi+1:li])
			gdep := scan.Dep{}
			gdep.Direct = true
			gdep.Type = "gradle"
			gdep.Name = line[fi+1 : li]
			//todo discuss and implement
			//gdep.Version = strs[1] // Unable to determine version. Some lines do not have version
			gdep.Source = g.FilePath
			gdeps = append(gdeps, gdep)
		}
	}
	return gdeps, nil
}
