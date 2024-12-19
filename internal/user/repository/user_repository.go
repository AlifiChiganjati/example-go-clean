package repository

import (
	"database/sql"

	"github.com/AlifiChiganjati/example-go-clean/internal/user/domain"
)

type (
	UserRepositoryImpl interface {
		Get(id string) (domain.User, error)
	}

	UserRepository struct {
		db *sql.DB
	}
)

func NewUserRepository(db *sql.DB) UserRepositoryImpl {
	return &UserRepository{
		db: db,
	}
}

func (u *UserRepository) Get(id string) (domain.User, error) {
	var user domain.User
	query := `
  SELECT id,username,email,created_at,updated_at 
  FROM users 
  WHERE id = $1`
	err := u.db.QueryRow(query, id).Scan(&user.Id, &user.UserName, &user.Email, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		return domain.User{}, nil
	}
	return user, nil
}
