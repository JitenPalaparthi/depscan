package maven

import (
	"bufio"
	"encoding/xml"
	"os"
	"strings"

	scan "github.com/JitenPalaparthi/depscan/scanner"
	"github.com/golang/glog"
)

type Maven struct {
	FilePaths []string
}

type Result struct {
	XMLName      xml.Name     `xml:"dependencies"`
	Dependencies []Dependency `xml:"dependency"`
}
type Dependency struct {
	XMLName    xml.Name `xml:"dependency"`
	GroupID    string   `xml:"groupId"`
	ArtifactId string   `xml:"artifactId"`
	Version    string   `xml:"version"`
}

func (m *Maven) Scan() ([]scan.Dep, error) {
	gdeps := make([]scan.Dep, 0)
	result := new(Result)

	inFile, err := os.Open(m.FilePaths[0])
	if err != nil {
		glog.Infoln(err)
		return nil, err
	}
	defer inFile.Close()
	// byteValue, err := ioutil.ReadAll(xmlFile)
	// if err != nil {
	// 	return nil, err
	// }
	scanner := bufio.NewScanner(inFile)
	found := false
	str := `<?xml version="1.0" encoding="UTF-8"?>`
	//lines := make([]string, 0)
	for scanner.Scan() {
		line := scanner.Text()
		//line = strings.TrimSpace(line)
		if strings.Contains(line, "</dependencies>") {
			found = false
			str = str + "</dependencies>"
		}
		if strings.Contains(line, "<dependencies>") {
			glog.Infoln(line)
			found = true
		}
		if found {
			str = str + line
		}
	}

	err = xml.Unmarshal([]byte(str), result)
	if err != nil {
		glog.Infoln(err)
		return nil, err
	}
	for _, dependency := range result.Dependencies {
		gdep := scan.Dep{}
		gdep.Direct = true
		gdep.Type = "maven"
		gdep.Name = dependency.GroupID + ":" + dependency.ArtifactId
		gdep.Version = dependency.Version
		gdep.Source = m.FilePaths[0]
		gdeps = append(gdeps, gdep)
	}
	return gdeps, nil
}
