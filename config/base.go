package config

import (
	"fmt"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v2"
)

func LoadConf(filename string, s interface{}) {
	var path string
	path = filepath.Join("config", filename)
	yamlFile, _ := os.ReadFile(path)
	fmt.Println(string(yamlFile))
	yaml.Unmarshal(yamlFile, s)
	fmt.Println(s)

	if yamlFile, err := os.ReadFile(path); err != nil {
		panic(filename + " get error: " + err.Error())
	} else if err = yaml.Unmarshal(yamlFile, s); err != nil {
		panic(filename + " unmarshal error: " + err.Error())
	}
}
