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
		v1.GET("/product/blibli/detail", V1BlibliControllerHandler.GetBlibliProductDetail)
		v1.GET("/product/blibli/all", V1BlibliControllerHandler.GetBlibliProductWithDetail)
	}
}
