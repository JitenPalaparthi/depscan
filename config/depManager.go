package config

type DepManager struct {
	Lang     string `json:"lang"`
	DepTool  string `json:"depTool"`
	FileName string `json:"fileName"`
	FileExt  string `json:"fileExt"`
}
