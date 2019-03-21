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
