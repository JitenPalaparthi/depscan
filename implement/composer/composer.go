package composer

import (
	"bufio"
	"os"
	"strings"

	scan "github.com/JitenPalaparthi/depscan/scanner"
	"github.com/golang/glog"
)

type Composer struct {
	FilePaths []string
}

func (c *Composer) Scan() ([]scan.Dep, error) {
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
