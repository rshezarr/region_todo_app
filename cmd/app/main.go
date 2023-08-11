package main

import (
	"flag"
	"github.com/spf13/viper"
	"log"
	"todo_list/internal/app"
)

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
