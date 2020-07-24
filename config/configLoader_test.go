package config

import (
	"github.com/spf13/viper"
	"testing"
)

func TestLocalConfig(t *testing.T){
	LoadConfiguration()

	s := viper.GetString("propertySources.source.endpointa")
	if len(s) <= 0 {
		t.Fatal("Expected param not found")
	}
	t.Log("Fetched param successfully")
}
