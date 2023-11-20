package helper

import (
	"bytes"
	"encoding/json"
	"os"
	"path/filepath"
	"strings"
)

func IsElementExist(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}
	return false
}

func IsFileAndExtExists(s []string, str string) bool {
	for _, v := range s {
		if strings.Contains(v, "*.") {
			if filepath.Ext(v) == filepath.Ext(str) {
				//fmt.Println("XXXXXXX---", v, str)
				return true
			}
		}
		if v == str {
			return true
		}
	}
	return false
}

func FileToMap(filename string) (map[string]any, error) {
	bytes, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	mp := make(map[string]any)
	err = json.Unmarshal(bytes, &mp)
	if err != nil {
		return nil, err
	}
	return mp, nil
}

func JsonPrettyPrint(in string) string {
	var out bytes.Buffer
	err := json.Indent(&out, []byte(in), "", "\t")
	if err != nil {
		return in
	}
	return out.String()
}
