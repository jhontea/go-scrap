package v1

import (
	"github.com/gin-gonic/gin"

	"github.com/jhontea/go-scrap/controllers"
	"github.com/jhontea/go-scrap/services/blibli"
)

type V1BlibliController struct {
	controllers.Controllers
	BlibliService blibli.V1BlibliService
}

func (handler *V1BlibliController) GetBlibliProduct(context *gin.Context) {
	result, err := handler.BlibliService.GetBlibliProduct()

	if err != nil {
		response := handler.ResponseHelper.ErrorResponse(result, err.Error())
		context.JSON(response.Code, response)
		return
	}

	response := handler.ResponseHelper.SuccessResponse(result, "success retrieve data")
	context.JSON(response.Code, response)
}

func (handler *V1BlibliController) GetBlibliProductInfo(context *gin.Context) {
	result, err := handler.BlibliService.GetBlibliProductInfo()

	if err != nil {
		response := handler.ResponseHelper.ErrorResponse(result, err.Error())
		context.JSON(response.Code, response)
		return
	}

	response := handler.ResponseHelper.SuccessResponse(result, "success retrieve data")
	context.JSON(response.Code, response)
}

func (handler *V1BlibliController) GetBlibliProductWithInfo(context *gin.Context) {
	result, err := handler.BlibliService.GetBlibliProductWithInfo()

	if err != nil {
		response := handler.ResponseHelper.ErrorResponse(result, err.Error())
		context.JSON(response.Code, response)
		return
	}

	response := handler.ResponseHelper.SuccessResponse(result, "success retrieve data")
	context.JSON(response.Code, response)
}

func (handler *V1BlibliController) GetBlibliProductSummary(context *gin.Context) {
	result, err := handler.BlibliService.GetBlibliProductSummary()

	if err != nil {
		response := handler.ResponseHelper.ErrorResponse(result, err.Error())
		context.JSON(response.Code, response)
		return
	}

	response := handler.ResponseHelper.SuccessResponse(result, "success retrieve data")
	context.JSON(response.Code, response)
}

func (handler *V1BlibliController) GetBlibliProductInfoSummary(context *gin.Context) {
	result, err := handler.BlibliService.GetBlibliProductInfoSummary()

	if err != nil {
		response := handler.ResponseHelper.ErrorResponse(result, err.Error())
		context.JSON(response.Code, response)
		return
	}

	response := handler.ResponseHelper.SuccessResponse(result, "success retrieve data")
	context.JSON(response.Code, response)
}
