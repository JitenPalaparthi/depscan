package implement

import (
	scan "github.com/JitenPalaparthi/depscan/scanner"
)

type Npm struct {
	FilePath string
}

func (m *Npm) Scan() ([]scan.Dep, error) {
	gdeps := make([]scan.Dep, 0)
	return gdeps, nil
}
