package application_test

import (
	"github.com/codeedu/go-hexagonal/application"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestProduct_Enable(t *testing.T) {
	product := application.Product{}
	product.Name = "Banana"
	product.Status = application.DISABLED
	product.Price = 10

	// a variável err recebe a verificação do metodo Enable()
	err := product.Enable()
	require.Nil(t, err) // verifica que t não pode ser igual a err

	// altera o valor de price e renova a variável err 
	product.Price = 0
	err = product.Enable()

	// testa novamente, esperando uma mensagem de erro (a mesma determinada no product.go)
	require.Equal(t, "The price must be greater then zero to enable the product", err.Error())
}