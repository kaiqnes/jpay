package service

import (
	"errors"
	"fmt"
	"github.com/nuno/nunes-jumia/src/dto"
	"github.com/nuno/nunes-jumia/src/entity"
	"github.com/nuno/nunes-jumia/src/model"
	"github.com/nuno/nunes-jumia/src/repository"
	"regexp"
	"strings"
)

//go:generate mockgen -source=./customer-service.go -destination=./mocks/customer-service_mock.go
type CustomerService interface {
	GetCustomers(limit, offset int, params map[string]string) (dto.CustomerOutputDto, error)
}

type customerService struct {
	repository repository.CustomerRepository
}

func NewCustomerService(repository repository.CustomerRepository) CustomerService {
	return &customerService{
		repository: repository,
	}
}

func (service customerService) GetCustomers(limit, offset int, params map[string]string) (dto.CustomerOutputDto, error) {
	total, customers, err := service.repository.GetCustomers(limit, offset)
	if err != nil {
		errMsg := fmt.Sprintf("Fail to retrieve customers in DB. Err: %s", err.Error())
		return dto.CustomerOutputDto{}, errors.New(errMsg)
	}

	outputDto := buildCustomerOutputDto(total, limit, offset, customers, params)

	return outputDto, nil
}

func buildCustomerOutputDto(total int64, limit, offset int, customers []model.Customer, params map[string]string) (outputDto dto.CustomerOutputDto) {
	regexGetCountryCodeAndPhoneNumber := "^\\((\\d{3})\\)\\s((?:.*))$"
	matcher, _ := regexp.Compile(regexGetCountryCodeAndPhoneNumber)

	for _, customer := range customers {
		matches := matcher.FindStringSubmatch(customer.Phone)
		customerDto := buildCustomerDto(customer, matches)

		if len(params) > 0 {
			customerMatch := filterCustomerByParams(customerDto, params)
			if !customerMatch {
				continue
			}
		}

		outputDto.Customers = append(outputDto.Customers, customerDto)
	}

	outputDto.Total = total
	outputDto.Limit = limit
	outputDto.Offset = offset

	return
}

func filterCustomerByParams(customerDto dto.Customer, params map[string]string) (isValid bool) {
	for key, value := range params {
		switch key {
		case "country_name":
			isValid = strings.EqualFold(customerDto.CountryName, value)
		case "status":
			isValid = strings.EqualFold(customerDto.Status, value)
		}

		if !isValid {
			return
		}
	}
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
