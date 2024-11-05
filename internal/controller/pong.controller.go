package controller

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type PongController struct{}

func NewPongController() *PongController {
	return &PongController{}
}

func (pc *PongController) Pong(c *gin.Context) {
	fmt.Println("----------> My handle")
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
