package manager

import "github.com/AlifiChiganjati/example-go-clean/internal/user/repository"

type RepoManagerImpl interface {
	UserRepo() repository.UserRepositoryImpl
}

type RepoManager struct {
	infra InfraManagerImpl
}

func NewRepoManager(infra InfraManagerImpl) RepoManagerImpl {
	return &RepoManager{infra: infra}
}

func (r *RepoManager) UserRepo() repository.UserRepositoryImpl {
	return repository.NewUserRepository(r.infra.Connect())
}
