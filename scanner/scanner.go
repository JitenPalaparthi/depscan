package scanner

type Scanner interface {
	Scan() ([]Dep, error)
}

type Dep struct {
	Name         string         `json:"name" yaml:"name"`
	Version      string         `json:"version" yaml:"version"`
	Direct       bool           `json:"direct" yaml:"direct"`
	Source       string         `json:"source" yaml:"source"`
	Type         string         `json:"type" yaml:"type"`
	Dependencies map[string]any `json:"dependencies,omitempty" yaml:"dependencies,omitempty"`
}
