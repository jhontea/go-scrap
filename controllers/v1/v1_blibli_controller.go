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

func (handler *V1BlibliController) GetBlibliProductDetail(context *gin.Context) {
	result, err := handler.BlibliService.GetBlibliProductDetail()

	if err != nil {
		response := handler.ResponseHelper.ErrorResponse(result, err.Error())
		context.JSON(response.Code, response)
		return
	}

	response := handler.ResponseHelper.SuccessResponse(result, "success retrieve data")
	context.JSON(response.Code, response)
}

func (handler *V1BlibliController) GetBlibliProductWithDetail(context *gin.Context) {
	result, err := handler.BlibliService.GetBlibliProductWithDetail()

	if err != nil {
		response := handler.ResponseHelper.ErrorResponse(result, err.Error())
		context.JSON(response.Code, response)
		return
	}

	response := handler.ResponseHelper.SuccessResponse(result, "success retrieve data")
	context.JSON(response.Code, response)
}
