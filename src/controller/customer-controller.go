package controller

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/nuno/nunes-jumia/src/dto"
	"github.com/nuno/nunes-jumia/src/service"
	"net/http"
	"strconv"
	"strings"
)

const (
	limitKey      = "limit"
	offsetKey     = "offset"
	defaultLimit  = 10
	defaultOffset = 0
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
	limit, offset, err := extractQueryParams(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, dto.NewError(err.Error()))
		return
	}

	outputDto, err := controller.service.GetCustomers(limit, offset)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, dto.NewError(err.Error()))
		return
	}

	ctx.JSON(http.StatusOK, outputDto)
}

func extractQueryParams(ctx *gin.Context) (int, int, error) {
	var (
		inputLimit    = ctx.Query(limitKey)
		inputOffset   = ctx.Query(offsetKey)
		limit, offset int
	)

	if len(strings.TrimSpace(inputLimit)) > 0 {
		intLimit, err := parseInt(inputLimit, limitKey, 32)
		if err != nil {
			return 0, 0, err
		}
		limit = intLimit
	} else {
		limit = defaultLimit
	}

	if len(strings.TrimSpace(inputOffset)) > 0 {
		intOffset, err := parseInt(inputOffset, offsetKey, 32)
		if err != nil {
			return 0, 0, err
		}
		offset = intOffset
	} else {
		offset = defaultOffset
	}

	return limit, offset, nil
}

func parseInt(strValue, fieldName string, bitSize int) (int, error) {
	value, err := strconv.ParseInt(strValue, 10, bitSize)
	if err != nil {
		return 0, errors.New(fmt.Sprintf("error to parse %s", fieldName))
	}
	return int(value), nil
}
