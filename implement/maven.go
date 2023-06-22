package implement

import (
	scan "github.com/JitenPalaparthi/depscan/scanner"
)

type Maven struct {
	FilePath string
}

func (m *Maven) Scan() ([]scan.Dep, error) {
	gdeps := make([]scan.Dep, 0)
	return gdeps, nil
}
