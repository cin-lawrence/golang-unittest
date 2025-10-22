package planner

import (
	"context"
	"fmt"
	"time"

	"ch24/fakes/domain/ingredient"
	"ch24/fakes/domain/recipe"
)

type Planner struct {
	recipeBook RecipeBook
	pantry     Pantry
}

func NewPlanner(recipes RecipeBook, ingredientStore Pantry) *Planner {
	return &Planner{recipeBook: recipes, pantry: ingredientStore}
}

func (p *Planner) ScheduleMeal(ctx context.Context, r recipe.Recipe, _ time.Time) error {
	availableIngredients, err := p.pantry.GetIngredients(ctx)
	if err != nil {
		return err
	}

	if hasIngredients, missing := haveIngredients(availableIngredients, r); !hasIngredients {
		return ErrorMissingIngredients{
			MissingIngredients: missing,
		}
	}

	return p.pantry.Remove(ctx, r.Ingredients...)
}

func (p *Planner) SuggestRecipes(ctx context.Context) (recipe.Recipes, error) {
	availableIngredients, err := p.pantry.GetIngredients(ctx)
	if err != nil {
		return nil, err
	}

	recipes, err := p.recipeBook.GetRecipes(ctx)
	if err != nil {
		return nil, err
	}

	var suggestions recipe.Recipes
	for _, r := range recipes {
		if hasIngredients, _ := haveIngredients(availableIngredients, r); hasIngredients {
			suggestions = append(suggestions, r)
		}
	}
	return suggestions, nil
}

func haveIngredients(
	availableIngredients ingredient.Ingredients,
	recipe recipe.Recipe,
) (hasIngredients bool, missing ingredient.Ingredients) {
	for _, ingr := range recipe.Ingredients {
		if !availableIngredients.Has(ingr) {
			missing = append(missing, ingr)
		}
	}

	if len(missing) > 0 {
		return false, missing
	}
	return true, nil
}

type ErrorMissingIngredients struct {
	MissingIngredients ingredient.Ingredients
}

func (e ErrorMissingIngredients) Error() string {
	return fmt.Sprintf("missing ingredientts: %v", e.MissingIngredients)
}
