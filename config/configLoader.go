package config

import (
	"fmt"
	"github.com/spf13/viper"
	"log"
	"os"
)

type lConfig struct{
	PropertySource []propertySource
}

type propertySource struct {
	Name 	string
	Source 	map[string]interface{}
}

func LoadConfiguration(s string){
	localConfig(s)

}

func localConfig(s string){
	viper.SetConfigName("application-local")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(s)

	path, _ := os.Getwd()
	fmt.Println("---------", path)
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			panic("Config file was not found: " + err.Error())
		}else {
			panic("Config file found but another error was encountered: " + err.Error())
		}
	}

	var environment lConfig
	err := viper.Unmarshal(&environment)
	if err!=nil {
		log.Fatalf("Unable parse local config")
	}
	log.Printf("Local environment configuration successfully parsed")
}


