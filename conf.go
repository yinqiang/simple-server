package main

import (
	"encoding/gob"
	"os"
)

type Config struct {
	Ip   string
	Port int
	Path string
}

func LoadConfigOrDefault(filepath, defaultPath, defaultHost string, defaultPort int) (*Config, error) {
	conf, err := LoadConfig(filepath)
	if err != nil {
		if os.IsNotExist(err) {
			conf = &Config{
				Path: defaultPath,
				Ip:   defaultHost,
				Port: defaultPort,
			}
			if err := conf.Save(filepath); err != nil {
				return nil, err
			}
		} else {
			return nil, err
		}
	}
	return conf, nil
}

func LoadConfig(filepath string) (*Config, error) {
	f, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	d := gob.NewDecoder(f)
	c := &Config{}
	if err := d.Decode(c); err != nil {
		return nil, err
	}
	return c, nil
}

func (c *Config) Save(filepath string) error {
	f, err := os.OpenFile(filepath, os.O_CREATE|os.O_WRONLY, os.ModePerm)
	if err != nil {
		return err
	}
	defer f.Close()
	e := gob.NewEncoder(f)
	return e.Encode(c)
}
