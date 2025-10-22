package planner_test

import (
	"context"
	"errors"
	"testing"
	"time"

	"ch24/fakes/adapters/inmemory"
	"ch24/fakes/domain/ingredient"
	"ch24/fakes/domain/planner"
	"ch24/fakes/internal/expect"
	"ch24/fakes/internal/plannertest"
)

func TestRecipePlanner(t *testing.T) {
	t.Run("with in-memory store", func(t *testing.T) {
		(&RecipePlannerTest{
			CreateDependencies: func() (planner.RecipeBook, planner.Pantry, Cleanup) {
				return inmemory.NewRecipeStore(), inmemory.NewPantry(), func() {}
			},
		}).Test(t)
	})
}

type Cleanup func()

type RecipePlannerTest struct {
	CreateDependencies func() (planner.RecipeBook, planner.Pantry, Cleanup)
}

func (r *RecipePlannerTest) Test(t *testing.T) {
	t.Run("planning meals", func(t *testing.T) {
		t.Run("happy path, have ingredients for a recipe, schedule it, update pantry", func(t *testing.T) {
			var (
				ctx                          = context.Background()
				lasagna                      = plannertest.RandomRecipe()
				recipeBook, pantry, teardown = r.CreateDependencies()
				sut                          = planner.NewPlanner(recipeBook, pantry)
			)
			t.Cleanup(teardown)

			expect.NoErr(t, recipeBook.AddRecipes(ctx, lasagna))
			expect.NoErr(t, pantry.Store(ctx, lasagna.Ingredients...))

			expect.NoErr(t, sut.ScheduleMeal(ctx, lasagna, time.Now()))
			remainingIngredients, err := pantry.GetIngredients(ctx)
			expect.NoErr(t, err)
			expect.Equal(t, 0, len(remainingIngredients))
		})

		t.Run("returns a missing ingredients error if you try to schedule a meal without all the ingredients", func(t *testing.T) {
			var (
				ctx                         = context.Background()
				lasagna                     = plannertest.RandomRecipe()
				recipeBook, store, teardown = r.CreateDependencies()
				sut                         = planner.NewPlanner(recipeBook, store)
			)
			t.Cleanup(teardown)

			expect.NoErr(t, recipeBook.AddRecipes(ctx, lasagna))

			err := sut.ScheduleMeal(ctx, lasagna, time.Now())
			expect.Err(t, err)

			missingIngredientsErr, ok := err.(planner.ErrorMissingIngredients)
			expect.True(t, ok)
			expect.DeepEqual(t, planner.ErrorMissingIngredients{
				MissingIngredients: lasagna.Ingredients,
			}, missingIngredientsErr)
		})

		t.Run("when recipeBook fails to get ingredients, we get an error", func(t *testing.T) {
			var (
				ctx                          = context.Background()
				lasagna                      = plannertest.RandomRecipe()
				recipeBook, pantry, teardown = r.CreateDependencies()
				failingPantry                = planner.NewPantryDelegate(pantry)
			)
			t.Cleanup(teardown)

			failingPantry.GetIngredientsFunc = func(ctx context.Context) (ingredient.Ingredients, error) {
				return nil, errors.New("oh no")
			}

			sut := planner.NewPlanner(recipeBook, failingPantry)
			expect.NoErr(t, recipeBook.AddRecipes(ctx, lasagna))
			expect.NoErr(t, pantry.Store(ctx, lasagna.Ingredients...))

			err := sut.ScheduleMeal(ctx, lasagna, time.Now())
			expect.Err(t, err)
		})

		t.Run("returns the specific ingredients missing if you try to schedule a meal with some missing ingredients", func(t *testing.T) {
			var (
				ctx                          = context.Background()
				lasagna                      = plannertest.RandomRecipe()
				recipeBook, pantry, teardown = r.CreateDependencies()
				sut                          = planner.NewPlanner(recipeBook, pantry)
			)
			t.Cleanup(teardown)

			missingIngredient, ingredientsWeHave := lasagna.Ingredients[0], lasagna.Ingredients[1:]
			expect.NoErr(t, pantry.Store(ctx, ingredientsWeHave...))

			err := sut.ScheduleMeal(ctx, lasagna, time.Now())
			expect.Err(t, err)

			missingIngredientsErr, ok := err.(planner.ErrorMissingIngredients)
			expect.True(t, ok)
			expect.DeepEqual(t, planner.ErrorMissingIngredients{
				MissingIngredients: ingredient.Ingredients{missingIngredient},
			}, missingIngredientsErr)
		})
	})

	t.Run("suggesting recipes", func(t *testing.T) {
		t.Run("if don't have the ingredients for a meal, we can't make it", func(t *testing.T) {
			var (
				ctx                          = context.Background()
				pie                          = plannertest.RandomRecipe()
				recipeBook, pantry, teardown = r.CreateDependencies()
				sut                          = planner.NewPlanner(recipeBook, pantry)
			)
			t.Cleanup(teardown)

			expect.NoErr(t, recipeBook.AddRecipes(ctx, pie))

			recipes, err := sut.SuggestRecipes(ctx)
			expect.NoErr(t, err)
			plannertest.AssertDoesntHaveRecipe(t, recipes, pie)
		})

		t.Run("if we have the ingredients for a recipe we can make it", func(t *testing.T) {
			var (
				ctx                           = context.Background()
				bananaBread                   = plannertest.RandomRecipe()
				recipeInsufficientIngredients = plannertest.RandomRecipe()
				recipeBook, pantry, teardown  = r.CreateDependencies()
				sut                           = planner.NewPlanner(recipeBook, pantry)
			)
			t.Cleanup(teardown)

			expect.NoErr(t, recipeBook.AddRecipes(ctx, bananaBread, recipeInsufficientIngredients))
			expect.NoErr(t, pantry.Store(ctx, bananaBread.Ingredients...))

			recipes, err := sut.SuggestRecipes(ctx)
			expect.NoErr(t, err)
			plannertest.AssertHasRecipe(t, recipes, bananaBread)
			plannertest.AssertDoesntHaveRecipe(t, recipes, recipeInsufficientIngredients)
		})

		t.Run("if we have ingredients for 2 recipes, we can make both", func(t *testing.T) {
			var (
				ctx                         = context.Background()
				bananaBread                 = plannertest.RandomRecipe()
				bananaMilkshake             = plannertest.RandomRecipe()
				recipeBook, store, teardown = r.CreateDependencies()
				sut                         = planner.NewPlanner(recipeBook, store)
			)
			t.Cleanup(teardown)

			expect.NoErr(t, recipeBook.AddRecipes(ctx, bananaBread, bananaMilkshake))
			expect.NoErr(t, store.Store(ctx, bananaBread.Ingredients...))
			expect.NoErr(t, store.Store(ctx, bananaMilkshake.Ingredients...))

			recipes, err := sut.SuggestRecipes(ctx)
			expect.NoErr(t, err)
			plannertest.AssertHasRecipe(t, recipes, bananaBread)
			plannertest.AssertHasRecipe(t, recipes, bananaMilkshake)
		})
	})
}
