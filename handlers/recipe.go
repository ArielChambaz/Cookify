package handlers

import (
	"cookify/models"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

// Handler to retrieve all recipes
func GetRecipes(c *gin.Context) {
	recipes, err := models.GetAllRecipes()
	if err != nil {
		log.Printf("Error fetching recipes: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, recipes)
}

// Handler to add a new recipe
func AddRecipe(c *gin.Context) {
	// Parse form data
	name := c.PostForm("name")
	category := c.PostForm("category")
	ingredients := strings.Split(c.PostForm("ingredients"), ",")
	steps := strings.Split(c.PostForm("steps"), ",")
	imageURL := c.PostForm("image_url")

	// Log the received data
	log.Printf("Received data: name=%s, category=%s, ingredients=%v, steps=%v, imageURL=%s", name, category, ingredients, steps, imageURL)

	// Create a new recipe
	newRecipe := models.Recipe{
		Name:        name,
		Category:    category,
		Ingredients: ingredients,
		Steps:       steps,
		ImageURL:    imageURL,
	}

	// Add the recipe to the database
	err := models.AddRecipe(newRecipe)
	if err != nil {
		log.Printf("Failed to add recipe: %v", err)
		c.String(http.StatusInternalServerError, "Failed to add recipe: %v", err)
		return
	}

	// Redirect to the recipes list
	c.Redirect(http.StatusSeeOther, "/recipes")
}

// Handler to edit a recipe
func EditRecipe(c *gin.Context) {
	id := c.Param("id")
	recipe, err := models.GetRecipeByID(id)
	if err != nil {
		c.String(http.StatusInternalServerError, "Error fetching recipe: %v", err)
		return
	}
	c.HTML(http.StatusOK, "edit_recipe.html", recipe)
}

// Handler to update a recipe
func UpdateRecipe(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.String(http.StatusBadRequest, "Invalid recipe ID")
		return
	}

	name := c.PostForm("name")
	category := c.PostForm("category")
	ingredients := strings.Split(c.PostForm("ingredients"), ",")
	steps := strings.Split(c.PostForm("steps"), ",")
	imageURL := c.PostForm("image_url")

	recipe := models.Recipe{
		ID:          id,
		Name:        name,
		Category:    category,
		Ingredients: ingredients,
		Steps:       steps,
		ImageURL:    imageURL,
	}

	err = models.UpdateRecipe(recipe)
	if err != nil {
		c.String(http.StatusInternalServerError, "Failed to update recipe: %v", err)
		return
	}

	c.Redirect(http.StatusSeeOther, "/recipes")
}

// Handler to delete a recipe
func DeleteRecipe(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.String(http.StatusBadRequest, "Invalid recipe ID")
		return
	}

	err = models.DeleteRecipe(id)
	if err != nil {
		c.String(http.StatusInternalServerError, "Failed to delete recipe: %v", err)
		return
	}

	c.Status(http.StatusNoContent)
}

// Handler to search recipes
func SearchRecipes(c *gin.Context) {
	query := c.Query("q")
	recipes, err := models.SearchRecipes(query)
	if err != nil {
		c.String(http.StatusInternalServerError, "Error searching recipes: %v", err)
		return
	}
	c.HTML(http.StatusOK, "recipes.html", recipes)
}
