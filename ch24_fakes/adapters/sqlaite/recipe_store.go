package sqlaite

import (
	"context"
	"fmt"

	ent2 "ch24/fakes/adapters/sqlaite/ent"

	"ch24/fakes/domain/ingredient"
	"ch24/fakes/domain/recipe"
)

type RecipeStore struct {
	client *ent2.Client
}

func NewRecipeStore(client *ent2.Client) *RecipeStore {
	return &RecipeStore{
		client: client,
	}
}

func (r *RecipeStore) GetRecipes(ctx context.Context) (recipe.Recipes, error) {
	all, err := r.client.Recipe.Query().WithRecipeingredient().All(ctx)
	if err != nil {
		return nil, err
	}

	var recipes recipe.Recipes
	for _, r := range all {
		var ingrs ingredient.Ingredients
		for _, ingr := range r.Edges.Recipeingredient {
			first, err := ingr.QueryIngredient().First(ctx)
			if err != nil {
				return nil, err
			}

			ingrs = append(ingrs, ingredient.Ingredient{
				Name:     first.Name,
				Quantity: uint(ingr.Quantity),
			})
		}

		recipes = append(recipes, recipe.Recipe{
			Name:        r.Name,
			MealType:    recipe.MealType(r.MealType),
			Description: r.Description,
			Ingredients: ingrs,
		})
	}

	return recipes, nil
}

func (r *RecipeStore) AddRecipes(ctx context.Context, recipes ...recipe.Recipe) error {
	for _, newRecipe := range recipes {
		var recipeIngredients []*ent2.RecipeIngredient
		for _, newIngredient := range newRecipe.Ingredients {
			savedIngredient, err := CreateIngredientIfNotExists(ctx, r.client, newIngredient)
			if err != nil {
				return err
			}

			ri, err := r.client.RecipeIngredient.Create().AddIngredient(savedIngredient).SetQuantity(int(newIngredient.Quantity)).Save(ctx)
			if err != nil {
				return fmt.Errorf("could not create recipe ingredient: %v %w", newIngredient, err)
			}

			recipeIngredients = append(recipeIngredients, ri)
		}

		err := r.client.Recipe.Create().
			SetName(newRecipe.Name).
			AddRecipeingredient(recipeIngredients...).
			SetMealType(int(newRecipe.MealType)).
			SetDescription(newRecipe.Description).
			Exec(ctx)

		if err != nil {
			return err
		}
	}

	return nil
}
