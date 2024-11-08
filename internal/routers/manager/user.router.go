package manager

import "github.com/gin-gonic/gin"

type UserRouter struct{}

func (ur *UserRouter) InitUserRouter(Router *gin.RouterGroup) {
	// public router
	// userRouterPublic := Router.Group("/admin")
	// {
	// 	userRouterPublic.POST("/login")
	// }
	// private router
	userRouterPrivate := Router.Group("/admin/user")
	// userRouterPrivate.Use(limiter())
	// userRouterPrivate.Use(Authen())
	// useRouterPrivate.Use(Permission())
	{
		userRouterPrivate.POST("/active_user")
	}
}
