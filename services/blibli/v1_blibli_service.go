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

func (service *V1BlibliService) GetBlibliProductDetailUrl(url string) (responses.BlibliProductDetailResponse, error) {
	urlHelper := helpers.UrlHelper{
		Method: constants.METHOD_GET,
		Url:    url,
	}

	response, err := urlHelper.ProcessRequest()

	var blibliResponse responses.BlibliProductDetailResponse

	if err != nil {
		return blibliResponse, err
	}

	json.Unmarshal([]byte(response.Body), &blibliResponse)

	return blibliResponse, nil
}

func (service *V1BlibliService) GetBlibliProductWithDetail() (responses.BlibliProductWithDetailResponse, error) {
	urlHelper := helpers.UrlHelper{
		Method: constants.METHOD_GET,
		Url:    "https://www.blibli.com/backend/search/products?page=1&start=0&searchTerm=nike%20shoes&intent=false&merchantSearch=true&sort=10&category=OL-1000001&customUrl=olahraga-aktivitas-luar-ruang&shipment=blibli&showFacet=true",
	}

	response, err := urlHelper.ProcessRequest()

	var blibliResponse responses.BlibliProductWithDetailResponse

	if err != nil {
		return blibliResponse, err
	}

	json.Unmarshal([]byte(response.Body), &blibliResponse)

	for key, products := range blibliResponse.Data.Products {
		url := "https://www.blibli.com/backend/product/products/" + products.FormattedID + "/_info?defaultItemSku="
		responseDetail, _ := service.GetBlibliProductDetailUrl(url)

		blibliResponse.Data.Products[key].Additional = url
		blibliResponse.Data.Products[key].Options = responseDetail.Data.Options
		blibliResponse.Data.Products[key].PriceInfo = responseDetail.Data.Price

	}

	return blibliResponse, nil
}
