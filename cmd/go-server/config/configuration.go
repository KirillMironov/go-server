package config

import (
	"github.com/spf13/viper"
)

var Config DBConfig

type DBConfig struct {
	Database struct {
		ConnectionString string `yaml:"ConnectionString"`
	} `yaml:"Database"`
}

func ReadConfig() error {
	v := viper.New()
	v.SetConfigName("service")
	v.SetConfigType("yaml")
	v.AddConfigPath("../../../config")
	v.SetEnvPrefix("go-server")
	v.AutomaticEnv()
	if err := v.ReadInConfig(); err != nil {
		return err
	}
	return v.Unmarshal(&Config)
}
