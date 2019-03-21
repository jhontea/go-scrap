package responses

type BlibliResponse struct {
	Code   int         `json:"code"`
	Data   interface{} `json:"data"`
	Status string      `json:"status"`
}

type BlibliProductResponse struct {
	Code   int        `json:"code"`
	Data   BlibliData `json:"data"`
	Status string     `json:"status"`
}

type BlibliData struct {
	Paging   BlibliPaging     `json:"paging"`
	Products []BlibliProdcuts `json:"products"`
}

type BlibliPaging struct {
	Page      int `json:"page"`
	TotalItem int `json:"total_item"`
	TotalPage int `json:"total_page"`
}

type BlibliProdcuts struct {
	Brand       string      `json:"brand"`
	Images      []string    `json:"images"`
	ItemSKU     string      `json:"itemSku"`
	Name        string      `json:"name"`
	Price       interface{} `json:"price"`
	Sku         string      `json:"sku"`
	Status      string      `json:"status"`
	Url         string      `json:"url"`
	FormattedID string      `json:"formattedId"`
}

type BlibliProductDetailResponse struct {
	Code   int                 `json:"code"`
	Data   BlibliProductDetail `json:"data"`
	Status string              `json:"status"`
}

type BlibliProductDetail struct {
	Options interface{} `json:"options"`
	Price   interface{} `json:"price"`
}
