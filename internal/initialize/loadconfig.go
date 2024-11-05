package initialize

import (
	"fmt"

	"github.com/nghiatk54/go-ecommerce-backend-api/global"
	"github.com/spf13/viper"
)

func LoadConfig() {
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
	if err := viper.Unmarshal(&global.Config); err != nil {
		fmt.Printf("Unalble to decode config %v", err)
	}
}
