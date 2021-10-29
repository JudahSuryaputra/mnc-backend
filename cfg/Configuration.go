package cfg

import (
	"log"

	"github.com/spf13/viper"
)

func Init() {
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./cfg")
	viper.SetConfigName("local")

	if err := viper.ReadInConfig(); err != nil {
		log.Printf("Fatal error config file: %s\n", err)
	}
}
