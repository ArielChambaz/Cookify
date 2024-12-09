# Cookify

**Cookify** is a web application for managing cooking recipes. It allows users to add, edit, delete, and view recipes. The application is built using the Gin framework for Go.

## Contributor

This project has been done by Felly Florenza Willianto and Ariel Chambaz

---

## Installation

To set up Cookify on your local machine, follow these steps:

1. **Clone the repository:**

   ```bash
   git clone https://github.com/your-username/cookify.git
   cd cookify
   ```
2. **Install dependencies:**

   ```bash
   go mod download
   ```

3. **Run the application:**

   ```bash
   go run main.go
   ```

## Usage

### API Routes

- `GET /api/recipes`: Retrieve all recipes.
- `POST /api/recipes`: Add a new recipe.
- `POST /api/recipes/:id`: Update an existing recipe.
- `DELETE /api/recipes/:id`: Delete a recipe.
- `GET /recipes/search`: Search for recipes.

### HTML Routes

- `GET /`: Home page.
- `GET /recipes`: List of recipes.
- `GET /recipes/add`: Form to add a new recipe.
- `GET /recipes/edit/:id`: Form to edit an existing recipe.

## Models

### Recipe

The `Recipe` model is defined in [models/recipe.go](models/recipe.go):
```go

type Recipe struct {
    ID          int
    Name        string
    Category    string
    Ingredients []string
    Steps       []string
    ImageURL    string
}
```

## Handlers

Handlers are defined in [handlers/recipe.go](handlers/recipe.go):

- `GetRecipes`: Retrieve all recipes.
- `AddRecipe`: Add a new recipe.
- `EditRecipe`: Retrieve a recipe by ID for editing.
- `UpdateRecipe`: Update an existing recipe.
- `DeleteRecipe`: Delete a recipe.
- `SearchRecipes`: Search for recipes.

## Templates

HTML templates are located in the `templates` directory:

- `index.html`: Home page.
- `recipes.html`: List of recipes.
- `add_recipe.html`: Form to add a new recipe.
- `edit_recipe.html`: Form to edit an existing recipe.

## Static Files

Static files (CSS, images, etc.) are located in the `static` directory.

## Contributing

1. Fork the project.
2. Create your feature branch (`git checkout -b feature/AmazingFeature`).
3. Commit your changes (`git commit -m 'Add some AmazingFeature'`).
4. Push to the branch (`git push origin feature/AmazingFeature`).
5. Open a Pull Request.

Thank you for using Cookify! Bon appétit!