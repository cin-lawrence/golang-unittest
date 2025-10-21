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
