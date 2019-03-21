package responses

type BlibliResponse struct {
	Code   int         `json:"code"`
	Data   interface{} `json:"data"`
	Status string      `json:"status"`
}

// #1
type BlibliProductResponse struct {
	Code   int        `json:"code"`
	Data   BlibliData `json:"data"`
	Status string     `json:"status"`
}

// #1.1
type BlibliData struct {
	Paging   BlibliPaging     `json:"paging"`
	Products []BlibliProdcuts `json:"products"`
}

// #1.1.1
type BlibliPaging struct {
	Page      int `json:"page"`
	TotalItem int `json:"total_item"`
	TotalPage int `json:"total_page"`
}

// #1.1.2
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

// #2
type BlibliProductDetailResponse struct {
	Code   int                 `json:"code"`
	Data   BlibliProductDetail `json:"data"`
	Status string              `json:"status"`
}

// #2.1
type BlibliProductDetail struct {
	Options interface{} `json:"options"`
	Price   interface{} `json:"price"`
}

// #3
type BlibliProductWithDetailResponse struct {
	Code   int                     `json:"code"`
	Data   BlibliProductWithDetail `json:"data"`
	Status string                  `json:"status"`
}

// #3.1
type BlibliProductWithDetail struct {
	Paging   BlibliPaging               `json:"paging"`
	Products []BlibliProdcutsWithDetail `json:"products"`
}

type BlibliProdcutsWithDetail struct {
	Brand       string      `json:"brand"`
	Images      []string    `json:"images"`
	ItemSKU     string      `json:"itemSku"`
	Name        string      `json:"name"`
	Price       interface{} `json:"price"`
	Sku         string      `json:"sku"`
	Status      string      `json:"status"`
	Url         string      `json:"url"`
	FormattedID string      `json:"formattedId"`
	Additional  string      `json:"additional"`
	Options     interface{} `json:"options"`
	PriceInfo   interface{} `json:"price_info"`
}
