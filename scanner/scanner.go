package scanner

type Scanner interface {
	Scan() ([]Dep, error)
}

type Dep struct {
	Name    string
	Version string
	Direct  bool
	Source  string
	Type    string
}
