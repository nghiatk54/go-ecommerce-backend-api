package initialize

import (
	"github.com/nghiatk54/go-ecommerce-backend-api/global"
	"github.com/nghiatk54/go-ecommerce-backend-api/pkg/logger"
)

func InitLogger() {
	global.Logger = logger.NewLogger(global.Config.Logger)
}
