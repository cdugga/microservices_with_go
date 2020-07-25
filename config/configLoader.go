package config

import (
	"github.com/spf13/viper"
	"log"
)

type lConfig struct{
	PropertySource []propertySource
}

type propertySource struct {
	Name 	string
	Source 	map[string]interface{}
}

func LoadConfiguration(s string, profile string){
	if profile == DEV{ localConfig(s); unmarshal(); }
}

func localConfig(s string){
	viper.SetConfigName("application-local")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(s)

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			panic("Config file was not found: " + err.Error())
		}else {
			panic("Config file found but another error was encountered: " + err.Error())
		}
	}
}

func unmarshal(){
	var environment lConfig
	err := viper.Unmarshal(&environment)
	if err!=nil {
		log.Fatalf("Unable parse local config")
	}
	log.Printf("Local environment configuration successfully parsed")
}

