package initialize

import (
	"fmt"

	"github.com/nghiatk54/go-ecommerce-backend-api/global"
	"go.uber.org/zap"
)

func Run() {
	// Load configuration
	LoadConfig()
	m := global.Config.Mysql
	fmt.Println("Loading configuration mysql", m.UserName, m.Password)
	InitLogger()
	global.Logger.Info("Config log ok!!", zap.String("ok", "success"))
	InitMysql()
	InitRedis()
	r := InitRouter()
	r.Run(":8002")
}
