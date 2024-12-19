package manager

import "github.com/AlifiChiganjati/example-go-clean/internal/user/usecase"

type UsecaseManagerImpl interface {
	UserUseCase() usecase.UserUsecaseImpl
}

type UsecaseManager struct {
	repo RepoManagerImpl
}

func NewUseCaseManager(repo RepoManagerImpl) UsecaseManagerImpl {
	return &UsecaseManager{repo: repo}
}

func (u *UsecaseManager) UserUseCase() usecase.UserUsecaseImpl {
	return usecase.NewUserUsecase(u.repo.UserRepo())
}
