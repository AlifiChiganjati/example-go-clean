package delivery

import (
	"fmt"
	"log"

	"github.com/AlifiChiganjati/example-go-clean/config"
	"github.com/AlifiChiganjati/example-go-clean/internal/manager"
	"github.com/gin-gonic/gin"
)

type (
	ServerImpl interface {
		setupRoutes()
		Run()
	}

	Server struct {
		control manager.ControlManagerImpl
		engine  *gin.Engine
		host    string
	}
)

func NewServer() ServerImpl {
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatal(err)
	}

	cfgIntance, _ := cfg.(*config.Config)

	infra, err := manager.NewInfraManager(cfgIntance)
	if err != nil {
		log.Fatal(err)
	}

	repo := manager.NewRepoManager(infra)
	uc := manager.NewUseCaseManager(repo)
	control := manager.NewControlManager(uc)
	host := fmt.Sprintf(":%s", cfgIntance.ApiPort)
	engine := gin.Default()

	return &Server{
		control: control,
		engine:  engine,
		host:    host,
	}
}

func (s *Server) setupRoutes() {
	s.engine.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "Welcome to the API!"})
	})

	s.engine.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "OK"})
	})
}

func (s *Server) Run() {
	s.setupRoutes()
	if err := s.engine.Run(s.host); err != nil {
		log.Fatal("server can't run")
	}
}
