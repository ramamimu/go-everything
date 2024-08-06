package product

import "context"

type RepositoryProduct interface {
	AddProduct(ctx context.Context, productID string, amount uint, price float64) error
}

type product struct {
	repository RepositoryProduct
}

func NewProduct(r RepositoryProduct) *product {
	return &product{
		repository: r,
	}
}

func (p product) AddProductItem(ctx context.Context, productID string, amount uint, price float64) error {
	// some business logics...
	err := p.repository.AddProduct(ctx, productID, amount, price)
	if err != nil {
		return err
	}
	// some business logics...
	return nil
}
