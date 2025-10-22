package sqlaite

import (
	"context"

	ent2 "ch24/fakes/adapters/sqlaite/ent"
	_ "github.com/mattn/go-sqlite3"

	entingr "ch24/fakes/adapters/sqlaite/ent/ingredient"
	"ch24/fakes/adapters/sqlaite/ent/pantry"
	"ch24/fakes/domain/ingredient"
)

type Pantry struct {
	client *ent2.Client
}

func NewPantry(client *ent2.Client) *Pantry {
	return &Pantry{
		client: client,
	}
}

func (p *Pantry) GetIngredients(ctx context.Context) (ingredient.Ingredients, error) {
	persistedPantries, err := p.client.Pantry.Query().WithIngredient().All(ctx)
	if err != nil {
		return nil, err
	}

	var allIngredients ingredient.Ingredients
	for _, pantryItem := range persistedPantries {
		if pantryItem.Quantity == 0 {
			continue
		}
		allIngredients = append(allIngredients, ingredient.Ingredient{
			Name:     pantryItem.Edges.Ingredient.Name,
			Quantity: uint(pantryItem.Quantity),
		})
	}
	return allIngredients, nil
}

func (p *Pantry) Store(ctx context.Context, ingredients ...ingredient.Ingredient) error {
	for _, newIngredient := range ingredients {
		if err := p.addOrIncrementIngredient(ctx, newIngredient); err != nil {
			return err
		}
	}
	return nil
}

func (p *Pantry) Remove(ctx context.Context, toRemove ...ingredient.Ingredient) error {
	for _, ingr := range toRemove {
		err := p.client.Pantry.Update().
			Where(pantry.HasIngredientWith(entingr.Name(ingr.Name))).
			AddQuantity(-int(ingr.Quantity)).
			Exec(ctx)

		if err != nil {
			return err
		}
	}
	return nil
}

func (p *Pantry) addOrIncrementIngredient(ctx context.Context, newIngredient ingredient.Ingredient) error {
	savedIngredient, err := CreateIngredientIfNotExists(ctx, p.client, newIngredient)
	if err != nil {
		return err
	}

	err = p.client.Pantry.Create().
		SetIngredientID(savedIngredient.ID).
		SetQuantity(int(newIngredient.Quantity)).
		Exec(ctx)

	if ent2.IsConstraintError(err) {
		err = p.client.Pantry.Update().
			Where(pantry.HasIngredientWith(entingr.ID(savedIngredient.ID))).
			AddQuantity(int(newIngredient.Quantity)).
			Exec(ctx)

		if err != nil {
			return err
		}
	}
	return nil
}
