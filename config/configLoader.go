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

func LoadConfiguration(){
	localConfig()
}

func localConfig(){
	viper.SetConfigName("application-local")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")

	path, _ := os.Getwd()
	fmt.Println("---------", path)
	if err := viper.ReadInConfig(); err != nil {
		panic("Error reading local config file: " + err.Error())
	}

	var environment lConfig
	err := viper.Unmarshal(&environment)
	if err!=nil {
		log.Fatalf("Unable parse local config")
	}
	log.Printf("Local environment configuration successfully parsed")
}


