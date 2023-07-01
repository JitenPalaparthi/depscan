package config

type DepManager struct {
	Lang      string   `json:"lang"`
	DepTool   string   `json:"depTool"`
	FileNames []string `json:"fileNames"`
	FileExt   string   `json:"fileExt"`
}
