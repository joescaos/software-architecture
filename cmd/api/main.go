package main

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/gorm"

	"software-architecture/internal/adapters/api"
	"software-architecture/internal/adapters/repository"
	"software-architecture/internal/core/services"
)

func main() {
	// Load environment variables
	if err := godotenv.Load(); err != nil {
		log.Printf("No .env file found, using environment variables")
	}

	// Database connection
	dsn := buildDSN()
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	log.Println("Database connection established")

	// Migrate the schema
	db.AutoMigrate(&repository.ProductModel{})

	// Repository  initialization
	productRepo := repository.NewGormProductRepository(db)

	//Service initialization
	productService := services.NewProductService(productRepo)

	// Handler initialization
	productHandler := api.NewProductHandler(productService)

	// Router setup
	router := gin.Default()
	api.RegisterRoutes(router, productHandler)

	// Start server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Printf("Starting server on port %s", port)
	if err := router.Run(":" + port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}

}

func buildDSN() string {
	// Intentar usar DATABASE_URL completo primero (más flexible)
	if dsn := os.Getenv("DATABASE_URL"); dsn != "" {
		log.Println("Usando DATABASE_URL desde variables de entorno")
		return dsn
	}

	// build DSN from individual components
	host := getEnv("DB_HOST", "localhost")
	user := getEnv("DB_USER", "postgres")
	password := getEnv("DB_PASSWORD", "postgres")
	dbname := getEnv("DB_NAME", "products_db")
	port := getEnv("DB_PORT", "5432")
	sslmode := getEnv("DB_SSLMODE", "disable")

	// Formato DSN para PostgreSQL (gorm/postgres driver)
	// Más información: https://github.com/go-gorm/postgres
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		host, user, password, dbname, port, sslmode,
	)

	log.Printf("Construyendo DSN con host=%s, user=%s, dbname=%s", host, user, dbname)
	return dsn
}

// getEnv gets an environment variable or returns a default value if not set
func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
