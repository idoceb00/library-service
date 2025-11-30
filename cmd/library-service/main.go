package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/idoceb00/library-service/internal/config"
	"github.com/idoceb00/library-service/internal/database"
	"github.com/idoceb00/library-service/internal/handlers"
	"github.com/idoceb00/library-service/internal/repositories"
	"github.com/idoceb00/library-service/internal/routes"
	"github.com/idoceb00/library-service/internal/services"
)

func main() {
	cfg := config.LoadConfig()

	database.Connect(cfg)
	database.Migrate()

	bookRepo := repositories.NewBookRepository(database.DB)
	bookService := services.NewBookService(bookRepo)
	bookHandler := handlers.NewBookHandler(bookService)

	router := gin.Default()

	routes.RegisterBookRoutes(router, bookHandler)

	log.Println("🚀 Server running on http://localhost:8080")
	router.Run(":8080")
}
