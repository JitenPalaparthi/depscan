package implement

import (
	"errors"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"

	"github.com/JitenPalaparthi/depscan/config"
	"github.com/JitenPalaparthi/depscan/scanner"
)

type Implement struct {
	Config              *config.Config
	Path                string
	Paths               []string
	Exts                []string
	Depth               uint8
	DirCount, FileCount int
}

// New is a function that is used to create/instantiate implement object
func New(config *config.Config, path string, depth uint8) (*Implement, error) {
	if config == nil {
		return nil, errors.New("nil config")
	}
	if path == "" {
		return nil, errors.New("empty path")
	}
	if depth <= 0 {
		return nil, errors.New("invalid depth")
	}
	implement := new(Implement)
	implement.Config = config
	implement.Path = path
	implement.Depth = depth
	return implement, nil

}

func (i *Implement) Feed() error {
	if i.Config == nil || i.Path == "" || i.Depth <= 0 {
		return errors.New("use New function to create Implement object")
	}
	maxDepth := strings.Count(i.Path, string(os.PathSeparator)) + int(i.Depth)
	filepath.WalkDir(i.Path, func(p string, d fs.DirEntry, err error) error {
		if d.IsDir() && strings.Count(p, string(os.PathSeparator)) > maxDepth {
			fmt.Println("skip", i.Path)
			return fs.SkipDir
		}
		if d.IsDir() {
			i.DirCount++
		} else {
			i.FileCount++
		}
		if !d.IsDir() {
			if isElementExist(i.Config.GetDepFiles(), d.Name()) {
				i.Paths = append(i.Paths, p)
			}
			i.Exts = append(i.Exts, filepath.Ext(d.Name()))
		}
		return nil
	})
	return nil
}

func (i *Implement) ScanAll(iScanners ...scanner.Scanner) ([]scanner.Dep, error) {
	var deps []scanner.Dep
	for _, s := range iScanners {
		d, err := s.Scan()
		if err != nil {
			return deps, err
		}
		deps = append(deps, d...)
	}
	return deps, nil
}

func isElementExist(s []string, str string) bool {
	for _, v := range s {
		if strings.TrimSpace(v) == strings.TrimSpace(str) {
			return true
		}
	}
	return false
}
