package usecase

import (
	"fmt"

	"github.com/AlifiChiganjati/example-go-clean/internal/user/domain"
	"github.com/AlifiChiganjati/example-go-clean/internal/user/repository"
)

type (
	UserUsecaseImpl interface {
		FindById(id string) (domain.User, error)
	}

	UserUsecase struct {
		repo repository.UserRepositoryImpl
	}
)

func NewUserUsecase(repo repository.UserRepositoryImpl) UserUsecaseImpl {
	return &UserUsecase{
		repo: repo,
	}
}

func (u *UserUsecase) FindById(id string) (domain.User, error) {
	user, err := u.repo.Get(id)
	if err != nil {
		return domain.User{}, fmt.Errorf("user with ID %s not found", id)
	}
	return user, nil
}
