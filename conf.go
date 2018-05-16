package main

import (
	"encoding/json"
	"io/ioutil"
)

type Config struct {
	Ip         string `json:"ip"`
	Port       int    `json:"port"`
	StaticPath string `json:"staticPath"`
}

func loadConf(filename string) (*Config, error) {
	conf := &Config{
		StaticPath: "./static",
	}
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(data, conf)
	if err != nil {
		return nil, err
	}
	return conf, nil
}
