package product_test

import (
	"context"
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	mock_product "github.com/ramamimu/go-everything/2-mockery/mock"
	"github.com/ramamimu/go-everything/2-mockery/product"
	"github.com/stretchr/testify/assert"
)

func TestAddProductItem_Negative(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mock_product.NewMockRepositoryProduct(ctrl)
	mockRepo.EXPECT().AddProduct(context.Background(), "product-1", uint(2), float64(2000)).Return(errors.New("got an error"))

	p := product.NewProduct(mockRepo)
	err := p.AddProductItem(context.Background(), "product-1", 2, 2000)

	assert.Error(t, err, "expected error")
	assert.ErrorContains(t, err, "got an error")
}

func TestAddProductItem_Positive(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mock_product.NewMockRepositoryProduct(ctrl)
	mockRepo.EXPECT().AddProduct(context.Background(), "product-1", uint(2), float64(2000)).Return(nil)

	p := product.NewProduct(mockRepo)
	err := p.AddProductItem(context.Background(), "product-1", 2, 2000)

	assert.Nil(t, err, "expected error is nil")
	assert.NoError(t, err, "expected no error")
}
