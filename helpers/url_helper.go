package helpers

import (
	"bytes"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/jhontea/go-scrap/responses"
)

type UrlHelper struct {
	Method    string
	Url       string
	FormParam string
	Header    []UrlHeader
}

type UrlHeader struct {
	Field string
	Value interface{}
}

func (handler *UrlHelper) ProcessRequest() (responses.UrlHttpResponse, error) {
	var result responses.UrlHttpResponse

	// Http request
	client := &http.Client{}
	req, err := http.NewRequest(handler.Method, handler.Url, bytes.NewBuffer([]byte(handler.FormParam)))
	if err != nil {
		log.Println("Error new request | UrlHelper | ProcessRequest")
		return result, err
	}

	// Set header
	// -- Content type
	req.Header.Add("Content-Type", "application/json")
	// set user agent
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/66.0.3359.181 Safari/537.36")
	// TO DO process header

	// Do request
	resp, err := client.Do(req)
	if err != nil {
		log.Println("Error doing request | UrlHelper | ProcessRequest")
		return result, err
	}

	defer resp.Body.Close()

	// Get string body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("Error read body")
		log.Println(err.Error())
		return result, err
	}

	result.Status = resp.Status
	result.StatusCode = resp.StatusCode
	result.Body = string(body)
	result.Header = resp.Header

	return result, nil
}
