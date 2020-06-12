package glb

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"strings"
)

type Config struct {
	Database struct {
		Host     string `json:"host"`
		Port     int    `json:"port"`
		User     string `json:"user"`
		Password string `json:"password"`
		DBname   string `json:"dbname"`
	} `json:"database"`
	Domain    string `json:"domain"`
	FileStore string `json:"file_store"`
}

func (c *Config) Read() {
	conf, err := ioutil.ReadFile("./settings.json")
	if err != nil {
		panic(err)
	}
	if err := json.Unmarshal(conf, &c); err != nil {
		panic(err)
	}
	
	if strings.Contains(c.FileStore, "~") {
		homeDir, err := os.UserHomeDir()
		if err != nil {
			panic(err)
		}
		strings.ReplaceAll(c.FileStore, "~", homeDir)
	}
}

func Read() Config {
	var c Config
	conf, _ := ioutil.ReadFile("./settings.json")
	json.Unmarshal(conf, &c)
	
	if strings.Contains(c.FileStore, "~") {
		homeDir, err := os.UserHomeDir()
		if err != nil {
			panic(err)
		}
		c.FileStore = strings.ReplaceAll(c.FileStore, "~", homeDir)
	}
	
	return c
}
