package cmd

import (
	"fmt"
	"github.com/alyssonvitor500/go-hexagonal/adapters/cli"
	"github.com/spf13/cobra"
)

var action string
var productId string
var productName string
var productPrice float64

// cliCmd represents the cli command
var cliCmd = &cobra.Command{
	Use:   "cli",
	Short: "A command line tool for managing products",
	Long:  `cli is a command line tool for managing products`,
	Run: func(cmd *cobra.Command, args []string) {
		result, err := cli.Run(productService, action, productId, productName, productPrice)
		if err != nil {
			fmt.Println(err.Error())
		}

		fmt.Println(result)
	},
}

func init() {
	rootCmd.AddCommand(cliCmd)
	cliCmd.Flags().StringVarP(&action, "action", "a", "enable", "Enable / Disable a product")
	cliCmd.Flags().StringVarP(&productId, "id", "i", "", "Product ID")
	cliCmd.Flags().StringVarP(&productName, "product", "n", "", "Product name")
	cliCmd.Flags().Float64VarP(&productPrice, "price", "p", 0, "Product price")
}
