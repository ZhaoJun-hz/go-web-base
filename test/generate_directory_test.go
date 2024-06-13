package test

import (
	"encoding/json"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

var stRootDir string
var stSeparator string
var jsonData map[string]any

const jsonFileName = "directory.json"

func loadJson() {
	stSeparator = string(filepath.Separator)
	stWorkDir, _ := os.Getwd()
	stRootDir = stWorkDir[:strings.LastIndex(stWorkDir, stSeparator)]

	bytes, _ := os.ReadFile(stWorkDir + stSeparator + jsonFileName)
	err := json.Unmarshal(bytes, &jsonData)
	if err != nil {
		panic("failed to load json: " + err.Error())
	}
}

func parseMap(mapData map[string]any, parentDir string) {
	for key, value := range mapData {
		switch value.(type) {
		case string:
			{
				path := value.(string)
				if path == "" {
					continue
				}
				if parentDir != "" {
					path = parentDir + stSeparator + path
					if key == "text" {
						parentDir = path
					}
				} else {
					parentDir = path
				}
				createDir(path)
			}
		case []any:
			{
				parseArray(value.([]any), parentDir)
			}
		}
	}
}

func parseArray(jsonData []any, parentDir string) {
	for _, value := range jsonData {
		mapValue := value.(map[string]any)
		parseMap(mapValue, parentDir)
	}
}

func createDir(path string) {
	if path == "" {
		return
	}
	err := os.MkdirAll(stRootDir+stSeparator+path, fs.ModePerm)
	if err != nil {
		panic("createDir Error: " + err.Error())
	}
}

func TestGenerateDirectory(t *testing.T) {
	loadJson()
	parseMap(jsonData, "")
}
