package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"virtual-library/routes"

	"github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"virtual-library/config"
	"virtual-library/models"
)

func main() {
	logger := logrus.New()
	logger.SetFormatter(&logrus.JSONFormatter{})
	logger.SetOutput(os.Stdout)

	cfg, err := config.LoadConfig()
	if err != nil {
		logger.Fatalf("Error...")
	}

	// Connect to the database
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=UTC",
		cfg.DB.Host, cfg.DB.User, cfg.DB.Password, cfg.DB.Name, cfg.DB.Port)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		logger.Fatalf("Failed to connect to database: %v", err)
	}

	err = db.AutoMigrate(&models.Book{})
	if err != nil {
		logger.Fatalf("Failed to migrate database: %v", err)
	}

	router := routes.SetupRouter(db)
	err = router.Run(":8080")
	if err != nil {
		log.Fatal(err)
	}

	// Start the server
	port := cfg.Server.Port
	logger.Infof("Server is running on port %s", port)
	if err := router.Run(":" + port); err != nil && err != http.ErrServerClosed {
		logger.Fatalf("Failed to start server: %v", err)
	}
}
