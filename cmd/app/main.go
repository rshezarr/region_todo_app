package main

import (
	"flag"
	"github.com/spf13/viper"
	"log"
	"todo_list/internal/app"
)

// @title Todo App API
// @description API for managing todo lists.
// @version 1.0
// @host localhost:9090
// @BasePath /v1
// @schemes http
// @schemes https
// @produces json
// @consumes json

func main() {
	if err := initConfig(); err != nil {
		log.Fatalf("error occured while parsing configs: %v", err)
	}

	app.Run()
}

func initConfig() error {
	//set flag
	var configPath = flag.String("config-path", "configs/", "path to config file")

	//parse flag
	flag.Parse()

	//set config file path, name and type
	viper.SetConfigName("configs")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(*configPath)

	//read config by property above
	return viper.ReadInConfig()
}
