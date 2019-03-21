package responses

// #2 Product Info
type BlibliProductInfoResponse struct {
	Code   int               `json:"code"`
	Data   BlibliProductInfo `json:"data"`
	Status string            `json:"status"`
}

// #2.1
type BlibliProductInfo struct {
	Options interface{}     `json:"options"`
	Price   BlibliInfoPrice `json:"price"`
}

// #2.2.1
type BlibliInfoPrice struct {
	Discount int     `json:"discount"`
	Listed   float32 `json:"listed"`
	Offered  float32 `json:"offered"`
}

// #3 Product list with info
type BlibliProductWithInfoResponse struct {
	Code   int                   `json:"code"`
	Data   BlibliProductWithInfo `json:"data"`
	Status string                `json:"status"`
}

// #3.1
type BlibliProductWithInfo struct {
	Paging   BlibliPaging             `json:"paging"`
	Products []BlibliProdcutsWithInfo `json:"products"`
}

// #3.2
type BlibliProdcutsWithInfo struct {
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
