package responses

type BlibliResponse struct {
	Code   int        `json:"code"`
	Data   BlibliData `json:"data"`
	Status string     `json:"status"`
}

type BlibliData struct {
	Paging   interface{} `json:"paging"`
	Products interface{} `json:"products"`
}
