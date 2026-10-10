package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	"payflow/internal/database"
	"payflow/internal/handler"
	"payflow/internal/repository"
	"payflow/internal/service"
)

func main() {
	err := godotenv.Load()
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	databaseURL := os.Getenv("DATABASE_URL")
	if databaseURL == "" {
		log.Fatal("DATABASE_URL environment variable is not set")
	}

	db, err := database.Connect(databaseURL)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()

	merchantRepository := repository.NewMerchantRepository(db)
	// merchantUserRepository := repository.NewMerchantUserRepository(db)

	merchantService := service.NewMerchantService(
		merchantRepository,
	)

	merchantHandler := handler.NewMerchantHandler(
		merchantService,
	)

	r := gin.Default()

	r.GET("/health", handler.Health)
	r.GET("/api/v1/merchants", merchantHandler.GetAll)
	r.POST("api/v1/merchants", merchantHandler.Create)
	r.GET("api/v1/merchants/:id", merchantHandler.GetByID)
	r.PUT("api/v1/merchants/:id", merchantHandler.Update)
	r.DELETE("api/v1/merchants/:id", merchantHandler.Delete)
	r.PATCH("/api/v1/merchants/:id/activate", merchantHandler.Activate)

	log.Println("PayFlow API running on :8080")

	if err := r.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
