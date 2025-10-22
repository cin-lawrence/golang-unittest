package sqlaite

import (
	"context"
	"fmt"

	ent2 "ch24/fakes/adapters/sqlaite/ent"
	"ch24/fakes/domain/ingredient"
)

func CreateIngredientIfNotExists(ctx context.Context, client *ent2.Client, newIngredient ingredient.Ingredient) (*ent2.Ingredient, error) {
	id, err := client.Ingredient.Create().SetName(newIngredient.Name).OnConflict().Ignore().ID(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to create ingredient %v: %w", newIngredient, err)
	}

	return client.Ingredient.GetX(ctx, id), nil
}
