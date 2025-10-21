package planner

import "context"

type API1Decorator struct {
	delegate           API1
	CreateCustomerFunc func(context.Context, string) (API1Customer, error)
	GetCustomerFunc    func(context.Context, string) (API1Customer, error)
	UpdateCustomerFunc func(context.Context, string, string) error
}

var _ API1 = &API1Decorator{}

func NewAPI1Delegate(delegate API1) *API1Decorator {
	return &API1Decorator{delegate: delegate}
}

func (a *API1Decorator) CreateCustomer(ctx context.Context, name string) (API1Customer, error) {
	if a.CreateCustomerFunc != nil {
		return a.CreateCustomerFunc(ctx, name)
	}
	return a.delegate.CreateCustomer(ctx, name)
}

func (a *API1Decorator) GetCustomer(ctx context.Context, id string) (API1Customer, error) {
	if a.GetCustomerFunc != nil {
		return a.GetCustomerFunc(ctx, id)
	}
	return a.delegate.GetCustomer(ctx, id)
}

func (a *API1Decorator) UpdateCustomer(ctx context.Context, id string, name string) error {

	if a.UpdateCustomerFunc != nil {
		a.UpdateCustomerFunc(ctx, id, name)
	}
	return a.delegate.UpdateCustomer(ctx, id, name)
}
