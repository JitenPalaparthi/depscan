package implement

import (
	"encoding/json"
	"errors"
	"io/fs"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/JitenPalaparthi/depscan/config"
	"github.com/JitenPalaparthi/depscan/helper"
	"github.com/JitenPalaparthi/depscan/scanner"
	"github.com/golang/glog"
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

var (
	ErrNilConfig               = errors.New("nil config")
	ErrEmptyPath               = errors.New("empty path")
	ErrInvalidOutfile          = errors.New("invalid outfile")
	ErrInvalidOutfileFormat    = errors.New("invalid outfile format.It must be json | yaml | yml")
	ErrNewImplement            = errors.New("use New function to create Implement object")
	ErrNoDataToGenerateOutfile = errors.New("no data to generate output file")
	ErrPathDoesNotExist        = errors.New("path does not exist")
	ErrUnsupportedNPMVersion   = errors.New("unsupported npm version.Currently it supports only lockfileVersion-1,2 and 3 only.Old formats are not supported")
)

// New is a function that is used to create/instantiate implement object
func New(config *config.Config, path string, outfile string, depth uint8) (*Implement, error) {
	if config == nil {
		return nil, ErrNilConfig
	}
	if path == "" {
		return nil, ErrEmptyPath
	}
	if _, err := os.Stat(path); os.IsNotExist(err) {
		// path/to/whatever does not exist
		return nil, ErrPathDoesNotExist
	}
	// if depth < 0 {
	// 	return nil, errors.New("invalid depth")
	// }
	if outfile == "" {
		return nil, ErrInvalidOutfile
	}
	if filepath.Ext(outfile) == "json" || filepath.Ext(outfile) == "yaml" || filepath.Ext(outfile) == "yml" {
		return nil, ErrInvalidOutfileFormat
	}

	implement := new(Implement)
	implement.Config = config
	implement.Path = path
	implement.Depth = depth
	implement.Outfile = outfile
	return implement, nil
}

func (i *Implement) Feed() error {
	if i.Config == nil || i.Path == "" {
		return ErrNewImplement
	}
	maxDepth := strings.Count(i.Path, string(os.PathSeparator)) + int(i.Depth)
	glog.Infoln("evaluating depth param.Current depth value:", i.Depth)
	filepath.WalkDir(i.Path, func(p string, d fs.DirEntry, err error) error {
		if d.IsDir() && strings.Count(p, string(os.PathSeparator)) > maxDepth {
			glog.Info("skip path(s):", p)
			return fs.SkipDir
		}
		if d.IsDir() {
			i.DirCount++
		} else {
			i.FileCount++
		}
		if d.IsDir() && helper.IsElementExist(i.Config.IgnoreDirs, d.Name()) {
			glog.Infoln("This is an ignore directory so skipping it. path:", p)
			return fs.SkipDir
		}
		if !d.IsDir() {
			if helper.IsElementExist(i.Config.GetDepFiles(), d.Name()) {
				i.Paths = append(i.Paths, p)
				Dep := i.Config.GetDepManagerByFileName(d.Name()) // added to add language even if no js or other programming files. Just based on the Dep file. For example requirements.txt
				i.Languages = append(i.Languages, Dep.Lang)       //
			}

			// The below code does these things.
			// 1- Get programming file extensions. Example .java,.py .js etc
			// Check whether that extension exists in the config file.
			// if the ext exists add it to the implement object exts.
			// based on ext add languages to the implemen object.
			// This logic makes sure that it does not append multiple entries of same extension and also multiple entries of langauges.
			// so duplicates of implement.Exts and implement.Langauges
			if helper.IsElementExist(i.Config.GetExtensions(), filepath.Ext(d.Name())) {
				if !helper.IsElementExist(i.Exts, filepath.Ext(d.Name())) {
					i.Exts = append(i.Exts, filepath.Ext(d.Name()))
					Dep := i.Config.GetDepManagerByExt(filepath.Ext(d.Name()))
					i.Languages = append(i.Languages, Dep.Lang)
				}
			}
		}
		return nil
	})
	return nil
}

// ScanAll scans all language implementations based on Scanner interface
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

// Write dependency data to the file.
func (i *Implement) Write(deps []scanner.Dep) error {
	if len(deps) == 0 {
		return ErrNoDataToGenerateOutfile
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
