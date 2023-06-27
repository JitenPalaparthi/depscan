package implement

import (
	scnr "github.com/JitenPalaparthi/depscan/scanner"
)

type Output struct {
	Path      string     `json:"path" yaml:"path"`
	Count     uint16     `json:"count" yaml:"count"`
	Languages []string   `json:"languages" yaml:"languages"`
	Items     []scnr.Dep `json:"items" yaml:"items"`
}
