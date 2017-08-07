package config

import (
	"github.com/spf13/viper"
	"fmt"
)

func init() {
	viper.SetConfigFile("./config/config.json")
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println(err.Error())
		panic(err.Error())
	}
}

func GetString(key string)(string) {
	return viper.GetString(key)
}

func GetInt(key string)(int) {
	return viper.GetInt(key)
}