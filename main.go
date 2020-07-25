package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/cdugga/microservices_with_go/config"
	"github.com/cdugga/microservices_with_go/controllers"
	"github.com/go-openapi/runtime/middleware"
	gohandlers "github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/spf13/viper"
	"log"
	"net/http"
	"os"
	"os/signal"
	"path/filepath"
	"runtime"
	"time"
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

func shutdownHook(s http.Server, l *log.Logger){
	sigChannel := make(chan os.Signal)
	signal.Notify(sigChannel, os.Interrupt)
	signal.Notify(sigChannel, os.Kill)

	sig := <- sigChannel
	l.Println("Received terminate, gracefully shutting down", sig)

	tc, _ := context.WithTimeout(context.Background(), 30*time.Second)
	s.Shutdown(tc)
}

func startServer(s http.Server, l *log.Logger){
	go func(){
		err := s.ListenAndServe()
		if err != nil{
			l.Fatal(err)
		}
		l.Println("Started server")
	}()
}

func createServer(sm *mux.Router)http.Server{
	ch := gohandlers.CORS(gohandlers.AllowedOrigins([]string{"*"}))

	return http.Server{
		Addr: ":8080",
		Handler: ch(sm),
		IdleTimeout: 120*time.Second,
		ReadTimeout: 5*time.Second,
		WriteTimeout: 10*time.Second,
	}
}

func createServerMux() *mux.Router {
	sm := mux.NewRouter()
	getRouter := sm.Methods(http.MethodGet).Subrouter()
	getRouter.HandleFunc("/cryptoprices", controllers.HelloWorldHandler)

	opts := middleware.RedocOpts{SpecURL: "/cryptoprices/swagger.json"}
	sh := middleware.Redoc(opts, nil)

	getRouter.Handle("/docs", sh)
	return sm
}

func main(){

	logger := log.New(os.Stdout, "demo.api", log.LstdFlags)

	config.LoadConfiguration(basepath, viper.GetString("profile"), logger)
	fmt.Println(viper.IsSet("propertySources.source.endpointa"))

	sm := createServerMux()
	s := createServer(sm)

	startServer(s, logger)
	shutdownHook(s, logger)
}

