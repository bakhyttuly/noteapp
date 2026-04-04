package handlers

import (
	"net/http"
	"noteapp/config"
	"noteapp/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetNotes(c *gin.Context) {
	var notes []models.Note

	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "5"))
	categoryID := c.Query("category_id")

	if page < 1 {
		page = 1
	}

	if limit < 1 {
		limit = 5
	}

	offset := (page - 1) * limit

	query := config.DB.Preload("Category")

	if categoryID != "" {
		query = query.Where("category_id = ?", categoryID)
	}

	query.Offset(offset).Limit(limit).Find(&notes)

	c.JSON(http.StatusOK, notes)
}

func GetNoteByID(c *gin.Context) {
	id := c.Param("id")

	var note models.Note
	if err := config.DB.Preload("Category").First(&note, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Note not found"})
		return
	}

	c.JSON(http.StatusOK, note)
}

func CreateNote(c *gin.Context) {
	var note models.Note

	if err := c.ShouldBindJSON(&note); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
		return
	}

	if note.Title == "" || note.Content == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Title and content are required"})
		return
	}

	if note.CategoryID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Category ID is required"})
		return
	}

	var category models.Category
	if err := config.DB.First(&category, note.CategoryID).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Category not found"})
		return
	}

	config.DB.Create(&note)
	config.DB.Preload("Category").First(&note, note.ID)

	c.JSON(http.StatusCreated, note)
}

func UpdateNote(c *gin.Context) {
	id := c.Param("id")

	var note models.Note
	if err := config.DB.First(&note, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Note not found"})
		return
	}

	var input models.Note
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
		return
	}

	if input.Title == "" || input.Content == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Title and content are required"})
		return
	}

	if input.CategoryID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Category ID is required"})
		return
	}

	var category models.Category
	if err := config.DB.First(&category, input.CategoryID).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Category not found"})
		return
	}

	note.Title = input.Title
	note.Content = input.Content
	note.CategoryID = input.CategoryID

	config.DB.Save(&note)
	config.DB.Preload("Category").First(&note, note.ID)

	c.JSON(http.StatusOK, note)
}

func DeleteNote(c *gin.Context) {
	id := c.Param("id")

	var note models.Note
	if err := config.DB.First(&note, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Note not found"})
		return
	}

	config.DB.Delete(&note)
	c.JSON(http.StatusOK, gin.H{"message": "Note deleted successfully"})
}
