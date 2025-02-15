package main

import (
	"github.com/gin-gonic/gin"
	"github.com/mehedicode-lab/go-clean-architecture/config"
	"github.com/mehedicode-lab/go-clean-architecture/internal/domain"
	"github.com/mehedicode-lab/go-clean-architecture/internal/infrastructure/entities"
	"github.com/mehedicode-lab/go-clean-architecture/internal/interfaces/http"
	"github.com/mehedicode-lab/go-clean-architecture/internal/interfaces/routes"
	"github.com/mehedicode-lab/go-clean-architecture/internal/usecases"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	config.LoadConfig()
	cfg := config.AppConfig
	db := config.ConnectDatabase(cfg.Rds)
	db.AutoMigrate(&domain.User{})

	repo := entities.NewUserRepo(db)
	service := usecases.NewUserService(repo)
	handler := http.NewAuthHandler(service)

	r := gin.Default()
	routes.AuthRoutes(r.Group("/auth"), handler)

	r.Run(":8080")
}
