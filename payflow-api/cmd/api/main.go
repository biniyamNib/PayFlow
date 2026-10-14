package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	"payflow/internal/database"
	"payflow/internal/handler"
	"payflow/internal/middleware"
	"payflow/internal/repository"
	"payflow/internal/service"
)

func main() {
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

	merchantService := service.NewMerchantService(merchantRepository,)
	merchantUserervice := service.NewMerchantUserService(db)

	merchantHandler := handler.NewMerchantHandler(merchantService,)
	merchantUserHandler := handler.NewMerchantUserHandler(merchantUserervice,)

	r := gin.Default()

	r.GET("/health", handler.Health)
	r.GET("/api/v1/merchants", merchantHandler.GetAll)
	r.POST("/api/v1/merchants", merchantHandler.Create)
	r.GET("/api/v1/merchants/:id", middleware.AuthMiddleware(), merchantHandler.GetByID)
	r.PUT("/api/v1/merchants/:id", middleware.AuthMiddleware(), middleware.RequireRole("OWNER"), merchantHandler.Update)
	r.PATCH("/api/v1/merchants/:id/activate", middleware.AuthMiddleware(), middleware.RequireRole("OWNER"), merchantHandler.Activate)
	r.POST("/api/v1/merchant-users/register", merchantUserHandler.Register)
	r.POST("/api/v1/merchant-users/login", merchantUserHandler.Login)

	r.GET("/api/v1/protected", middleware.AuthMiddleware(), func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H {
			"message": "You are authenticated",
		})
	})
	r.GET("/api/v1/me", middleware.AuthMiddleware(), merchantHandler.Me)
	r.DELETE("/api/v1/merchants/:id", middleware.AuthMiddleware(), middleware.RequireRole("OWNER"), merchantHandler.Delete)
	r.GET("/api/v1/owner-only", middleware.AuthMiddleware(), middleware.RequireRole("OWNER"),
		func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "Welcome, OWNER",
			})
		},
	)

	log.Println("PayFlow API running on :8080")

	if err := r.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
