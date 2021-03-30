package config

import (
	"github.com/KirillMironov/go-server/domain"
	"github.com/spf13/viper"
	"os"
)

var Config domain.Config

func LoadConfiguration() error {
	viper.SetConfigName("service")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.AddConfigPath("../config")

	err := viper.ReadInConfig()
	if err != nil {
		Config.Port = os.Getenv("Port")
		Config.Database.ConnectionString = os.Getenv("ConnectionString")
		Config.Security.JWTKey = os.Getenv("JWTKey")
		return nil
	}

	err = viper.Unmarshal(&Config)
	if err != nil {
		return err
	}

	return nil
}
