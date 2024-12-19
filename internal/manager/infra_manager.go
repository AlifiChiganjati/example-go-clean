package manager

import (
	"database/sql"
	"fmt"

	"github.com/AlifiChiganjati/example-go-clean/config"
	_ "github.com/lib/pq"
)

type (
	InfraManagerImpl interface {
		Connect() *sql.DB
		openConn() error
	}

	InfraManager struct {
		db  *sql.DB
		cfg *config.Config
	}
)

func NewInfraManager(cfg *config.Config) (InfraManagerImpl, error) {
	conn := &InfraManager{
		cfg: cfg,
	}

	if err := conn.openConn(); err != nil {
		return nil, err
	}

	return conn, nil
}

func (i *InfraManager) openConn() error {
	dns := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", i.cfg.Host, i.cfg.Port, i.cfg.User, i.cfg.Password, i.cfg.Name)

	db, err := sql.Open(i.cfg.Driver, dns)
	if err != nil {
		return fmt.Errorf("failed to open connection %v", err.Error())
	}
	i.db = db
	return nil
}

func (i *InfraManager) Connect() *sql.DB {
	return i.db
}
