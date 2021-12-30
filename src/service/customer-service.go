package service

import (
	"fmt"
	"github.com/nuno/nunes-jumia/src/dto"
	"github.com/nuno/nunes-jumia/src/handler"
	"github.com/nuno/nunes-jumia/src/repository"
	"net/http"
)

//go:generate mockgen -source=./customer-service.go -destination=./mocks/customer-service_mock.go
type CustomerService interface {
	GetCustomers(limit, offset int, params map[string]string) (dto.CustomerOutputDto, handler.Errorx)
}

type customerService struct {
	repository repository.CustomerRepository
}

func NewCustomerService(repository repository.CustomerRepository) CustomerService {
	return &customerService{
		repository: repository,
	}
}

func (service customerService) GetCustomers(limit, offset int, params map[string]string) (dto.CustomerOutputDto, handler.Errorx) {
	total, customers, err := service.repository.GetCustomers(limit, offset, params)
	if err != nil {
		errMsg := fmt.Sprintf("Fail to retrieve customers in DB. Err: %s", err.Error())
		return dto.CustomerOutputDto{}, handler.NewError(http.StatusInternalServerError, errMsg)
	}

	outputDto := dto.NewCustomerOutputDto(total, limit, offset, customers)

	return outputDto, nil
}
