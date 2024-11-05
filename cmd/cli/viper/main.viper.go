package main

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	Server struct {
		Port int `mapstructure:"port"`
	} `mapstructure:"server"`
	Databases []struct {
		User     string `mapstructure:"user"`
		Password string `mapstructure:"password"`
		Host     string `mapstructure:"host"`
		DbName   string `mapstructure:"dbName"`
	} `mapstructure:"databases"`
}

func main() {
	viper := viper.New()
	viper.AddConfigPath("./config") // path to config
	viper.SetConfigName("local")    // name of config file (without extension)
	viper.SetConfigType("yaml")     // extension of config file

	// Read config
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("failed to read config file: %w", err))
	}
	// Read server config
	fmt.Println("Server port::", viper.GetInt("server.port"))
	fmt.Println("Security jwt key::", viper.GetString("security.jwt.key"))
	// Config struct
	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		fmt.Printf("Unalble to decode config %v", err)
	}

	fmt.Println("Config port::", config.Server.Port)
	for _, db := range config.Databases {
		fmt.Printf("database User: %s, Password: %s, Host: %s \n", db.User, db.Password, db.Host)
	}

}
