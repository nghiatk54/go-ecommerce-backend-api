package global

import (
	"github.com/nghiatk54/go-ecommerce-backend-api/pkg/logger"
	"github.com/nghiatk54/go-ecommerce-backend-api/pkg/setting"
)

var (
	Config setting.Config
	Logger *logger.LoggerZap
	// Mdb    *gorm.DB
)
