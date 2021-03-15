package config

import (
	"github.com/spf13/viper"
	"os"
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

	err := v.ReadInConfig()
	if err != nil {
		Config.Database.ConnectionString = os.Getenv("ConnectionString")
		Config.Security.JWTKey = os.Getenv("JWTKey")
		return nil
	}

	return v.Unmarshal(&Config)
}
