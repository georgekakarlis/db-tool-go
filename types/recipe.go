package types

import "time"

type Recipe struct {
	ID          int       `db:"id"`
	Name        string    `db:"name"`
	Content     string    `db:"content"`
	CreatedAt   time.Time `db:"created_at"`
	UpdatedAt   time.Time `db:"updated_at"`
	CookingTime int       `db:"cooking_time"`
	Servings    int       `db:"servings"`
	Difficulty  string    `db:"difficulty"`
	TotalTime   int       `db:"total_time"`
	IngredientID int      `db:"ingredient_id"`
}

type Ingredient struct {
	ID          int       `db:"id"`
	Name        string    `db:"name"`
	CreatedAt   time.Time `db:"created_at"`
	UpdatedAt   time.Time `db:"updated_at"`
	Description string    `db:"description"`
}

type Review struct {
	ID        int       `db:"id"`
	RecipeID  int       `db:"recipe_id"`
	Rating    int       `db:"rating"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}