package routes

import (
	"noteapp/handlers"
	"noteapp/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	r.POST("/register", handlers.Register)
	r.POST("/login", handlers.Login)

	auth := r.Group("/")
	auth.Use(middleware.AuthMiddleware())
	{
		auth.GET("/notes", handlers.GetNotes)
		auth.GET("/notes/:id", handlers.GetNoteByID)
		auth.POST("/notes", handlers.CreateNote)
		auth.PUT("/notes/:id", handlers.UpdateNote)
		auth.DELETE("/notes/:id", handlers.DeleteNote)

		auth.GET("/categories", handlers.GetCategories)
		auth.GET("/categories/:id", handlers.GetCategoryByID)
		auth.POST("/categories", handlers.CreateCategory)
		auth.PUT("/categories/:id", handlers.UpdateCategory)
		auth.DELETE("/categories/:id", handlers.DeleteCategory)
	}
}
