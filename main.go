package main

import (
	"cookify/handlers"
	"cookify/models"
	"html/template"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	funcMap := template.FuncMap{
		"join": strings.Join,
	}

	r.SetFuncMap(funcMap)
	r.LoadHTMLGlob("templates/*")

	// Serve static files
	r.Static("/static", "./static")

	// Routes for the API
	r.GET("/api/recipes", handlers.GetRecipes)
	r.POST("/api/recipes", handlers.AddRecipe)
	r.POST("/api/recipes/:id", handlers.UpdateRecipe)
	r.DELETE("/api/recipes/:id", handlers.DeleteRecipe)

	// Routes for HTML
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})

	r.GET("/recipes", func(c *gin.Context) {
		recipes, err := models.GetAllRecipes()
		if err != nil {
			c.String(http.StatusInternalServerError, "Error fetching recipes")
			return
		}
		c.HTML(http.StatusOK, "recipes.html", recipes)
	})

	r.GET("/recipes/add", func(c *gin.Context) {
		c.HTML(http.StatusOK, "add_recipe.html", nil)
	})

	r.GET("/recipes/edit/:id", handlers.EditRecipe)

	r.GET("/recipes/search", handlers.SearchRecipes) // New route for searching recipes

	// Start the server
	r.Run(":8080")
}
