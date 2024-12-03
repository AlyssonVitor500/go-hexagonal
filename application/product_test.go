package application_test

import (
	"testing"
	"github.com/alyssonvitor500/go-hexagonal"

func TestProduct_Enable(t *testing.T) {
	product := application.Product{}
	product.Name = "Hello"
	product.Status = application.DISABLED
	product.Price = 10

	err := product.Enable()
	require.Nil(t, err)

	product.Price = 0

	err = product.Enable()
	require.Equal(t, "the price must be greater than zero", err.Error())
}
