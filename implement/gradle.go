package implement

import (
	scan "github.com/JitenPalaparthi/depscan/scanner"
)

type Gradle struct {
	FilePath string
}

func (g *Gradle) Scan() ([]scan.Dep, error) {
	gdeps := make([]scan.Dep, 0)
	// inFile, err := os.Open(filepath)
	// if err != nil {
	// 	log.Println(err)
	// 	return nil, nil
	// }
	// defer inFile.Close()

	// scanner := bufio.NewScanner(inFile)
	// for scanner.Scan() {

	// 	line := scanner.Text()
	// 	line = strings.TrimSpace(line)
	// 	if line[0] == '#' {
	// 		continue
	// 	}
	// 	strs := strings.Split(line, "==")
	// 	if len(strs) == 2 {
	// 		gdep := scan.Dep{}
	// 		gdep.Direct = true
	// 		gdep.Type = "gradle"
	// 		gdep.Name = strs[0]
	// 		gdep.Version = strs[1]
	// 		gdep.Source = filepath
	// 		gdeps = append(gdeps, gdep)
	// 	}
	// 	// the actual logic goes here
	// }

	return gdeps, nil
}
