package main

import (
	"github.com/gin-gonic/gin"

	v1Controller "github.com/jhontea/go-scrap/controllers/v1"
)

func Route(route *gin.Engine) {

	v1 := route.Group("/v1")
	{
		V1BlibliControllerHandler := &v1Controller.V1BlibliController{}
		v1.GET("/product/blibli", V1BlibliControllerHandler.GetBlibliProduct)
		v1.GET("/product/blibli/info", V1BlibliControllerHandler.GetBlibliProductInfo)
		v1.GET("/product/blibli/all", V1BlibliControllerHandler.GetBlibliProductWithInfo)
		v1.GET("/product/blibli/summary", V1BlibliControllerHandler.GetBlibliProductSummary)
		v1.GET("/product/blibli/info-summary", V1BlibliControllerHandler.GetBlibliProductInfoSummary)
	}
}
