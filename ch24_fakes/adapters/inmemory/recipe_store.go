package inmemory

import (
	"context"

	"ch24/fakes/domain/recipe"
)

type RecipeStore struct {
	Recipes recipe.Recipes
}

func NewRecipeStore() *RecipeStore {
	return &RecipeStore{}
}

func (s *RecipeStore) GetRecipes(_ context.Context) (recipe.Recipes, error) {
	return s.Recipes, nil
}

func (s *RecipeStore) AddRecipes(_ context.Context, r ...recipe.Recipe) error {
	s.Recipes = append(s.Recipes, r...)
	return nil
}
