package ingredient

type Ingredient struct {
	Name     string
	Quantity uint
}

type Ingredients []Ingredient

func (i *Ingredients) Has(ingr Ingredient) bool {
	for _, item := range *i {
		if item.Name == ingr.Name && item.Quantity >= ingr.Quantity {
			return true
		}
	}
	return false
}

func (i *Ingredients) NumberOf(ingrName string) uint {
	for _, item := range *i {
		if item.Name == ingrName {
			return item.Quantity
		}
	}
	return 0
}

func (i *Ingredients) Remove(ingr Ingredient) {
	for idx, item := range *i {
		if item.Name == ingr.Name {
			(*i)[idx].Quantity -= ingr.Quantity
			if (*i)[idx].Quantity == 0 {
				*i = append((*i)[:idx], (*i)[idx+1:]...)
			}
		}
	}
}
