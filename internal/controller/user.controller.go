package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/nghiatk54/go-ecommerce-backend-api/internal/service"
	"github.com/nghiatk54/go-ecommerce-backend-api/pkg/response"
)

type UserController struct {
	userService service.IUserService
}

func NewUserController(
	userService service.IUserService,
) *UserController {
	return &UserController{
		userService: userService,
	}
}

func (uc *UserController) Register(c *gin.Context) {
	result := uc.userService.Register("", "")
	response.SuccessResponse(c, result, nil)
}

// // controller -> service -> repo -> models -> dbs
// func (uc *UserController) GetUsersByID(c *gin.Context) {
// 	response.SuccessResponse(c, 20001, []string{"tipjs", "m10", "cr7"})
// 	// response.ErrorResponse(c, 20003)
// }
