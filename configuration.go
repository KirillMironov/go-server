package main

import (
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"log"
)

type Conf struct {
	Database struct {
		ConnectionString string `yaml:"ConnectionString"`
	} `yaml:"Database"`
}

func ReadConfiguration(filename string, conf *Conf) {
	file, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Printf("Unable to read yaml")
	}

	err = yaml.Unmarshal(file, conf)
	if err != nil {
		log.Printf("Unmarshal error")
	}
}
