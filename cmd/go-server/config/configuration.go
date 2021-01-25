package config

import (
	"github.com/spf13/viper"
)

var Config DBConfig

type DBConfig struct {
	Database struct {
		ConnectionString string
	}
	Security struct {
		JWTKey string
	}
}

func ReadConfig() error {
	v := viper.New()
	v.SetConfigName("service")
	v.SetConfigType("yaml")
	v.AddConfigPath("../../../config")
	v.AddConfigPath("../../config")
	v.SetEnvPrefix("go-server")
	v.AutomaticEnv()

	err := v.ReadInConfig()
	if err != nil {
		return err
	}
	return v.Unmarshal(&Config)
}
