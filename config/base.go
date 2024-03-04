package config

import (
	"os"
	"path/filepath"

	"gopkg.in/yaml.v2"
)

func LoadConf(filename string, s interface{}) {
	path := filepath.Join("config", filename)

	if yamlFile, err := os.ReadFile(path); err != nil {
		panic(filename + " get error: " + err.Error())
	} else if err = yaml.Unmarshal(yamlFile, s); err != nil {
		panic(filename + " unmarshal error: " + err.Error())
	}
}
