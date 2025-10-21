package planner

import (
	"context"
	"testing"

	"ch24/fakes/internal/expect"
)

type API1Customer struct {
	Name string
	ID   string
}

type API1 interface {
	CreateCustomer(context.Context, string) (API1Customer, error)
	GetCustomer(context.Context, string) (API1Customer, error)
	UpdateCustomer(context.Context, string, string) error
}

type API1Contract struct {
	NewAPI1 func() API1
}

func (c *API1Contract) Test(t *testing.T) {
	t.Run("can create, get and update a customer", func(t *testing.T) {
		var (
			ctx  = context.Background()
			sut  = c.NewAPI1()
			name = "Bob"
		)

		customer, err := sut.CreateCustomer(ctx, name)
		expect.NoErr(t, err)

		got, err := sut.GetCustomer(ctx, customer.ID)
		expect.NoErr(t, err)
		expect.Equal(t, customer, got)

		newName := "Robert"
		expect.NoErr(t, sut.UpdateCustomer(ctx, customer.ID, newName))

		got, err = sut.GetCustomer(ctx, customer.ID)
		expect.NoErr(t, err)
		expect.Equal(t, newName, got.Name)
	})

	t.Run("the system will not allow you to add 'Dave' as a customer", func(t *testing.T) {
		var (
			ctx  = context.Background()
			sut  = c.NewAPI1()
			name = "Dave"
		)

		_, err := sut.CreateCustomer(ctx, name)
		expect.Err(t, err)
	})
}
