package main

import (
	"flag"
	"fmt"
	"github.com/cdugga/microservices_with_go/config"
	"github.com/spf13/viper"
)


func init(){
	profile := flag.String("profile", "local", "Choose operational mode")

	flag.Parse()
	viper.Set("profile", *profile)
}



func main(){
	fmt.Println("Hello")
	config.LoadConfiguration()
}

