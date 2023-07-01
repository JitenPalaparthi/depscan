package implement

import (
	scan "github.com/JitenPalaparthi/depscan/scanner"
)

type Npm struct {
	FilePaths []string
}

func (m *Npm) Scan() ([]scan.Dep, error) {
	gdeps := make([]scan.Dep, 0)

	// bytes, err := ioutil.ReadFile(m.FilePath)
	// if err != nil {
	// 	return nil, err
	// }

	return gdeps, nil
}
