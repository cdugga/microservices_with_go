package main

import (
	"flag"
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
	profile := flag.String("profile", "local", "Choose operational mode")
	flag.Parse()
	viper.Set("profile", *profile)
}

func main(){
	config.LoadConfiguration(basepath, profile)
}

