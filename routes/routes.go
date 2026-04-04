package routes

import (
	"noteapp/handlers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	r.GET("/notes", handlers.GetNotes)
	r.GET("/notes/:id", handlers.GetNoteByID)
	r.POST("/notes", handlers.CreateNote)
	r.PUT("/notes/:id", handlers.UpdateNote)
	r.DELETE("/notes/:id", handlers.DeleteNote)

	r.GET("/categories", handlers.GetCategories)
	r.GET("/categories/:id", handlers.GetCategoryByID)
	r.POST("/categories", handlers.CreateCategory)
	r.PUT("/categories/:id", handlers.UpdateCategory)
	r.DELETE("/categories/:id", handlers.DeleteCategory)
}
