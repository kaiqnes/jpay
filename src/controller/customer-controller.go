package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/nuno/nunes-jumia/src/dto"
	"github.com/nuno/nunes-jumia/src/service"
	"net/http"
)

//go:generate mockgen -source=./customer-controller.go -destination=./mocks/customer-controller_mock.go
type CustomerController interface {
	SetupRoutes(router *gin.Engine) *gin.Engine
	GetCustomers(ctx *gin.Context)
}

type customerController struct {
	service service.CustomerService
}

func NewCustomerController(service service.CustomerService) CustomerController {
	return &customerController{
		service: service,
	}
}

func (controller customerController) SetupRoutes(router *gin.Engine) *gin.Engine {
	router.GET("/customer/search", controller.GetCustomers)

	return router
}

func (controller customerController) GetCustomers(ctx *gin.Context) {
	outputDto, err := controller.service.GetCustomers()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, dto.NewError(err.Error()))
		return
	}

	ctx.JSON(http.StatusOK, outputDto)
}
