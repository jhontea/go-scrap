package controllers

import (
	"github.com/jhontea/go-scrap/helpers"
)

type Controllers struct {
	ResponseHelper helpers.ResponseHelper
}

type ControllersInterface interface {
	Get()
	Create()
	Delete()
	Update()
}
