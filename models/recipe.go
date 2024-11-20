package models

import (
	"database/sql"
	"log"
	"strings"

	_ "github.com/mattn/go-sqlite3"
)

type Recipe struct {
	ID          int      `json:"id"`
	Name        string   `json:"name"`
	Ingredients []string `json:"ingredients"`
	Steps       []string `json:"steps"`
	ImageURL    string   `json:"image_url,omitempty"`
	Category    string   `json:"category,omitempty"`
}

var db *sql.DB

// init initializes the database and creates the recipes table if it doesn't exist
func init() {
	var err error
	db, err = sql.Open("sqlite3", "./cookify.db")
	if err != nil {
		log.Fatal(err)
	}

	createTableSQL := `
    CREATE TABLE IF NOT EXISTS recipes (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        name TEXT,
        ingredients TEXT,
        steps TEXT,
        image_url TEXT,
        category TEXT
    );`
	_, err = db.Exec(createTableSQL)
	if err != nil {
		log.Fatal(err)
	}
}

func GetAllRecipes() ([]Recipe, error) {
	rows, err := db.Query("SELECT id, name, ingredients, steps, image_url, category FROM recipes")
	if err != nil {
		log.Printf("Error querying recipes: %v", err)
		return nil, err
	}
	defer rows.Close()

	var recipes []Recipe
	for rows.Next() {
		var r Recipe
		var ingredients, steps string
		err := rows.Scan(&r.ID, &r.Name, &ingredients, &steps, &r.ImageURL, &r.Category)
		if err != nil {
			log.Printf("Error scanning recipe: %v", err)
			return nil, err
		}
		r.Ingredients = strings.Split(ingredients, ",")
		r.Steps = strings.Split(steps, ",")
		recipes = append(recipes, r)
	}

	if err = rows.Err(); err != nil {
		log.Printf("Error iterating over recipes: %v", err)
		return nil, err
	}

	return recipes, nil
}

func GetRecipeByID(id string) (Recipe, error) {
	var r Recipe
	var ingredients, steps string
	err := db.QueryRow("SELECT id, name, ingredients, steps, image_url, category FROM recipes WHERE id = ?", id).Scan(&r.ID, &r.Name, &ingredients, &steps, &r.ImageURL, &r.Category)
	if err != nil {
		return r, err
	}
	r.Ingredients = strings.Split(ingredients, ",")
	r.Steps = strings.Split(steps, ",")
	return r, nil
}

func AddRecipe(r Recipe) error {
	ingredients := strings.Join(r.Ingredients, ",")
	steps := strings.Join(r.Steps, ",")

	statement, err := db.Prepare("INSERT INTO recipes (name, ingredients, steps, image_url, category) VALUES (?, ?, ?, ?, ?)")
	if err != nil {
		return err
	}
	_, err = statement.Exec(r.Name, ingredients, steps, r.ImageURL, r.Category)
	return err
}

func UpdateRecipe(r Recipe) error {
	ingredients := strings.Join(r.Ingredients, ",")
	steps := strings.Join(r.Steps, ",")

	statement, err := db.Prepare("UPDATE recipes SET name = ?, ingredients = ?, steps = ?, image_url = ?, category = ? WHERE id = ?")
	if err != nil {
		return err
	}
	_, err = statement.Exec(r.Name, ingredients, steps, r.ImageURL, r.Category, r.ID)
	return err
}

func DeleteRecipe(id int) error {
	statement, err := db.Prepare("DELETE FROM recipes WHERE id = ?")
	if err != nil {
		return err
	}
	_, err = statement.Exec(id)
	return err
}
