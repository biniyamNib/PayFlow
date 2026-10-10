package main

import (
	"log"

	"github.com/gin-gonic/gin"

	"payflow/internal/database"
	"payflow/internal/handler"
	"payflow/internal/repository"
	"payflow/internal/service"
)

func main() {
	databaseURL := "postgres://postgres:Bini%40postgres1@localhost:5432/PayFlow?sslmode=disable"
	db, err := database.Connect(databaseURL)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	merchantRepository := repository.NewMerchantRepository(db)

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
