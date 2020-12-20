package main

import (
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"log"
	"testing"
)

type Conf struct {
	ConnectionString string `yaml:"ConnectionString"`
}

func TestRead(t *testing.T)  {
	var conf Conf

	file, err := ioutil.ReadFile("service.yaml")
	if err != nil {
		log.Printf("err")
	}

	err = yaml.Unmarshal(file, &conf)
	if err != nil {
		log.Printf("Unmarshal error")
	}
	log.Printf(conf.ConnectionString)
}
