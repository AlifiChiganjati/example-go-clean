package v1

import (
	"github.com/AlifiChiganjati/example-go-clean/internal/user/controller"
	"github.com/gin-gonic/gin"
)

type UserRouter struct {
	ucon controller.UserController
	rg   *gin.RouterGroup
}

func NewUserRoutes(ucon controller.UserController, rg *gin.RouterGroup) *UserRouter {
	return &UserRouter{
		ucon: ucon,
		rg:   rg,
	}
}

func (ur *UserRouter) Route() {
	user := ur.rg.Group("/user")
	user.POST("/register", ur.ucon.GetByIdHanlder)
}
