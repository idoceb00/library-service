package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/idoceb00/library-service/internal/models"
	"github.com/idoceb00/library-service/internal/services"
)

type BookHandler struct {
	service *services.BookService
}

func NewBookHandler(service *services.BookService) *BookHandler {
	return &BookHandler{service: service}
}

func (h *BookHandler) GetAll(c *gin.Context) {
	books, err := h.service.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, books)
}

func (h *BookHandler) GetByID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	book, err := h.service.GetById(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "book not found"})
		return
	}

	c.JSON(http.StatusOK, book)
}

func (h *BookHandler) Create(c *gin.Context) {
	var book models.Book

	if err := c.ShouldBindJSON(&book); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if err := h.service.Create(&book); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, book)
}

func (h *BookHandler) Update(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	var data models.Book
	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	book, err := h.service.Update(uint(id), &data)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "book not found"})
		return
	}

	c.JSON(http.StatusOK, book)
}

func (h *BookHandler) Delete(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	if err := h.service.Delete(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusNoContent)
}
