package implement

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/fs"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/JitenPalaparthi/depscan/config"
	"github.com/JitenPalaparthi/depscan/scanner"
	"gopkg.in/yaml.v2"
)

type Implement struct {
	Config              *config.Config
	Path                string
	Outfile             string
	Paths               []string
	Exts                []string
	Languages           []string
	Depth               uint8
	DirCount, FileCount int
}

type Formatter interface {
	Marshal(v any) ([]byte, error)
}

// New is a function that is used to create/instantiate implement object
func New(config *config.Config, path string, outfile string, depth uint8) (*Implement, error) {
	if config == nil {
		return nil, errors.New("nil config")
	}
	if path == "" {
		return nil, errors.New("empty path")
	}
	if depth <= 0 {
		return nil, errors.New("invalid depth")
	}
	if outfile == "" {
		return nil, errors.New("invalid outfile")
	}
	if filepath.Ext(outfile) == "json" || filepath.Ext(outfile) == "yaml" || filepath.Ext(outfile) == "yml" {
		return nil, errors.New("invalid outfile format.It must be json | yaml | yml")
	}

	implement := new(Implement)
	implement.Config = config
	implement.Path = path
	implement.Depth = depth
	implement.Outfile = outfile
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
			if isElementExist(i.Config.GetExtensions(), filepath.Ext(d.Name())) {
				if !isElementExist(i.Exts, filepath.Ext(d.Name())) {
					i.Exts = append(i.Exts, filepath.Ext(d.Name()))
					// fmt.Println("------->", d.Name())
					// fmt.Println("----------->", filepath.Ext(d.Name()))
					Dep := i.Config.GetDepManagerByExt(filepath.Ext(d.Name()))
					i.Languages = append(i.Languages, Dep.Lang)
				}
			}
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

func (i *Implement) Write(deps []scanner.Dep) error {
	if len(deps) == 0 {
		return errors.New("no files to generate output file")
	}
	output := new(Output)
	output.Count = uint16(len(deps))
	output.Items = deps
	output.Languages = i.Languages
	output.Path = i.Path
	ext := filepath.Ext(i.Outfile)

	switch ext {
	case ".json":
		data, err := json.Marshal(output)
		if err != nil {
			return err
		}
		return ioutil.WriteFile(i.Outfile, data, 0655)
	case ".yaml", ".yml":
		data, err := yaml.Marshal(output)
		if err != nil {
			return err
		}
		return ioutil.WriteFile(i.Outfile, data, 0655)
	}
	return nil
}

func isElementExist(s []string, str string) bool {
	for _, v := range s {
		if strings.TrimSpace(v) == strings.TrimSpace(str) {
			return true
		}
	}
	return false
}
