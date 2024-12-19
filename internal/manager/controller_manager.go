package manager

import "github.com/AlifiChiganjati/example-go-clean/internal/user/controller"

type ControlManagerImpl interface {
	UserController() controller.UserControllerImpl
}

type ControlManager struct {
	usecaseManager UsecaseManagerImpl
}

func NewControlManager(usecaseManager UsecaseManagerImpl) ControlManagerImpl {
	return &ControlManager{usecaseManager: usecaseManager}
}

func (c *ControlManager) UserController() controller.UserControllerImpl {
	return controller.NewUserController(c.usecaseManager.UserUseCase())
}
