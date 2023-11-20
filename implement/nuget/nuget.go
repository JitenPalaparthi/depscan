package nuget

import (
	"bufio"
	"os"
	"strings"

	"log"

	scan "github.com/JitenPalaparthi/depscan/scanner"
)

type Nuget struct {
	FilePaths []string
}

func (p *Nuget) Scan() ([]scan.Dep, error) {
	gdeps := make([]scan.Dep, 0)
	inFile, err := os.Open(p.FilePaths[0])
	duplicateDeps := make(map[string]bool)
	if err != nil {
		log.Println(err)
		return nil, nil
	}
	defer inFile.Close()

	scanner := bufio.NewScanner(inFile)

	for scanner.Scan() {
		line := scanner.Text()
		line = strings.TrimSpace(line)

		if strings.HasPrefix(line, "<PackageReference") && strings.HasSuffix(line, "/>") && strings.Contains(line, "Include=") && strings.Contains(line, "Version=") {
			lines := strings.Split(line, " ")
			versionName, version := "", ""
			for _, v := range lines {
				if strings.Contains(v, "Include=") {
					fi, li := strings.Index(v, `"`), strings.LastIndex(v, `"`)
					versionName = v[fi+1 : li]
				} else if strings.Contains(v, "Version=") {
					fi, li := strings.Index(v, `"`), strings.LastIndex(v, `"`)
					version = v[fi+1 : li]
				} else {
					continue
				}
			}
			if _, ok := duplicateDeps[versionName+version]; !ok { // to avoid duplicates
				duplicateDeps[versionName+version] = true
				gdep := scan.Dep{}
				gdep.Direct = true
				gdep.Type = ".net"
				gdep.Name = versionName
				gdep.Version = version
				gdep.Source = p.FilePaths[0]
				gdeps = append(gdeps, gdep)
			}
		} else {
			continue
		}

	}
	return gdeps, nil
}
