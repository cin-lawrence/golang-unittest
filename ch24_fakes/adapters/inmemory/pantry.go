package inmemory

import (
	"context"

	"ch24/fakes/domain/ingredient"
)

type Pantry struct {
	ingredients ingredient.Ingredients
}

func NewPantry() *Pantry {
	return &Pantry{}
}

func (p *Pantry) Remove(_ context.Context, i ...ingredient.Ingredient) error {
	for _, ingr := range i {
		p.ingredients.Remove(ingr)
	}
	return nil
}

func (p *Pantry) GetIngredients(_ context.Context) (ingredient.Ingredients, error) {
	return p.ingredients, nil
}

func (p *Pantry) Store(_ context.Context, ingredients ...ingredient.Ingredient) error {
	for idx, ingr := range ingredients {
		if p.ingredients.Has(ingr) {
			p.ingredients[idx].Quantity += ingr.Quantity
		} else {
			p.ingredients = append(p.ingredients, ingr)
		}
	}
	return nil
}
