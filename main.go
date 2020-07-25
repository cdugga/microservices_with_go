package main

import (
	"flag"
	"fmt"
	"github.com/cdugga/microservices_with_go/config"
	"github.com/spf13/viper"
	"path/filepath"
	"runtime"
)

var (
	_, b, _, _ = runtime.Caller(0)
	basepath   = filepath.Dir(b)
)

func init(){
	profile := flag.String("profile", "dev", "Choose operational mode")
	flag.Parse()
	viper.Set("profile", *profile)
}

func main(){
	config.LoadConfiguration(basepath, viper.GetString("profile"))

	fmt.Println(viper.IsSet("propertySources.source.endpointa"))
}

