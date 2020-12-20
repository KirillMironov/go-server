package main

import (
	"testing"
)

func TestRead(t *testing.T)  {
	var conf Conf
	ReadConfiguration("service.yaml", &conf)

	if len(conf.Database.ConnectionString) == 0 {
		t.Fatal("Unable to read yaml")
	}
}
