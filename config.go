package main

import (
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"
)

var config Config = Config{}

type AdapterConfig struct {
	Name      string
	Title     string
	Type      string
	Arguments map[string]string
}

type Config struct {
	Host     string
	Port     string
	Adapters []AdapterConfig
}

func loadConfigFile(filename string) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	err = yaml.Unmarshal([]byte(data), &config)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
}
