//go:build wireinject

package wire

import (
	"github.com/google/wire"
	"github.com/nghiatk54/go-ecommerce-backend-api/internal/controller"
	"github.com/nghiatk54/go-ecommerce-backend-api/internal/repo"
	"github.com/nghiatk54/go-ecommerce-backend-api/internal/service"
)

func InitUserRouterHandler() (*controller.UserController, error) {
	wire.Build(
		repo.NewUserRepository,
		service.NewUserService,
		controller.NewUserController,
	)

	return new(controller.UserController), nil
}
