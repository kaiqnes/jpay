package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/nuno/nunes-jumia/src/handler"
	"github.com/nuno/nunes-jumia/src/service"
	"net/http"
	"strconv"
	"strings"
)

const (
	countryNameKey = "country_name"
	statusKey      = "status"
	limitKey       = "limit"
	offsetKey      = "offset"
	defaultLimit   = 10
	defaultOffset  = 0
)

//go:generate mockgen -source=./customer-controller.go -destination=./mocks/customer-controller_mock.go
type CustomerController interface {
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

func (controller customerController) GetCustomers(ctx *gin.Context) {
	limit, offset, params, errx := extractQueryParams(ctx)
	if errx != nil {
		ctx.JSON(errx.Status(), errx.JSON())
	}

	outputDto, errx := controller.service.GetCustomers(limit, offset, params)
	if errx != nil {
		ctx.IndentedJSON(errx.Status(), errx.JSON())
	}

	ctx.JSON(http.StatusOK, outputDto)
}

func extractQueryParams(ctx *gin.Context) (int, int, map[string]string, handler.Errorx) {
	var (
		inputCountryName = ctx.Query(countryNameKey)
		inputStatus      = ctx.Query(statusKey)
		inputLimit       = ctx.Query(limitKey)
		inputOffset      = ctx.Query(offsetKey)

		params        = make(map[string]string)
		limit, offset int
	)

	if len(strings.TrimSpace(inputCountryName)) > 0 {
		params[countryNameKey] = inputCountryName
	}

	if len(strings.TrimSpace(inputStatus)) > 0 {
		params[statusKey] = inputStatus
	}

	if len(strings.TrimSpace(inputLimit)) > 0 {
		intLimit, errx := parseInt(inputLimit, limitKey, 32)
		if errx != nil {
			return 0, 0, map[string]string{}, errx
		}
		limit = intLimit
	} else {
		limit = defaultLimit
	}

	if len(strings.TrimSpace(inputOffset)) > 0 {
		intOffset, errx := parseInt(inputOffset, offsetKey, 32)
		if errx != nil {
			return 0, 0, map[string]string{}, errx
		}
		offset = intOffset
	} else {
		offset = defaultOffset
	}

	return limit, offset, params, nil
}

func parseInt(strValue, fieldName string, bitSize int) (int, handler.Errorx) {
	value, err := strconv.ParseInt(strValue, 10, bitSize)
	if err != nil {
		errx := handler.NewError(http.StatusBadRequest, fmt.Sprintf("error to parse %s", fieldName))
		return 0, errx
	}
	return int(value), nil
}
