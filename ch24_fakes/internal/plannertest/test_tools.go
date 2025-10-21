package plannertest

import (
	"math/rand"

	"github.com/google/uuid"

	"ch24/fakes/domain/ingredient"
	"ch24/fakes/domain/recipe"
)

func RandomRecipes() recipe.Recipes {
	return recipe.Recipes{
		RandomRecipe(),
		RandomRecipe(),
		RandomRecipe(),
	}
}

func RandomRecipe() recipe.Recipe {
	return recipe.Recipe{
		Name:        uuid.New().String(),
		Ingredients: RandomIngredients(),
	}
}

func RandomIngredients() ingredient.Ingredients {
	return ingredient.Ingredients{
		RandomIngredient(),
		RandomIngredient(),
		RandomIngredient(),
	}
}

func RandomIngredient() ingredient.Ingredient {
	return ingredient.Ingredient{
		Name:     uuid.New().String(),
		Quantity: uint(rand.Intn(10)) + 1,
	}
}
