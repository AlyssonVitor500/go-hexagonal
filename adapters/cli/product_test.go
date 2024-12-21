package cli_test

import (
	"fmt"
	"github.com/alyssonvitor500/go-hexagonal/adapters/cli"
	mock_application "github.com/alyssonvitor500/go-hexagonal/application/mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestRun(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// Defining params
	productName := "Product Test"
	productPrice := 25.99
	productStatus := "enabled"
	productId := "abc"

	// Defining mocks
	productMock := mock_application.NewMockProductInterface(ctrl)
	productMock.EXPECT().GetID().Return(productId).AnyTimes()
	productMock.EXPECT().GetName().Return(productName).AnyTimes()
	productMock.EXPECT().GetPrice().Return(productPrice).AnyTimes()
	productMock.EXPECT().GetStatus().Return(productStatus).AnyTimes()

	productServiceMock := mock_application.NewMockProductServiceInterface(ctrl)
	productServiceMock.EXPECT().Get(productId).Return(productMock, nil).AnyTimes()
	productServiceMock.EXPECT().Create(productName, productPrice).Return(productMock, nil).AnyTimes()
	productServiceMock.EXPECT().Enable(gomock.Any()).Return(productMock, nil).AnyTimes()
	productServiceMock.EXPECT().Disable(gomock.Any()).Return(productMock, nil).AnyTimes()

	// Create entry test
	expectedResult := fmt.Sprintf(
		"Product ID %s with the name %s has been created with the price %f and status %s",
		productMock.GetID(),
		productMock.GetName(),
		productMock.GetPrice(),
		productMock.GetStatus())

	result, err := cli.Run(productServiceMock, "create", "", productName, productPrice)
	require.Nil(t, err)
	require.Equal(t, expectedResult, result)

	// Enabled entry test
	expectedResult = fmt.Sprintf("Product ID %s has been enabled", productMock.GetID())
	result, err = cli.Run(productServiceMock, "enable", productId, "", 0)
	require.Nil(t, err)
	require.Equal(t, expectedResult, result)

	// Disabled entry test
	expectedResult = fmt.Sprintf("Product ID %s has been disabled", productMock.GetID())
	result, err = cli.Run(productServiceMock, "disable", productId, "", 0)
	require.Nil(t, err)
	require.Equal(t, expectedResult, result)

	// Default entry test
	expectedResult = fmt.Sprintf(
		"Product ID: %s\nName: %s\nPrice: %f\nStatus: %s",
		productMock.GetID(),
		productMock.GetName(),
		productMock.GetPrice(),
		productMock.GetStatus())

	result, err = cli.Run(productServiceMock, "default", productId, "", 0)
	require.Nil(t, err)
	require.Equal(t, expectedResult, result)
}
