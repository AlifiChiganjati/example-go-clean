package controller

import (
	"net/http"

	"github.com/AlifiChiganjati/example-go-clean/internal/user/usecase"
	"github.com/AlifiChiganjati/example-go-clean/pkg/response"
	"github.com/gin-gonic/gin"
)

type (
	UserControllerImpl interface {
		GetByIdHanlder(c *gin.Context)
	}

	UserController struct {
		uc usecase.UserUsecaseImpl
	}
)

func NewUserController(uc usecase.UserUsecaseImpl) UserControllerImpl {
	return &UserController{uc: uc}
}

func (ucon *UserController) GetByIdHanlder(c *gin.Context) {
	id := c.Param("id")

	data, err := ucon.uc.FindById(id)
	if err != nil {
		response.SendErrorResponse(c, http.StatusBadRequest, err.Error())
	}

	response.SendSingleResponse(c, "succes", data)
}
