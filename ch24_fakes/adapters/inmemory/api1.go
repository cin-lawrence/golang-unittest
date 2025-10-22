package inmemory

import (
	"context"
	"errors"
	"strconv"

	"ch24/fakes/domain/planner"
)

func NewAPI1() *API1 {
	return &API1{customers: make(map[string]planner.API1Customer)}
}

type API1 struct {
	idx       int
	customers map[string]planner.API1Customer
}

func (a *API1) CreateCustomer(ctx context.Context, name string) (planner.API1Customer, error) {
	if name == "Dave" {
		return planner.API1Customer{}, errors.New("dave not allowed")
	}

	newCustomer := planner.API1Customer{
		Name: name,
		ID:   strconv.Itoa(a.idx),
	}
	a.customers[newCustomer.ID] = newCustomer
	a.idx++
	return newCustomer, nil
}

func (a *API1) GetCustomer(ctx context.Context, id string) (planner.API1Customer, error) {
	return a.customers[id], nil
}

func (a *API1) UpdateCustomer(ctx context.Context, id string, name string) error {
	customer := a.customers[id]
	customer.Name = name
	a.customers[id] = customer
	return nil
}
