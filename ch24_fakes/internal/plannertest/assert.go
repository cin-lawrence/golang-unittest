package plannertest

import (
	"testing"

	"ch24/fakes/domain/recipe"
	"ch24/fakes/internal/expect"
)

func AssertHasRecipe(t testing.TB, recipes recipe.Recipes, expected recipe.Recipe) {
	t.Helper()

	_, found := recipes.FindByName(expected.Name)
	expect.True(t, found)
}

func AssertDoesntHaveRecipe(t testing.TB, recipes recipe.Recipes, expected recipe.Recipe) {
	t.Helper()

	_, found := recipes.FindByName(expected.Name)
	expect.False(t, found)
}
