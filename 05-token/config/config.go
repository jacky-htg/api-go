package config

import (
	"github.com/spf13/viper"
	"github.com/jacky-htg/api-go/05-token/libraries"
)

func init() {
	viper.SetConfigFile("./config/config.json")
	err := viper.ReadInConfig()
	libraries.CheckError(err)
}

func GetString(key string)(string) {
	return viper.GetString(key)
}

func GetInt(key string)(int) {
	return viper.GetInt(key)
}
