package helpers

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProcessRequestGetBlibli(t *testing.T) {

	urlHelper := UrlHelper{
		Method: "GET",
		Url:    "https://www.blibli.com/backend/search/products?page=1&start=0&searchTerm=nike%20shoes&intent=false&merchantSearch=true&sort=10&category=OL-1000001&customUrl=olahraga-aktivitas-luar-ruang&shipment=blibli&showFacet=true",
	}

	actualResult, err := urlHelper.ProcessRequest()

	expectedStatusCode := 200

	assert.Nil(t, err)
	assert.Equal(t, expectedStatusCode, actualResult.StatusCode, "Status must be OK (200)")

	fmt.Println(actualResult)
}
