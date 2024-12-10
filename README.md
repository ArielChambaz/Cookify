# Cookify

**Cookify** is a web application for managing cooking recipes. It allows users to add, edit, delete, and view recipes. The application is built using the Gin framework for Go.

## Contributor

This project has been done by Felllay Fiorenza Willianto (5025221110) and Ariel Chambaz ()

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
  
  ![image](https://github.com/user-attachments/assets/ad4dedc2-6de6-4626-baec-51d155de3188)

- `recipes.html`: List of recipes
  
  ![image](https://github.com/user-attachments/assets/888f52b5-0bce-4d00-85e5-89487ad339ae)


```html
 ...

   body>
       <!-- Navbar -->
       <header class="navbar">
           <a href="/" class="navbar-brand">Cookify</a>
           <div class="navbar-nav2">
               <a href="/"> <b>Dashboard</b> </a>
               <a href="/recipes"> <b>Recipes</b> </a>
               <a href="/recipes/add"> <b>Add Recipe</b> </a>
           </div>
       </header>
   
       <div class="container">
           <div class="title">
               <h1 class="mt-5"><b> Discover your Recipes! </b></h1>
           </div>
           <div class="d-flex justify-content-end mb-3">
               <form action="/recipes/search" method="GET" class="form-inline">
                   <input type="text" name="q" class="form-control mr-2" placeholder="Search recipes...">
                   <button type="submit" class="btn btn-primary">Search</button>
               </form>
           </div>

   ...
      

```

- `add_recipe.html`: Form to add a new recipe.
  
  ![image](https://github.com/user-attachments/assets/0e584dab-3d12-4371-9110-e63e11e22e6c)

  ```html

  ...
      <body>
       <!-- Navbar -->
       <header class="navbar">
           <a href="/" class="navbar-brand">Cookify</a>
           <div class="navbar-nav2">
               <a href="/"> <b>Dashboard</b> </a>
               <a href="/recipes"> <b>Recipes</b> </a>
               <a href="/recipes/add"> <b>Add Recipe</b> </a>
           </div>
       </header>
   
       <div class="container">
           <h1 class="mb-4">Tell us your secret Recipe!</h1>
           <form action="/api/recipes" method="POST">
               <div class="form-group">
                   <label for="name">Recipe Name :</label>
                   <input type="text" id="name" name="name" class="form-control" required>
               </div>
   
               <div class="form-group">
                   <label for="category">Category :</label>
                   <input type="text" id="category" name="category" class="form-control">
               </div>
   
               <div class="form-group">
                   <label for="ingredients">Ingredients (comma-separated) :</label>
                   <textarea id="ingredients" name="ingredients" rows="4" class="form-control" required></textarea>
               </div>
   
               <div class="form-group">
                   <label for="steps">Steps (comma-separated) :</label>
                   <textarea id="steps" name="steps" rows="4" class="form-control" required></textarea>
               </div>
   
               <div class="form-group">
                   <label for="image_url">Image URL :</label>
                   <input type="url" id="image_url" name="image_url" class="form-control">
               </div>
   
               <button type="submit" class="btn btn-primary">Add Recipe</button>
           </form>
    </div>

  ...

  ```html

- `edit_recipe.html`: Form to edit an existing recipe.
  
  ![image](https://github.com/user-attachments/assets/ef847aba-4993-4cea-9537-620ad424d367)

  ```html
  ...
      <body>
       <!-- Navbar -->
       <header class="navbar">
           <a href="/" class="navbar-brand">Cookify</a>
           <div class="navbar-nav2">
               <a href="/"> <b>Dashboard</b> </a>
               <a href="/recipes"> <b>Recipes</b> </a>
               <a href="/recipes/add"> <b>Add Recipe</b> </a>
           </div>
       </header>
   
       <div class="container">
           <h1 class="mb-4">Edit Your Recipe</h1>
           <!-- Edit Recipe Form -->
           <form action="/api/recipes/{{.ID}}" method="POST">
               <input type="hidden" name="_method" value="PUT"> <!-- Simulate PUT request -->
               <div class="form-group">
                   <label for="name">Recipe Name :</label>
                   <input type="text" id="name" name="name" class="form-control" value="{{.Name}}" required>
               </div>
   
               <div class="form-group">
                   <label for="category">Category :</label>
                   <input type="text" id="category" name="category" class="form-control" value="{{.Category}}">
               </div>
  ...
  ```html

- - `ViewRecipes` : View full recipes.

## Static Files

Static files (CSS, images, etc.) are located in the `static` directory.

## Contributing

1. Fork the project.
2. Create your feature branch (`git checkout -b feature/AmazingFeature`).
3. Commit your changes (`git commit -m 'Add some AmazingFeature'`).
4. Push to the branch (`git push origin feature/AmazingFeature`).
5. Open a Pull Request.

Thank you for using Cookify! Bon appétit!
