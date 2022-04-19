package config

import (
	"anya-day/routes"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Server struct {
	Router *gin.Engine
	DB     *gorm.DB
}

func NewServer() *Server {
	return &Server{}
}

func (s *Server) Init() {
	db := ConnectDataBase()
	s.DB = db

	s.Router = gin.Default()
	routes.Attach(s.Router, "db", db)
	routes.InitAPI(s.Router)
}

func (s *Server) Run() {
	sqlDB, _ := s.DB.DB()
	defer sqlDB.Close()

	s.Router.Run()
}
