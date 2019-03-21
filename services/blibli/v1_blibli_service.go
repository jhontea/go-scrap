package blibli

import (
	"encoding/json"

	"github.com/jhontea/go-scrap/constants"
	"github.com/jhontea/go-scrap/helpers"
	"github.com/jhontea/go-scrap/responses"
)

type V1BlibliService struct {
}

func (service *V1BlibliService) GetBlibliProduct() (responses.BlibliProductResponse, error) {
	urlHelper := helpers.UrlHelper{
		Method: constants.METHOD_GET,
		Url:    "https://www.blibli.com/backend/search/products?page=1&start=0&searchTerm=nike%20shoes&intent=false&merchantSearch=true&sort=10&category=OL-1000001&customUrl=olahraga-aktivitas-luar-ruang&shipment=blibli&showFacet=true",
	}

	response, err := urlHelper.ProcessRequest()

	var blibliResponse responses.BlibliProductResponse

	if err != nil {
		return blibliResponse, err
	}

	json.Unmarshal([]byte(response.Body), &blibliResponse)

	return blibliResponse, nil
}

func (service *V1BlibliService) GetBlibliProductDetail() (responses.BlibliProductDetailResponse, error) {
	urlHelper := helpers.UrlHelper{
		Method: constants.METHOD_GET,
		Url:    "https://www.blibli.com/backend/product/products/pc--MTA-2999604/_info?defaultItemSku=",
	}

	response, err := urlHelper.ProcessRequest()

	var blibliResponse responses.BlibliProductDetailResponse

	if err != nil {
		return blibliResponse, err
	}

	json.Unmarshal([]byte(response.Body), &blibliResponse)

	return blibliResponse, nil
}
