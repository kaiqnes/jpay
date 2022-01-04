package service

import (
	"errors"
	"fmt"
	"github.com/nuno/nunes-jumia/src/dto"
	"github.com/nuno/nunes-jumia/src/entity"
	"github.com/nuno/nunes-jumia/src/model"
	"github.com/nuno/nunes-jumia/src/repository"
	"regexp"
)

//go:generate mockgen -source=./customer-service.go -destination=./mocks/customer-service_mock.go
type CustomerService interface {
	GetCustomers(limit, offset int) (dto.CustomerOutputDto, error)
}

type customerService struct {
	repository repository.CustomerRepository
}

func NewCustomerService(repository repository.CustomerRepository) CustomerService {
	return &customerService{
		repository: repository,
	}
}

func (service customerService) GetCustomers(limit, offset int) (dto.CustomerOutputDto, error) {
	total, customers, err := service.repository.GetCustomers(limit, offset)
	if err != nil {
		errMsg := fmt.Sprintf("Fail to retrieve customers in DB. Err: %s", err.Error())
		return dto.CustomerOutputDto{}, errors.New(errMsg)
	}

	outputDto := buildCustomerOutputDto(total, limit, offset, customers)

	return outputDto, nil
}

func buildCustomerOutputDto(total int64, limit, offset int, customers []model.Customer) (outputDto dto.CustomerOutputDto) {
	regexGetCountryCodeAndPhoneNumber := "^\\((\\d{3})\\)\\s((?:.*))$"
	matcher, _ := regexp.Compile(regexGetCountryCodeAndPhoneNumber)

	for _, customer := range customers {
		matches := matcher.FindStringSubmatch(customer.Phone)
		outputDto.Customers = append(outputDto.Customers, buildCustomerDto(customer, matches))
	}

	outputDto.Total = total
	outputDto.Limit = limit
	outputDto.Offset = offset

	return
}

func buildCustomerDto(customer model.Customer, matches []string) (newCustomer dto.Customer) {
	newCustomer.SetFormattedName(customer.Name)

	if matches == nil || len(matches) == 0 {
		return dto.NewUnidentifiedCustomer(customer.Name, customer.Phone)
	}

	countryCode := matches[1]
	phoneNumber := matches[2]
	customerCountry := entity.Countries[countryCode]
	status := customerCountry.IsValidPhoneNumber(customer.Phone)

	return dto.NewIdentifiedCustomer(customer.Name, phoneNumber, countryCode, customerCountry.Name, status)
}
