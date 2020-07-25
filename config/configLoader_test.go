package config

import (
	"github.com/spf13/viper"
	"testing"
)

func TestLocalConfig(t *testing.T){

	LoadConfiguration("../")
	s := viper.IsSet("propertySources.source.endpointa")
	if s != true {
		t.Fatal("Expected param not found")
	}
	t.Log("Fetched param successfully")
}
