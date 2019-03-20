package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	"github.com/jhontea/go-scrap/db"
	"github.com/jhontea/go-scrap/helpers"
	"github.com/jhontea/go-scrap/responses"
)

func main() {
	// load .env
	fmt.Println("loading .env")
	godotenv.Load()
	fmt.Println("success loading .env")

	db.DBInit()

	route := gin.New()
	route.Use(gin.Logger())
	route.Use(gin.Recovery())

	group := route.Group("/v1")
	{
		group.GET("/product/blibli", GetBlibliProduct)
	}

	serverPort := os.Getenv("PORT")
	serverString := fmt.Sprintf("run on port %s", serverPort)
	fmt.Println(serverString)

	if err := http.ListenAndServe(":"+serverPort, route); err != nil {
		log.Fatal(err)
	}
}

func GetBlibliProduct(context *gin.Context) {
	responseHelper := helpers.ResponseHelper{}
	urlHelper := helpers.UrlHelper{
		Method: "GET",
		Url:    "https://www.blibli.com/backend/search/products?page=1&start=0&searchTerm=nike%20shoes&intent=false&merchantSearch=true&sort=10&category=OL-1000001&customUrl=olahraga-aktivitas-luar-ruang&shipment=blibli&showFacet=true",
	}

	result, _ := urlHelper.ProcessRequest()

	var blibliResponse responses.BlibliResponse
	json.Unmarshal([]byte(result.Body), &blibliResponse)

	response := responseHelper.SuccessResponse(blibliResponse, "success retrieve data")
	context.JSON(response.Code, response)
}
