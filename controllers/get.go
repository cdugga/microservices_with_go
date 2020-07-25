package controllers

import (
	"fmt"
	"net/http"
)

func HelloWorldHandler(rw http.ResponseWriter, r *http.Request){
	fmt.Println("inside handler")
}
