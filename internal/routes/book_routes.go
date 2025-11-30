package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/idoceb00/library-service/internal/handlers"
)

func RegisterBookRoutes(r *gin.Engine, handler *handlers.BookHandler) {
	books := r.Group("/books")
	{
		books.GET("/", handler.GetAll)
		books.GET("/:id", handler.GetByID)
		books.POST("/", handler.Create)
		books.PUT("/:id", handler.Update)
		books.DELETE("/:id", handler.Delete)
	}
}
