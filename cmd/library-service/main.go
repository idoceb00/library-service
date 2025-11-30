package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/idoceb00/library-service/internal/config"
	"github.com/idoceb00/library-service/internal/database"
)

func main() {
	cfg := config.LoadConfig()

	database.Connect(cfg)
	database.Migrate()

	router := gin.Default()

	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})

	log.Println("🚀 Server running on http://localhost:8080")
	router.Run(":8080")
}
