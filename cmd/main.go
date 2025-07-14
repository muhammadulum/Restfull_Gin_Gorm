// @title        My API
// @version      1.0
// @description  This is a sample server
// @host         localhost:3000
// @BasePath     /api

package main

import (
	"gin_restfull_api/config"
	"gin_restfull_api/internal/domain"
	"gin_restfull_api/internal/handler"
	"gin_restfull_api/internal/repository"
	"gin_restfull_api/internal/usecase"
	"gin_restfull_api/pkg/middleware"

	"github.com/gin-gonic/gin"
 // sesuaikan dengan nama module kamu
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
	db.AutoMigrate(&domain.User{},&domain.Customer{})

	r := gin.Default()

	repo := repository.NewUserRepository(db)
	uc := usecase.NewUserUseCase(repo)
	h := handler.NewUserHandler(uc)

	customerRepo := repository.NewCustomerRepository(db)
	customerUseCase := usecase.NewCustomerUseCase(customerRepo)
	customerHandler := handler.NewCustomerHandler(customerUseCase)



	api := r.Group("/api")
	{
		api.POST("/register", h.Register)
		api.POST("/login", h.Login)
		api.POST("/refresh", h.RefreshToken)
		api.GET("/profile", middleware.JWTProtected("admin"), h.Profile)

		api.GET("/customers", customerHandler.GetAll)
		api.POST("/customers", customerHandler.Create)
		api.GET("/customers/:id", customerHandler.GetByID)
		api.PUT("/customers/:id", customerHandler.Update)
		api.DELETE("/customers/:id", customerHandler.Delete)

		//r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))



	}

	r.Run(":3000")
}
