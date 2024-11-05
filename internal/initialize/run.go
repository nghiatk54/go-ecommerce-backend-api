package initialize

import (
	"fmt"

	"github.com/nghiatk54/go-ecommerce-backend-api/global"
)

func Run() {
	// Load configuration
	LoadConfig()
	m := global.Config.Mysql
	fmt.Println("Loading configuration mysql", m.Username, m.Password)
	InitLogger()
	InitMysql()
	InitRedis()
	r := InitRouter()
	r.Run(":8002")
}
