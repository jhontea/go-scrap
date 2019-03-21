package helpers

import (
	"net/http"

	"github.com/jhontea/go-scrap/constants"
	"github.com/jhontea/go-scrap/responses"
)

type ResponseHelper struct {
}

func ResponseHelperHandler() ResponseHelper {
	return ResponseHelper{}
}

func (handler *ResponseHelper) BadResponse(data interface{}, message string) responses.Response {
	response := responses.Response{
		Code:    http.StatusBadRequest,
		Data:    data,
		Message: message,
		Status:  constants.STATUS_FAILED,
	}

	return response
}

func (handler *ResponseHelper) ForbiddenResponse(data interface{}, message string) responses.Response {
	response := responses.Response{
		Code:    http.StatusForbidden,
		Data:    data,
		Message: message,
		Status:  constants.STATUS_FAILED,
	}

	return response
}

func (handler *ResponseHelper) SuccessResponse(data interface{}, message string) responses.Response {
	response := responses.Response{
		Code:    http.StatusOK,
		Data:    data,
		Message: message,
		Status:  constants.STATUS_SUCCESS,
	}

	return response
}

func (handler *ResponseHelper) FailedResponse(data interface{}, message string) responses.Response {
	response := responses.Response{
		Code:    http.StatusOK,
		Data:    data,
		Message: message,
		Status:  constants.STATUS_FAILED,
	}

	return response
}

func (handler *ResponseHelper) ErrorResponse(data interface{}, message string) responses.Response {
	response := responses.Response{
		Code:    http.StatusInternalServerError,
		Data:    data,
		Message: message,
		Status:  constants.STATUS_FAILED,
	}

	return response
}
