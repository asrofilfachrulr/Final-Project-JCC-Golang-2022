package config

import (
	"anya-day/routes"
	"anya-day/validation"
	"log"

	"github.com/gin-gonic/gin"
	"gopkg.in/go-playground/validator.v9"
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

	v := validator.New()
	validation.RegisterAll(v)

	routes.Attach(s.Router, "db", db)
	routes.Attach(s.Router, "validator", v)

	routes.InitAPI(s.Router)
	log.Println("Server initiated...")
}

func (s *Server) Run() {
	sqlDB, _ := s.DB.DB()
	defer sqlDB.Close()

	log.Println("Running server at 8080...")
	s.Router.Run()
}
