package cli

import (
	"fmt"
	"github.com/alyssonvitor500/go-hexagonal/application"
)

func Run(service application.ProductServiceInterface, action string, productId string, productName string, price float64) (string, error) {

	var result = ""

	switch action {
	case "create":
		product, err := service.Create(productName, price)
		if err != nil {
			return result, err
		}

		result = fmt.Sprintf(
			"Product ID %s with the name %s has been created with the price %f and status %s",
			product.GetID(),
			product.GetName(),
			product.GetPrice(),
			product.GetStatus())
	case "enable":
		res, err := enableDisable(service, productId, false)
		if err != nil {
			return result, err
		}

		result = fmt.Sprintf("Product ID %s has been enabled", res.GetID())
	case "disabled":
		res, err := enableDisable(service, productId, true)
		if err != nil {
			return result, err
		}

		result = fmt.Sprintf("Product ID %s has been disabled", res.GetID())
	default:
		product, err := service.Get(productId)
		if err != nil {
			return result, err
		}
		result = fmt.Sprintf(
			"Product ID: %s\nName: %s\nPrice: %f\nStatus: %s",
			product.GetID(),
			product.GetName(),
			product.GetPrice(),
			product.GetStatus())
	}

	return result, nil
}

func enableDisable(service application.ProductServiceInterface, productId string, isDisable bool) (application.ProductInterface, error) {
	product, err := service.Get(productId)
	if err != nil {
		return nil, err
	}

	if isDisable {
		product, err = service.Disable(product)
	} else {
		product, err = service.Enable(product)
	}

	if err != nil {
		return nil, err
	}

	return product, nil
}
