package config

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/spf13/viper"
	"io/ioutil"
	"log"
	"net/http"
)


type sConfig struct{
	Name 	string	`json:"name"`
	Profiles	[]string `json:"profiles"`
	Label	interface{} `json:"label"`
	Version	interface{} `json:"version"`
	State	interface{} `json:"state"`
	PropertySource []propertySource `json:"propertySource"`
}

type lConfig struct{
	PropertySource []propertySource
}

type propertySource struct {
	Name 	string	`json:"name"`
	Source 	map[string]interface{} `json:"source"`
}

func LoadConfiguration(s string, profile string){
	if profile == DEV { var environment lConfig; localConfig(s); unmarshal(&environment) }

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

func unmarshal(v interface{}){
	err := viper.Unmarshal(v)
	if err!=nil {
		log.Fatalf("Unable parse local config")
	}
	log.Printf("Local environment configuration successfully parsed")
}

func basicAuthHeader()string{
	return base64.StdEncoding.EncodeToString([]byte(":"))
}

func remoteConfig()([]byte, error){
	req, err := http.NewRequest(http.MethodGet, "", nil)
	req.Header.Add("Authorization", "Basic " + basicAuthHeader())

	resp, err := Client.Do(req)
	if err!= nil {
		panic("Failed reading remote configuration. Shutting down:" + err.Error())
	}
	body, err := ioutil.ReadAll(resp.Body)
	return body, err
}

func parseRemote(body []byte){

	var environment sConfig
	err := json.Unmarshal(body, &environment)
	if err !=nil { panic("Cannot parse configuration, message: " + err.Error())}

	for key, value := range environment.PropertySource[0].Source {
		viper.Set(key, value)
		fmt.Printf("Loading config property %v => %v\n",key, value)
	}
	if viper.IsSet("server_name"){
		fmt.Printf("Successfully loaded configuration for service %s\n", viper.GetString("server_name"))
	}



}
