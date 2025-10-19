package main

import (
	"froztypoint/backend/internal/handler"
	"froztypoint/backend/internal/repository"
	"froztypoint/backend/internal/service"
	"log"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	dsn := "host=localhost user=reza password=Reza_wahyu123 dbname=froztypoint port=5432 sslmode=disable TimeZone=Asia/Jakarta"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("failed to connect database", err)
	}

	db.AutoMigrate(&repository.User{}, &repository.Product{})

	// pembuatan handler untuk menginisiasi handler dengan dependency service dan repository
	userRepo := repository.NewUserRepository(db)
	productRepo := repository.NewProductRepository(db)

	userService := service.NewUserService(userRepo)
	productService := service.NewProductService(productRepo)

	userHandler := handler.NewUserHandler(userService)
	productHandler := handler.NewProductHandler(productService)

	router := gin.Default()

	api := router.Group("/api")
	{
		userGroup := api.Group("/users")
		{
			userGroup.POST("/", userHandler.CreateUser)
			userGroup.GET("/:id", userHandler.GetUserByID)
		}

		productGroup := api.Group("/products")
		{
			productGroup.POST("/", productHandler.CreateProduct)
			productGroup.GET("/:id", productHandler.GetProductByID)
		}
	}
	log.Println("server running on:8080")

	router.Run(":8080")

}
