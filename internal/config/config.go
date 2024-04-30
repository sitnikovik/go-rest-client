package config

import (
	"gopkg.in/yaml.v3"
	"log"
	"os"
	"path"
	"runtime"
)

const configFilename = "../../config.yaml"

// Config is a struct to store configuration to test library
type Config struct {
	Proxy struct {
		Host     string `yaml:"host"`
		Port     int    `yaml:"port"`
		Login    string `yaml:"login"`
		Password string `yaml:"password"`
	} `yaml:"proxy"`
}

// FromFile reads config from file
func FromFile() Config {
	file, err := os.Open(currentDir() + "/" + configFilename)
	if err != nil {
		log.Fatal(err.Error())
	}
	defer func(file *os.File) {
		_ = file.Close()
	}(file)

	var conf Config
	if file != nil {
		decoder := yaml.NewDecoder(file)
		if err = decoder.Decode(&conf); err != nil {
			log.Fatal(err.Error())
		}
	}

	return conf
}

// currentDir returns path to current directory
func currentDir() string {
	_, b, _, _ := runtime.Caller(0)

	return path.Dir(b)
}
