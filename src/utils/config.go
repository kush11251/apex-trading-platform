package utils

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

type Config struct {
	ServerAddr string `json:"server_addr"`
}

var config Config

func InitConfig() {
	configBytes, err := ioutil.ReadFile("config.json")
	if err != nil {
		log.Fatal(err)
	}
	json.Unmarshal(configBytes, &config)
}
