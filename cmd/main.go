package main

import (
	"gin_restfull/config"
	"gin_restfull/internal/domain"
	"gin_restfull/internal/handler"
	"gin_restfull/internal/repository"
	"gin_restfull/internal/usecase"
	"gin_restfull/pkg/middleware"

	"github.com/gin-gonic/gin"
)

// @title User Auth API
// @version 1.0
// @description Backend sederhana untuk login, register, JWT auth, dan role-based access.
// @host localhost:3000
// @BasePath /api
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
func main() {
	db := config.InitPostgres()
	defer config.CloseDB(db)
	db.AutoMigrate(&domain.User{})

	r := gin.Default()

	repo := repository.NewUserRepository(db)
	uc := usecase.NewUserUseCase(repo)
	h := handler.NewUserHandler(uc)

	api := r.Group("/api")
	{
		api.POST("/register", h.Register)
		api.POST("/login", h.Login)
		api.POST("/refresh", h.RefreshToken)
		api.GET("/profile", middleware.JWTProtected("admin"), h.Profile)
	}

	r.Run(":3000")
}
