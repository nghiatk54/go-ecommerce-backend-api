package user

import (
	"github.com/gin-gonic/gin"
	"github.com/nghiatk54/go-ecommerce-backend-api/internal/wire"
)

type UserRouter struct{}

func (ur *UserRouter) InitUserRouter(Router *gin.RouterGroup) {
	// public router
	// This is non-dependency
	// userRepo := repo.NewUserRepository()
	// userService := service.NewUserService(userRepo)
	// userHandleNonDependency := controller.NewUserController(userService)
	// userRouterPublic := Router.Group("/user")
	userController, _ := wire.InitUserRouterHandler()

	// WIRE
	// Dependency Injection
	userRouterPublic := Router.Group("/user")
	{
		userRouterPublic.POST("/register", userController.Register) // register -> YES -> NO
		userRouterPublic.POST("/otp")
	}

	// private router
	userRouterPrivate := Router.Group("/user")
	// userRouterPrivate.Use(limiter())
	// userRouterPrivate.Use(Authen())
	// userRouterPrivate.Use(Permission())
	{
		userRouterPrivate.GET("/get_info")
	}

}
