package routers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	r := gin.Default()

	v1 := r.Group("/v1/2024")
	{
		v1.GET("/ping/:name", Pong) // v1/2024/ping
		v1.POST("/ping", Pong)
		v1.PUT("/ping", Pong)
		v1.DELETE("/ping", Pong)
		v1.PATCH("/ping", Pong)
		v1.OPTIONS("/ping", Pong)
		v1.HEAD("/ping", Pong)
	}

	v2 := r.Group("/v2/2024")
	{
		v2.GET("/ping", Pong) // v2/2024/ping
		v2.POST("/ping", Pong)
		v2.PUT("/ping", Pong)
		v2.DELETE("/ping", Pong)
		v2.PATCH("/ping", Pong)
		v2.OPTIONS("/ping", Pong)
		v2.HEAD("/ping", Pong)
	}

	return r
}

func Pong(c *gin.Context) {
	name := c.Param("name")
	uid := c.Query("uid")
	email := c.DefaultQuery("email", "nghiatk54@gmail.com")

	c.JSON(http.StatusOK, gin.H{ // map string
		"message": "pong...ping",
		"name":    name,
		"uid":     uid,
		"email":   email,
		"users":   []string{"John", "Messi", "CR7", "nghiatk54"},
	})
}
