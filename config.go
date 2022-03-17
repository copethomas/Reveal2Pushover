package main

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

const (
	ConfigFileName = "reveal2pushover_config.json"
)

type Config struct {
	ServerAddr      string
	PushOverUserKey string
	PushOverAppKey  string
	TenantId        string
}

func ConfigExists() bool {
	if _, err := os.Stat(ConfigFileName); err != nil {
		return false
	}
	return true
}

func WriteConfig() {
	c := Config{"127.0.0.1:8080", "user_key", "app_key", "tenant_id"}
	jsonData, _ := json.MarshalIndent(c, "", "	")
	err := ioutil.WriteFile(ConfigFileName, jsonData, 0600)
	if err != nil {
		Error.Fatal("Failed to create config file")
	}
}

func LoadConfig() (Config, error) {
	var c Config
	jsonData, err := ioutil.ReadFile(ConfigFileName)
	if err != nil {
		return c, err
	}
	err = json.Unmarshal(jsonData, &c)
	return c, err
}
