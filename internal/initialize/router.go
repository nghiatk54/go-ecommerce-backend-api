package initialize

import (
	"fmt"

	"github.com/gin-gonic/gin"
	c "github.com/nghiatk54/go-ecommerce-backend-api/internal/controller"
	"github.com/nghiatk54/go-ecommerce-backend-api/internal/middlewares"
)

func AA() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("Before --->AA")
		c.Next()
		fmt.Println("After --->AA")
	}
}

func BB() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("Before --->BB")
		c.Next()
		fmt.Println("After --->BB")
	}
}

func CC(c *gin.Context) {
	fmt.Println("Before --->CC")
	c.Next()
	fmt.Println("After --->CC")
}

func InitRouter() *gin.Engine {
	r := gin.Default()
	// use middleware
	r.Use(middlewares.AuthenMiddleware(), BB(), CC)

	v1 := r.Group("/v1/2024")
	{
		v1.GET("/ping/:name", c.NewPongController().Pong) // v1/2024/ping
		v1.GET("/user/1", c.NewUserController().GetUsersByID)
		// v1.POST("/ping", Pong)
		// v1.PUT("/ping", Pong)
		// v1.DELETE("/ping", Pong)
		// v1.PATCH("/ping", Pong)
		// v1.OPTIONS("/ping", Pong)
		// v1.HEAD("/ping", Pong)
	}

	// v2 := r.Group("/v2/2024")
	// {
	// 	v2.GET("/ping", Pong) // v2/2024/ping
	// 	v2.POST("/ping", Pong)
	// 	v2.PUT("/ping", Pong)
	// 	v2.DELETE("/ping", Pong)
	// 	v2.PATCH("/ping", Pong)
	// 	v2.OPTIONS("/ping", Pong)
	// 	v2.HEAD("/ping", Pong)
	// }

	return r
}
