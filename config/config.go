package config

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

var CFG *Config

type Config struct {
	LogLevel string `json:"log_level"`
	Port     string `json:"port"`
	Username string `json:"username"`
	Password string `json:"password"`
	Token    string `json:"token"`
}

func LoadConfig() *Config {
	cfg := new(Config)
	buf, err := ioutil.ReadFile("config.json")
	if err != nil {
		log.Printf("fail to read config.json", err.Error())
		return nil
	}

	if err := json.Unmarshal(buf, cfg); err != nil {
		log.Printf("fail to unmarshal cfg: %s", err.Error())
		return nil
	}

	if cfg == nil || !CheckConfig(cfg) {
		log.Println("unexpect configuuration")
		return nil
	}

	return cfg
}

func CheckConfig(cfg *Config) bool {
	if cfg.Port == "" {
		log.Println("port is empty")
		return false
	}

	if cfg.Token == "" {
		log.Println("token is empty")
		return false
	}

	return true
}
