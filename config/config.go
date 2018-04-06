package config

import (
	"log"
	"os"

	"github.com/spf13/viper"
)

// Viper exposes the Viper configuration
var Viper *viper.Viper

func init() {
	Viper = viper.New()
	Viper.SetConfigName("settings")     // name of config file (without extension)
	Viper.AddConfigPath("/etc/config/") // path to look for the config file in
	Viper.AddConfigPath("config")
	Viper.AddConfigPath("../config")             // call multiple times to add many search paths
	if err := Viper.ReadInConfig(); err != nil { // Handle errors reading the config file
		log.Fatalf("Fatal error config file: %s \n", err)
	}
}

// GetHostPort returns the host+port string for the server
func GetHostPort() string {
	host := Viper.GetString("app.host")
	p := Viper.GetString("app.port")
	if p == "" {
		p = os.Getenv("PORT")
	}
	return host + ":" + p
}

// GetStaticPath returns the path string for static files
func GetStaticPath() string {
	return Viper.GetString("app.StaticPath")
}
