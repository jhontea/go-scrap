package responses

// #4 Product Summary
type BlibliSummaryResponse struct {
	Code   int                  `json:"code"`
	Data   BlibliProductSummary `json:"data"`
	Status string               `json:"status"`
}

// #4.1
type BlibliProductSummary struct {
	Brand            BlibliSummaryBrand    `json:"brand"`
	Code             string                `json:"code"`
	IdType           string                `json:"idType"`
	Images           []BlibliSummaryImages `json:"images"`
	Name             string                `json:"name"`
	ProductItemCode  string                `json:"productItemCode"`
	Sku              string                `json:"sku"`
	Url              string                `json:"url"`
	UrlFriendlyName  string                `json:"urlFriendlyName"`
	UsageProductCode string                `json:"usageProductCode"`
}

// #4.2.1
type BlibliSummaryBrand struct {
	Code     string `json:"code"`
	Name     string `json:"name"`
	Logo     string `json:"Logo"`
	Official bool   `json:"Official"`
}

// #4.2.2
type BlibliSummaryImages struct {
	Full      string `json:"full"`
	Thumbnail string `json:"thumbnail"`
}

// #5 Product Info Summary
type BlibliInfoSummaryResponse struct {
	Code   int                      `json:"code"`
	Data   BlibliProductInfoSummary `json:"data"`
	Status string                   `json:"status"`
}

// #5.1
type BlibliProductInfoSummary struct {
	Brand            BlibliSummaryBrand    `json:"brand"`
	Code             string                `json:"code"`
	IdType           string                `json:"idType"`
	Images           []BlibliSummaryImages `json:"images"`
	Name             string                `json:"name"`
	ProductItemCode  string                `json:"productItemCode"`
	Sku              string                `json:"sku"`
	Url              string                `json:"url"`
	UrlFriendlyName  string                `json:"urlFriendlyName"`
	UsageProductCode string                `json:"usageProductCode"`
	BlibliProductInfo
}
