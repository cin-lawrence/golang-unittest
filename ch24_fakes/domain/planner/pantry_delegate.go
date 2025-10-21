package planner

import (
	"context"

	"ch24/fakes/domain/ingredient"
)

type PantryDelegate struct {
	delegate Pantry

	GetIngredientsFunc func(context.Context) (ingredient.Ingredients, error)
	StoreFunc          func(context.Context, ...ingredient.Ingredient) error
	RemoveFunc         func(context.Context, ...ingredient.Ingredient) error
}

func NewPantryDelegate(delegate Pantry) *PantryDelegate {
	return &PantryDelegate{delegate: delegate}
}

func (p *PantryDelegate) GetIngredients(ctx context.Context) (ingredient.Ingredients, error) {
	if p.GetIngredientsFunc != nil {
		return p.GetIngredientsFunc(ctx)
	}
	return p.delegate.GetIngredients(ctx)
}

func (p *PantryDelegate) Store(ctx context.Context, ingredients ...ingredient.Ingredient) error {
	if p.StoreFunc != nil {
		return p.StoreFunc(ctx, ingredients...)
	}
	return p.delegate.Store(ctx, ingredients...)
}

func (p *PantryDelegate) Remove(ctx context.Context, ingredients ...ingredient.Ingredient) error {
	if p.RemoveFunc != nil {
		return p.RemoveFunc(ctx)
	}
	return p.delegate.Remove(ctx)
}
