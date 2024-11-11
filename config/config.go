package config

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	SourceWebBuffer   uint32
	SourceApiBuffer   uint32
	SourceOtherBuffer uint32
}

var AppConfig Config

func init() {
	viper.SetConfigFile("go-commons/config/.env") // Adjust path to the .env file
	viper.AutomaticEnv()

	viper.SetDefault("SOURCE_WEB_BUFFER", 900)
	viper.SetDefault("SOURCE_API_BUFFER", 1200)
	viper.SetDefault("SOURCE_OTHER_BUFFER", 300)

	if err := viper.ReadInConfig(); err != nil {
		log.Printf("Error reading config file: %v", err)
	}

	AppConfig = Config{
		SourceWebBuffer:   viper.GetUint32("SOURCE_WEB_BUFFER"),
		SourceApiBuffer:   viper.GetUint32("SOURCE_API_BUFFER"),
		SourceOtherBuffer: viper.GetUint32("SOURCE_OTHER_BUFFER"),
	}
	log.Println("Configuration loaded:", AppConfig)
}
