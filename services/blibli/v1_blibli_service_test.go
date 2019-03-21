package blibli

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetBlibliProduct(t *testing.T) {
	service := V1BlibliService{}

	result, err := service.GetBlibliProduct()

	assert.Nil(t, err)

	fmt.Println(err)
	fmt.Println(result)
}

func TestGetBlibliProductDetail(t *testing.T) {
	service := V1BlibliService{}

	result, err := service.GetBlibliProductDetail()

	assert.Nil(t, err)

	fmt.Println(err)
	fmt.Println(result)
}

func TestGetBlibliProductWithDetail(t *testing.T) {
	service := V1BlibliService{}

	result, err := service.GetBlibliProductWithDetail()

	assert.Nil(t, err)

	fmt.Println(err)
	fmt.Println(result)
}
