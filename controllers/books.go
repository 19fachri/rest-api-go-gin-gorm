package controllers

import (
	"example/web-service-go/models"
	"example/web-service-go/validators"
	"net/http"

	"github.com/gin-gonic/gin"
)

func FindAllBooks(c *gin.Context) {
	var books []models.Book
	models.DB.Find(&books)
	c.JSON(http.StatusOK, gin.H{"books": books})
}

func CreateBook(c *gin.Context) {
	var input validators.CreateBookInput
	err := c.ShouldBindJSON(&input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	book := models.Book{Title: input.Title, Author: input.Author}
	models.DB.Create(&book)

	c.JSON(http.StatusCreated, gin.H{"book": book})
}

func FindBookById(c *gin.Context) {
	var book models.Book
	id := c.Param("id")

	err := models.DB.Where("id = ?", id).First(&book).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Data not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"book": book})
}

func UpdateBookById(c *gin.Context) {
	var book models.Book
	id := c.Param("id")
	err := models.DB.Where("id = ?", id).First(&book).Error
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Data not found"})
		return
	}

	var input validators.UpdateBookInput
	err = c.ShouldBindJSON(&input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	models.DB.Model(&book).Updates(models.Book{Title: input.Title, Author: input.Author})
	c.JSON(http.StatusOK, gin.H{"data": book})
}

func DeleteBook(c *gin.Context) {
	var book models.Book
	id := c.Param("id")
	err := models.DB.Where("id = ?", id).First(&book).Error
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Data not found"})
		return
	}
	models.DB.Delete(&book)
	c.JSON(http.StatusOK, gin.H{"message": "delete success"})
}
