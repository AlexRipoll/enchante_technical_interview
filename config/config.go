package config

import (
	"fmt"
	"github.com/spf13/viper"
)

var (
	Params *Configurations
)

func init() {
	Params = getEnvVariables()
}

type Configurations struct {
	Server   Server
	Database Database
	Jwt      Jwt
}

type Server struct {
	Port int
}

type Database struct {
	Name     string
	User     string
	Password string
	Host     string
	Port     int
}

type Jwt struct {
	Key string
	Ttl int64 // in seconds
}

func getEnvVariables() *Configurations {
	viper.SetConfigName("config")
	viper.SetConfigType("yml")
	viper.AddConfigPath("./")
	viper.AutomaticEnv()
	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	var configuration Configurations
	if err := viper.Unmarshal(&configuration); err != nil {
		fmt.Printf("Unable to decode into struct, %v", err)
	}
	return &configuration
}
