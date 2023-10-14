package main

import (
	"github.com/BurntSushi/toml"
	"log"
	"os"
	"strings"
)

type Config struct {
	Endpoint string `toml:"endpoint"`
}

func isExistConfigFile(path string) bool {
	entries, err := os.ReadDir(path)
	if err != nil {
		log.Fatal(err)
	}

	for _, e := range entries {
		if e.Name() == "nlog.config.toml" {
			return true
		}
	}
	return false
}

func getConfig() Config {
	configPath := getConfigPath()
	if configPath == "" {
		log.Fatal("nlog.config.toml not found")
	}

	var conf Config
	_, err := toml.DecodeFile(configPath, &conf)
	if err != nil {
		panic(err)
	}
	return conf
}

func getConfigPath() string {
	workDir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
		return ""
	}
	part := strings.Split(strings.ReplaceAll(workDir, "\\", "/"), "/")
	length := len(part)
	for i := 0; i < length; i++ {
		currentPath := strings.Join(part, "/")
		if isExistConfigFile(currentPath) {
			return currentPath + "/nlog.config.toml"
		}
		part = part[:len(part)-1]
	}
	return ""
}
